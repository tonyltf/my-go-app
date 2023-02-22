package model

import (
	"time"

	"github.com/google/uuid"
)

type Rate struct {
	ID           uuid.UUID `json:"id,omitempty"`
	CurrencyPair string    `json:"currency_pair"`
	ExchangeRate float64   `json:"exchange_rate,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type AvgRate struct {
	CurrencyPair  string    `json:"currency_pair"`
	ExchangeRate  float64   `json:"exchange_rate,omitempty"`
	FromCreatedAt time.Time `json:"from_created_at,omitempty"`
	ToCreatedAt   time.Time `json:"to_created_at,omitempty"`
}
