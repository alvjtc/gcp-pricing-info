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
	"context"

	"google.golang.org/api/cloudbilling/v1"

	"github.com/alvjtc/gcp-pricing-info/internal/google"
)

type SKU = map[string]map[string]map[string]string
type SKUPrice = map[string]*cloudbilling.PricingInfo

var SKUList SKU
var SKUPriceList SKUPrice

func InitData() SKU {
	SKUList = make(map[string]map[string]map[string]string)

	SKUList["europe-west1"] = make(map[string]map[string]string)

	SKUList["europe-west1"]["n1"] = make(map[string]string)
	SKUList["europe-west1"]["n1"]["cpu"] = "9431-52B1-2C4F"
	SKUList["europe-west1"]["n1"]["cpu1y"] = "4F49-1FC5-D994"
	SKUList["europe-west1"]["n1"]["cpu3y"] = "20F3-B410-FB36"
	SKUList["europe-west1"]["n1"]["ram"] = "39F4-0112-6F39"
	SKUList["europe-west1"]["n1"]["ram1y"] = "0FCC-C885-6989"
	SKUList["europe-west1"]["n1"]["ram3y"] = "B1FD-24D4-0892"

	SKUList["europe-west1"]["n2"] = make(map[string]string)
	SKUList["europe-west1"]["n2"]["cpu"] = "9F61-45D7-D4FB"
	SKUList["europe-west1"]["n2"]["cpu1y"] = "A121-1D02-4CFA"
	SKUList["europe-west1"]["n2"]["cpu3y"] = "1438-08DD-CC18"
	SKUList["europe-west1"]["n2"]["ram"] = "A109-54C1-7CB0"
	SKUList["europe-west1"]["n2"]["ram1y"] = "E6C5-0BFA-F6A6"
	SKUList["europe-west1"]["n2"]["ram3y"] = "FEA1-DB7A-4C41"

	return SKUList
}

func InitSKUPriceList(svc *cloudbilling.APIService) error {
	SKUPriceList = make(map[string]*cloudbilling.PricingInfo)

	err := svc.Services.Skus.List(google.ComputeSKU).Pages(context.Background(), func(res *cloudbilling.ListSkusResponse) (err error) {
		for _, sku := range res.Skus {
			SKUPriceList[sku.SkuId] = sku.PricingInfo[0]
		}
		return
	})
	if err != nil {
		return err
	}

	return nil
}
