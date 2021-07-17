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

/*type skuDataJSON struct {
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
}*/

type SKU = map[string]map[string]map[string]string

var SKUData SKU

func InitData() SKU {
	SKUData = make(map[string]map[string]map[string]string)

	SKUData["europe-west1"] = make(map[string]map[string]string)

	SKUData["europe-west1"]["n1"] = make(map[string]string)
	SKUData["europe-west1"]["n1"]["cpu"] = "9431-52B1-2C4F"
	SKUData["europe-west1"]["n1"]["cpu1y"] = "4F49-1FC5-D994"
	SKUData["europe-west1"]["n1"]["cpu3y"] = "20F3-B410-FB36"
	SKUData["europe-west1"]["n1"]["ram"] = "39F4-0112-6F39"
	SKUData["europe-west1"]["n1"]["ram1y"] = "0FCC-C885-6989"
	SKUData["europe-west1"]["n1"]["ram3y"] = "B1FD-24D4-0892"

	SKUData["europe-west1"]["n2"] = make(map[string]string)
	SKUData["europe-west1"]["n2"]["cpu"] = "9F61-45D7-D4FB"
	SKUData["europe-west1"]["n2"]["cpu1y"] = "A121-1D02-4CFA"
	SKUData["europe-west1"]["n2"]["cpu3y"] = "1438-08DD-CC18"
	SKUData["europe-west1"]["n2"]["ram"] = "A109-54C1-7CB0"
	SKUData["europe-west1"]["n2"]["ram1y"] = "E6C5-0BFA-F6A6"
	SKUData["europe-west1"]["n2"]["ram3y"] = "FEA1-DB7A-4C41"

	return SKUData
}
