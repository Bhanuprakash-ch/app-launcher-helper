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
	"testing"
	. "github.com/onsi/gomega"
)

type MockCloudController struct {

}

func (m *MockCloudController) ServicePlans(Name string) (*ResourceList, error) {
	if Name == "atk_plans_url" {
		servicePlans := &ResourceList{1, []Resource{Resource{ResourceMetadata{"atk_plan_guid", "atk_plan_url"}, ResourceEntity{"", ""}}}}
		return servicePlans, nil
	}
	otherPlans := &ResourceList{1, []Resource{Resource{ResourceMetadata{"other_plan_guid", "other_plan_url"}, ResourceEntity{"", ""}}}}
	return otherPlans, nil
}

func (m *MockCloudController) Services() (*ResourceList, error) {
	services := &ResourceList{1, []Resource{Resource{ResourceMetadata{"atk_guid","atk_url"}, ResourceEntity{"atk","atk_plans_url"}}}}
	// method not called in test
	return services, nil
}

func (m *MockCloudController ) Spaces(organization string) (*ResourceList, error) {
	var services *ResourceList
	// method not called in test
	return services, nil
}

func (m *MockCloudController ) SpaceSummary(space string) (*SpaceSummary, error) {
	var summary * SpaceSummary
	// method not called in test
	return summary, nil
}

func TestGetServicePlanId(t *testing.T) {
	RegisterTestingT(t)
	cc := new(MockCloudController)

	ssh := NewSpaceSummaryHelper()
	srv := NewAtkListService(cc, ssh)

	atkPlan, err := srv.servicePlanId("atk")
	Expect(atkPlan).To(Equal("atk_plan_guid"));

	sePlan, err := srv.servicePlanId("another")
	Expect(sePlan).To(Equal(""));
	Expect(err).NotTo(Equal(nil));
}
