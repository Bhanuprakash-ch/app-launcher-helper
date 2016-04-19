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
	"strings"

	"github.com/cloudfoundry/gosteno"
)

type SpaceSummaryHelper struct {
	logger          *gosteno.Logger
}

func NewSpaceSummaryHelper() SpaceSummaryHelper {
	return SpaceSummaryHelper{
		logger:          gosteno.NewLogger("space_summary_helper"),
	}
}

func (p *SpaceSummaryHelper) getMainGuidPart(guid string ) string {
	split := strings.Split(guid,"-")
	mainPart := split[0] + "-" + split[1] + "-" + split[2] + "-" + split[3]
	return mainPart
}

func (p *SpaceSummaryHelper) getAppsByService(planLabel string, summary *ExtendedSpaceSummary, apps map[string]Application) []AtkInstance{
	instances := []AtkInstance{}

	for _, s := range summary.Services {
		if s.ServicePlan.Service.Label == planLabel {
			serviceGuidSuffix := "-" + s.Guid[0:8]

			// Please note that the "if" statements below have side effects. They assign a value to a variable from a map,
			// but using different key. Be careful with possible "optimizations".
			if a, ok := apps[s.Name + serviceGuidSuffix]; ok {
				p.logger.Debug("App name (matched service name): " + s.Name)

				// Information about ATK Instance contains the service name (the name user entered in UI or CLI)
				// and the application URL. This is what makes sense to be presented to the users.
				instances = append(instances, AtkInstance{s.Name, p.getUrl(a.Urls), a.Guid, s.Guid, a.State, s.Metadata})
			} else if a, ok := apps[planLabel + serviceGuidSuffix]; ok {
				p.logger.Debug("App name (matched service plan): " + s.Name)

				// See the above comment.
				instances = append(instances, AtkInstance{s.Name, p.getUrl(a.Urls), a.Guid, s.Guid, a.State, s.Metadata})
			} else {
				p.logger.Warn("App not found for service: " + s.Guid)
			}
		}
	}
	return instances
}

func (p* SpaceSummaryHelper) getUrl(urls []string) string {
	if (len(urls) > 0) {
		return urls[0]
	}
	return "";
}
