package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastRate(t *testing.T) {
	ctx := context.Background()
	rate_1 := GetLastRate(ctx, "XXXXXX")
	assert.Nil(t, rate_1)

	rate_2 := GetLastRate(ctx, "BTCUSD")
	assert.NotEmpty(t, rate_2)
}
