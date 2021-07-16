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

package google

import "google.golang.org/api/cloudbilling/v1"

const (
	// ComputeSKU SKU Id for Google Compute Engine.
	ComputeSKU = "services/6F81-5844-456A"
)

// Googler type has the API Services from Google Cloud to make API calls.
type Googler struct {
	BillingService *cloudbilling.APIService
}
