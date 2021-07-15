package healthcheck

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	type response struct {
		Message string `json:"message"`
	}

	tests := []struct {
		name     string
		args     args
		expected response
	}{
		{"HealthCheck API Call", args{nil, httptest.NewRequest(http.MethodGet, "/healthcheck", nil)}, response{"Server is up and running"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			w := httptest.NewRecorder()

			Handler(w, tt.args.r)
			res := w.Result()
			defer func() {
				if err = res.Body.Close(); err != nil {
					t.Fatalf("%v", err)
				}
			}()

			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("%v", err)
			}

			var got response
			if err = json.Unmarshal(resBody, &got); err != nil {
				t.Fatalf("%v", err)
			}

			if got != tt.expected {
				t.Errorf("expected %+v but got %+v", tt.expected, got)
			}
		})
	}
}
