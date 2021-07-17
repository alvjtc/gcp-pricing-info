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

package main

import (
	"github.com/alvjtc/gcp-pricing-info/internal/api/compute"
	"github.com/alvjtc/gcp-pricing-info/internal/api/healthcheck"
	"github.com/alvjtc/gcp-pricing-info/internal/google"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port = ":" + os.Getenv("PORT")

func main() {
	err := google.Services.NewGoogler()
	if err != nil {
		log.Fatal(err)
	}

	s := mux.NewRouter()

	// HealthCheck Endpoint
	s.HandleFunc("/v1/healthcheck", healthcheck.Handler).Methods(http.MethodGet)

	// Compute Endpoint
	s.HandleFunc("/v1/compute", compute.Handler).Methods(http.MethodGet)
	compute.SKUData = compute.InitData()

	if port == ":" {
		port = ":8080"
	}

	log.Println("Running server on host", port)
	log.Fatal(http.ListenAndServe(port, s))
}
