package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	myFetcher, err := NewFetcher("random")
	if err != nil {
		t.Errorf("Error in NewFetcher: %v", err)
	}
	assert.NotNil(t, myFetcher)
	err = myFetcher.FetchRate("BTC", "USD")
	if err != nil {
		t.Errorf("Error in FetchRate: %v", err)
	}
	rate, err := myFetcher.TransformRate()
	if err != nil {
		t.Errorf("Error in TransformRate: %v", err)
	}
	assert.NotNil(t, rate)
}
