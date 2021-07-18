//Copyright 2020 Álvaro José Teijido Carpente
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package compute

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Handler responds to an HTTP request to get the pricing info for Compute.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params, err := parseRequest(r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price, err := getPrice(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseRequest(v url.Values) (ret Request, err error) {
	values := map[string]interface{}{
		"region": "", "family": "", "billing": "", "os": "", "instances": 0, "ram": 0.0, "hours": 0.0, "cpu": 0.0,
	}

	paramsStr := []string{"region", "family", "billing", "os"}
	paramsInt := []string{"instances"}
	paramsFloat := []string{"ram", "hours", "cpu"}

	for _, p := range paramsStr {
		if v[p] != nil {
			values[p] = strings.ToLower(v[p][0])
		}
	}

	for _, p := range paramsInt {
		if v[p] != nil {
			values[p], err = strconv.Atoi(v[p][0])
			if err != nil {
				return ret, err
			}
		}
	}

	for _, p := range paramsFloat {
		if v[p] != nil {
			values[p], err = strconv.ParseFloat(v[p][0], 64)
			if err != nil {
				return ret, err
			}
		}
	}

	ret = Request{
		values["instances"].(int),
		values["region"].(string),
		values["family"].(string),
		values["cpu"].(float64),
		values["ram"].(float64),
		values["billing"].(string),
		values["hours"].(float64),
		values["os"].(string),
	}

	return ret, nil
}
