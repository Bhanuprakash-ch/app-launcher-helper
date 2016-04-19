/**
 * Copyright (c) 2015 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package service

import (
	"errors"
	"sort"
	"github.com/cloudfoundry/gosteno"
)

type AtkInstance struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Guid        string `json:"guid"`
	ServiceGuid string `json:"service_guid"`
	State       string `json:"state"`
	Metadata    InstanceMetadata `json:"metadata"`
}

// AtkInstances
// swagger:response AtkInstancesResponse
type AtkInstancesResponse struct {
	// in: body
	Body AtkInstances
}

type AtkInstances struct {
	Instances         []AtkInstance `json:"instances"`
	ServicePlanGuid   string        `json:"service_plan_guid"`
}

type InstanceMetadata struct {
	CreatorGuid string `json:"creator_guid"`
	CreatorName string `json:"creator_name"`
}

type ByName []AtkInstance

func (a ByName) Len() int { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (a *AtkInstances) Append(another *AtkInstances) {
	if another.ServicePlanGuid != "" {
		a.ServicePlanGuid = another.ServicePlanGuid
	}

	a.Instances = append(a.Instances, another.Instances...)
}

func (a *AtkInstances) Sort() {
	sort.Sort(ByName(a.Instances))}

type AtkListService struct {
	SpaceSummaryHelper SpaceSummaryHelper
	cloudController    CloudController
	serviceCatalog	   ServiceCatalog
	logger          *gosteno.Logger
}

func NewAtkListService(cloudController CloudController, serviceCatalog ServiceCatalog, SpaceSummaryHelper SpaceSummaryHelper) *AtkListService {
	return &AtkListService{
		cloudController: 	cloudController,
		serviceCatalog: 	serviceCatalog,
		SpaceSummaryHelper: SpaceSummaryHelper,
		logger:          	gosteno.NewLogger("atk_list_service"),
	}
}

func (p *AtkListService) getSpaceList(orgId string) ([]string, error) {
	spaces, err := p.cloudController.Spaces(orgId)
	if err != nil {
		return nil, err
	}

	return spaces.IdList(), nil
}

func (p *AtkListService) servicePlanId(Name string) (string, error) {
	services, err := p.cloudController.ServicesFiltered(Name)
	if err != nil {
		return "", err
	}
	var servicePlansUrl string

	for _, r := range services.Resources {
		if r.Entity.Label == Name {
			servicePlansUrl = r.Entity.ServicePlansUrl
		}
	}

	if len(servicePlansUrl) == 0 {
		return "", errors.New("Service plans url is empty")
	}

	servicePlans, err := p.cloudController.ServicePlans(servicePlansUrl)
	if err != nil {
		return "", err
	}

	if len(servicePlans.Resources) == 0 {
		return "", errors.New("Could not find any service plan for: "+ Name)
	}

	return servicePlans.Resources[0].Metadata.Id, nil
}

func (p *AtkListService) getSpaceInstances(atkLabel string,
	space string,
	instanceChan chan AtkInstances,
	errorChan chan error) {
	summary, err := p.serviceCatalog.ExtendedSummary(space)

	if err != nil {
		errorChan <- err
		return
	}
	atkInstanceList := p.getInstancesFromSpaceSummary(atkLabel, summary);

	atkPlan, err := p.servicePlanId(atkLabel)

	if err != nil {
		p.logger.Warn("Failed to fetch service plan for label: " + atkLabel)
	}

	instanceChan <-AtkInstances{atkInstanceList, atkPlan}
}

func (p *AtkListService) getInstancesFromSpaceSummary(atkLabel string,
	summary *ExtendedSpaceSummary) []AtkInstance {
	apps := make(map[string]Application)
	for _, a := range summary.Apps {
		apps[a.Name] = a
	}

	instances := p.SpaceSummaryHelper.getAppsByService(atkLabel, summary, apps)
	return instances
}


func (p *AtkListService) GetAllInstances(atkLabel string, orgId string) (*AtkInstances, error) {
	spaceList, err := p.getSpaceList(orgId)
	if err != nil {
		return nil, err
	}

	instanceChan := make(chan AtkInstances)
	errorChan := make(chan error)

	for _, s := range spaceList {
		go p.getSpaceInstances(atkLabel, s, instanceChan, errorChan)
	}

	atkInstances := AtkInstances{}
	for _, _ = range spaceList {
		select {
		case spaceInstances := <-instanceChan:
			atkInstances.Append(&spaceInstances)
		case err = <-errorChan:
			p.logger.Warn(err.Error())
		}
	}

	return &atkInstances, nil
}
