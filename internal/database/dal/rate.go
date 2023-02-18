package dal

import (
	"context"
	"database/sql"
	"fmt"
	"my-go-app/internal/database/model"

	"github.com/google/uuid"
)

func Create(ctx context.Context, db *sql.DB, rate *model.Rate) (*model.Rate, error) {

	r := &model.Rate{
		ID:           uuid.New(),
		CurrencyPair: rate.CurrencyPair,
		ExchangeRate: rate.ExchangeRate,
		CreatedAt:    rate.CreatedAt,
	}

	const stmt = `INSERT INTO rate (id, currency_pair, exchange_rate, created_at) VALUES (?, ?, ?, ?)`
	if _, err := db.ExecContext(ctx, stmt, r.ID, r.CurrencyPair, r.ExchangeRate, r.CreatedAt); err != nil {
		return nil, fmt.Errorf("reate create error %w", err)
	}
	return r, nil

}
