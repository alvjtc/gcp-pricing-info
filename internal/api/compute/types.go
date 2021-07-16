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

const (
	// FamilyG1 Identification for G1 Small instances Family.
	FamilyG1 = "g1"
	// FamilyF1 Identification for F1 Micro instances Family.
	FamilyF1 = "f1"
	// FamilyN1 Identification for N1 instances Family.
	FamilyN1 = "n1"
	// FamilyN2 Identification for N2 instances Family.
	FamilyN2 = "n2"
	// FamilyN2D Identification for N2D instances Family.
	FamilyN2D = "n2d"
	// FamilyE2 Identification for E2 instances Family.
	FamilyE2 = "e2"
	// FamilyC2 Identification for C2 instances Family.
	FamilyC2 = "c2"
	// FamilyM1 Identification for M1 instances Family.
	FamilyM1 = "m1"
	// FamilyM2 Identification for M2 instances Family.
	FamilyM2 = "m2"

	// BillingPreemptible Identification for Preemptible instances.
	BillingPreemptible = "preemptible"
)

// Request type has all the fields in the HTTP request.
type Request struct {
	Instances int     `json:"instances"`
	Region    string  `json:"region"`
	Family    string  `json:"family"`
	CPU       float64 `json:"cpu"`
	RAM       float64 `json:"ram"`
	Billing   string  `json:"billing"`
	Hours     float64 `json:"hours"`
	Os        string  `json:"os"`
}

// Price type has all the fields in the HTTP response.
type Price struct {
	Currency      string  `json:"currency"`
	ComputePrice  float64 `json:"computePrice"`
	OsPrice       float64 `json:"osPrice"`
	EffectiveTime string  `json:"effectiveTime"`
}

// SUD has arrays with the hours and percentage for calculating the discounts.
type SUD struct {
	Hours      [4]float64
	Percentage [4]float64
}
