package healthcheck

import "net/http"

// Handler responds to a HealthCheck request to verify that the API is running
func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"message":"Server is up and running"}`))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
