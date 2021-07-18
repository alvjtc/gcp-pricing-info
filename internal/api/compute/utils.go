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
	"errors"
	"strings"

	"google.golang.org/api/cloudbilling/v1"
)

const hoursInMonth = 730

func getPrice(svc *cloudbilling.APIService, r Request) (p Price, err error) {
	if !isValidRequest(r) {
		return Price{}, errors.New("invalid request")
	}

	cpuSKUId := SKUList[r.Region][r.Family]["cpu"]
	if cpuSKUId == "" {
		return p, errors.New("invalid parameters")
	}

	ramSKUId := SKUList[r.Region][r.Family]["ram"]
	if ramSKUId == "" {
		return p, errors.New("invalid parameters")
	}

	price := r.CPU * (float64(SKUPriceList[cpuSKUId].PricingExpression.TieredRates[0].UnitPrice.Units) +
		float64(SKUPriceList[cpuSKUId].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9)

	p.ComputePrice += price

	price = r.RAM * (float64(SKUPriceList[ramSKUId].PricingExpression.TieredRates[0].UnitPrice.Units) +
		float64(SKUPriceList[ramSKUId].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9)

	p.ComputePrice += price

	p.applySUD(r.Billing, r.Hours, r.Family)
	p.applyOsPrice(r)
	p.ComputePrice, p.EffectiveTime = p.ComputePrice*float64(r.Instances), SKUPriceList[cpuSKUId].EffectiveTime

	p.Currency = "USD"
	return p, nil

	/*foundCPU, foundRAM, foundOS := false, false, false

	if r.Os == "linux" {
		foundOS = true
	}

	err = svc.Services.Skus.List(google.ComputeSKU).Pages(context.Background(), func(res *cloudbilling.ListSkusResponse) (err error) {
		if foundCPU && foundRAM && foundOS {
			return
		}

		for _, sku := range res.Skus {
			switch sku.SkuId {
			case cpuSKUId:
				foundCPU = true
				price := float64(sku.PricingInfo[0].PricingExpression.TieredRates[0].UnitPrice.Units) +
					float64(sku.PricingInfo[0].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9
				price *= r.CPU
				p.ComputePrice += price
			case ramSKUId:
				foundRAM = true
				price := float64(sku.PricingInfo[0].PricingExpression.TieredRates[0].UnitPrice.Units) +
					float64(sku.PricingInfo[0].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9
				price *= r.RAM
				p.ComputePrice += price
			}

			if foundCPU && foundRAM && foundOS {
				p.applySUD(r.Billing, r.Hours, r.Family)
				p.applyOsPrice(r)
				p.ComputePrice, p.EffectiveTime = p.ComputePrice*float64(r.Instances), sku.PricingInfo[0].EffectiveTime

				return
			}
		}
		return
	})
	if err != nil {
		return p, err
	}
	*/
	p.Currency = "USD"
	return p, nil
}

func isValidRequest(r Request) bool {
	return !(r.Hours > 730) && !(r.CPU < 0.0) && !(r.Instances < 0) && !(r.RAM < 0.0)
}

func (p *Price) applyOsPrice(r Request) {
	if strings.Contains(r.Family, "f1") || strings.Contains(r.Family, "g1") {
		p.OsPrice *= r.Hours
	} else {
		p.OsPrice *= r.Hours * r.CPU
	}
}

func getSUD(h float64) (r SUD) {
	const q = hoursInMonth * 0.25
	for i := 0; i < len(r.Hours); i++ {
		if h > q {
			r.Hours[i] = q
			h -= q
		} else {
			r.Hours[i] = h
			h = 0
		}
	}

	return
}

func (p *Price) applySUD(b string, h float64, t string) {
	sud := getSUD(h)

	if b != "sud" {
		sud.Percentage[0] = 1
		sud.Percentage[1] = 1
		sud.Percentage[2] = 1
		sud.Percentage[3] = 1
	} else {
		switch {
		case t == FamilyG1 || t == FamilyF1 || t == FamilyN1 || t == FamilyM1 || t == FamilyM2:
			sud.Percentage[0] = 1
			sud.Percentage[1] = 0.8
			sud.Percentage[2] = 0.6
			sud.Percentage[3] = 0.4
		case t == FamilyN2 || t == FamilyN2D || t == FamilyC2:
			sud.Percentage[0] = 1
			sud.Percentage[1] = 0.8678
			sud.Percentage[2] = 0.733
			sud.Percentage[3] = 0.6
		default:
			sud.Percentage[0] = 1
			sud.Percentage[1] = 1
			sud.Percentage[2] = 1
			sud.Percentage[3] = 1
		}
	}

	priceUnit := p.ComputePrice
	p.ComputePrice = 0

	for i := 0; i < len(sud.Percentage); i++ {
		p.ComputePrice += priceUnit * sud.Hours[i] * sud.Percentage[i]
	}
}
