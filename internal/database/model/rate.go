package model

import (
	"time"

	"github.com/google/uuid"
)

type Rate struct {
	ID           uuid.UUID `json:"id"`
	CurrencyPair string    `json:"currency_pair"`
	ExchangeRate float64   `json:"exchange_rate"`
	CreatedAt    time.Time `json:"created_at"`
}
