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
	"log"
	"net/http"
	"os"

	"github.com/alvjtc/gcp-pricing-info/internal/healthcheck"
	"github.com/gorilla/mux"
)

var port = ":" + os.Getenv("PORT")

func main() {
	s := mux.NewRouter()

	s.HandleFunc("/healthcheck", healthcheck.Handler).Methods(http.MethodGet)

	if port == ":" {
		port = ":8080"
	}

	log.Println("Running server on host", port)
	log.Fatal(http.ListenAndServe(port, s))
}
