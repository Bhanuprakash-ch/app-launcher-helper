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

func (rl *ResourceList) Contains(Id string) bool {
	for _, r := range rl.Resources {
		if r.Metadata.Id == Id {
			return true
		}
	}

	return false
}

func (rl *ResourceList) IdList() []string {
	ids := make([]string, rl.Count)
	for i, r := range rl.Resources {
		ids[i] = r.Metadata.Id
	}

	return ids
}

type CloudController interface {
	Spaces(organization string) (*ResourceList, error)
	SpaceSummary(space string) (*SpaceSummary, error)
	Services() (*ResourceList, error)
	ServicesFiltered(Name string) (*ResourceList, error)
	ServicePlans(Name string) (*ResourceList, error)
}
