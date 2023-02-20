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
	rate_1 := GetRate(ctx, "XXXXXX", nil)
	assert.Equal(t, rate_1, nil)

	rate_2 := GetRate(ctx, "BTCUSD", nil)
	assert.NotEmpty(t, rate_2)
}
func TestGetRate(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	unix := fmt.Sprintf("%v", now.Unix())
	rate_1 := GetRate(ctx, "XXXXXX", &unix)
	assert.Equal(t, rate_1, nil)

	rate_2 := GetRate(ctx, "BTCUSD", &unix)
	assert.NotEmpty(t, rate_2)
}
