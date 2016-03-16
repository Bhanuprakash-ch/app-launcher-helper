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
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cloudfoundry/gosteno"
)

type RestController struct {
	client      *http.Client
	apiUrl      string
	accessToken string
	logger      *gosteno.Logger
}

func (rc *RestController) doGet(path string, target interface{}) error {
	rc.logger.Debug("GET " + path)
	req, err := http.NewRequest("GET", rc.apiUrl+path, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "bearer "+rc.accessToken)

	resp, err := rc.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	dec := json.NewDecoder(resp.Body)

	return dec.Decode(&target)
}