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
package config

import (
	"os"
	"testing"
        s "github.com/trustedanalytics/app-launcher-helper/service"
	. "github.com/onsi/gomega"
)

func TestNewConfig(t *testing.T) {
	RegisterTestingT(t)

        vcapservices := s.VcapServices{[]s.UserProvided{s.UserProvided{UpsiCredentials: s.Credentials{Host:"serviceCatalogUrl"}, UpsiName:"servicecatalog"}}}

	expected := Config{
		ApiUrl:      "apiurl",
		TokenKeyUrl: "tokenKeyurl",
                VcapServicesRaw: "{\"user-provided\":[{\"credentials\":{\"host\":\"serviceCatalogUrl\"},\"name\":\"servicecatalog\"}]}",
                VcapServices: vcapservices,
                ServiceCatalogUrl: "serviceCatalogUrl",
	}

	os.Setenv("API_URL", expected.ApiUrl)
	os.Setenv("TOKEN_KEY_URL", expected.TokenKeyUrl)
        os.Setenv("VCAP_SERVICES", expected.VcapServicesRaw)

	c := NewConfig()

	Expect(c).To(Equal(&expected))
}
