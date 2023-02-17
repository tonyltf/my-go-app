package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Reponse struct{}

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/healthcheck", nil)
	if err != nil {
		t.Errorf("Error when doing health checking: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(rr, req)
}
