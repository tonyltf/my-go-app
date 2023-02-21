package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLastRate(t *testing.T) {
	ctx := context.Background()
	_, err := GetRate(ctx, "XXXXXX", nil)
	assert.NotNil(t, err)

	rate_2, err := GetRate(ctx, "BTCUSD", nil)
	if err != nil {
		t.Errorf("Error in TestGetLastRate: %v", err)
	}
	assert.NotEmpty(t, *rate_2)
}
func TestGetRate(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	unix := fmt.Sprintf("%v", now.Unix())
	_, err := GetRate(ctx, "XXXXXX", &unix)
	assert.NotNil(t, err)

	rate_2, err := GetRate(ctx, "BTCUSD", &unix)
	if err != nil {
		t.Errorf("Error in TestGetRate: %v", err)
	}
	assert.NotEmpty(t, *rate_2)
}
