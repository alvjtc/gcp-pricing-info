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
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port = ":" + os.Getenv("PORT")
var cacheInterval = os.Getenv("CACHE_INTERVAL")

func main() {
	if cacheInterval == "" {
		cacheInterval = "60"
	}

	cacheIntervalValue, err := strconv.Atoi(cacheInterval)
	if err != nil {
		log.Fatal(err)
	}

	err = google.Services.NewGoogler()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Google Services")

	cacheComputeBilling()

	go func() {
		for range time.Tick(time.Minute * time.Duration(cacheIntervalValue)) {
			cacheComputeBilling()
		}
	}()

	s := mux.NewRouter()

	// HealthCheck Endpoint
	s.HandleFunc("/v1/healthcheck", healthcheck.Handler).Methods(http.MethodGet)

	// Compute Endpoint
	s.HandleFunc("/v1/compute", compute.Handler).Methods(http.MethodGet)
	compute.SKUList = compute.InitData()

	if port == ":" {
		port = ":8080"
	}

	hostname, _ := os.Hostname()
	log.Printf("Running server on host %s%s\n", hostname, port)
	log.Fatal(http.ListenAndServe(port, s))
}

func cacheComputeBilling() {
	if err := compute.InitSKUPriceList(google.Services.BillingService); err != nil {
		log.Fatal(err)
	}
	log.Println("Cached Google Compute Billing data")
}
