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

func NewRestServiceCatalog(apiUrl string, accessToken string) service.ServiceCatalog {
	return &RestController{
		client:      &http.Client{},
		apiUrl:      apiUrl,
		accessToken: accessToken,
		logger:      gosteno.NewLogger("service_catalog"),
	}
}

func (rc *RestController) ExtendedSummary(space string) (*service.ExtendedSpaceSummary, error) {
	var summary service.ExtendedSpaceSummary
	return &summary, rc.doGet("/rest/service_instances/extended_summary?space="+space, &summary)
}
