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
)

const hoursInMonth = 730

func getPrice(r Request) (p Price, err error) {
	if !isValidRequest(r) {
		return Price{}, errors.New("invalid request")
	}

	cpuSKUId := SkuVMList[r.Region][r.Family]["cpu"]
	if cpuSKUId == "" {
		return p, errors.New("invalid parameters")
	}

	ramSKUId := SkuVMList[r.Region][r.Family]["ram"]
	if ramSKUId == "" {
		return p, errors.New("invalid parameters")
	}

	price := r.CPU * (float64(SkuPriceList[cpuSKUId].PricingExpression.TieredRates[0].UnitPrice.Units) +
		float64(SkuPriceList[cpuSKUId].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9)

	p.ComputePrice += price

	price = r.RAM * (float64(SkuPriceList[ramSKUId].PricingExpression.TieredRates[0].UnitPrice.Units) +
		float64(SkuPriceList[ramSKUId].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9)

	p.ComputePrice += price

	p.applySUD(r.Billing, r.Hours, r.Family)
	p.applyOsPrice(r)

	p.ComputePrice, p.EffectiveTime, p.Currency = p.ComputePrice*float64(r.Instances), SkuPriceList[cpuSKUId].EffectiveTime, "USD"

	return p, nil
}

func isValidRequest(r Request) bool {
	return !(r.Hours > 730) && !(r.CPU < 0.0) && !(r.Instances < 0) && !(r.RAM < 0.0)
}

func (p *Price) applyOsPrice(r Request) {
	var skuOS string

	switch r.Os {
	case "windows":
		switch r.Family {
		case "f1":
			skuOS = SkuOSList["f1"]
		case "g1":
			skuOS = SkuOSList["f1"]
		default:
			skuOS = SkuOSList["global"]
		}
	default:
		return
	}

	p.OsPrice = r.Hours * r.CPU * (float64(SkuPriceList[skuOS].PricingExpression.TieredRates[0].UnitPrice.Units) +
		float64(SkuPriceList[skuOS].PricingExpression.TieredRates[0].UnitPrice.Nanos)*1e-9)
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
