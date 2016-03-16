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
package cc

import (
	"net/http"

	"github.com/cloudfoundry/gosteno"

	"github.com/trustedanalytics/app-launcher-helper/service"
)

func NewRestCloudController(apiUrl string, accessToken string) service.CloudController {
	return &RestController{
		client:      &http.Client{},
		apiUrl:      apiUrl,
		accessToken: accessToken,
		logger:      gosteno.NewLogger("cc_client"),
	}
}

func (rc *RestController) Spaces(organization string) (*service.ResourceList, error) {
	var resList service.ResourceList
	return &resList, rc.doGet("/v2/organizations/"+organization+"/spaces", &resList)
}

func (rc *RestController) SpaceSummary(space string) (*service.SpaceSummary, error) {
	var summary service.SpaceSummary
	return &summary, rc.doGet("/v2/spaces/"+space+"/summary", &summary)
}

func (rc *RestController) Services() (*service.ResourceList, error) {
	var services service.ResourceList
	return &services, rc.doGet("/v2/services", &services)
}

func (rc *RestController) ServicePlans(servicePlansUrl string) (*service.ResourceList, error) {
	var services service.ResourceList
	return &services, rc.doGet(servicePlansUrl, &services)
}
