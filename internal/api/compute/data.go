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
	"io/ioutil"
	"os"
	"path/filepath"
)

type skuDataJSON struct {
	Compute struct {
		Region []struct {
			Name   string `json:"name"`
			Family []struct {
				Name  string `json:"name"`
				Cpu   string `json:"cpu"`
				Cpu1y string `json:"cpu1y"`
				Cpu3y string `json:"cpu3y"`
				Ram   string `json:"ram"`
				Ram1y string `json:"ram1y"`
				Ram3y string `json:"ram3y"`
			}
		}
	}
}

var SKUData skuDataJSON

func InitData() (skuDataJSON, error) {
	var data skuDataJSON

	filePath, err := filepath.Abs("./internal/data/skus.json")
	if err != nil {
		return data, err
	}

	jsonFile, err := os.Open(string(filePath))
	if err != nil {
		return data, err
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
