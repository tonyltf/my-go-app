package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Reponse struct{}

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "/v1/healthcheck", nil)
	if err != nil {
		t.Errorf("Error when doing health checking: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(rr, req)

	response := fmt.Sprintf("%s", rr.Body)
	assert.Equal(response, "ok", "Health checking should response \"ok\"")
}

func TestGetLastPrice(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "/v1/price/BTCUSD", nil)
	if err != nil {
		t.Errorf("Error when getting last price BTCUSD: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLastExchangePrice)
	handler.ServeHTTP(rr, req)

	if price, err := strconv.ParseFloat(fmt.Sprintf("%s", rr.Body), 64); err == nil {
		fmt.Println(price)
		assert.Greater(price, 0.0, "Price should be greater than 0")
	}
	if err != nil {
		t.Errorf("Error when getting last price BTCUSD: %v", err)
	}

	req, err = http.NewRequest("GET", "/v1/price/ETHUSD", nil)
	if err != nil {
		t.Errorf("Error when getting last price ETHUSD: %v", err)
	}

	handler.ServeHTTP(rr, req)

	if price, err := strconv.ParseFloat(fmt.Sprintf("%s", rr.Body), 64); err == nil {
		fmt.Println(price)
		assert.Greater(price, 0.0, "Price should be greater than 0")
	}
	if err != nil {
		t.Errorf("Error when getting last price ETHUSD: %v", err)
	}
}
