package dal

import (
	"context"
	"database/sql"
	"fmt"
	"my-go-app/internal/database/model"
	"time"

	"github.com/google/uuid"
)

func Create(ctx context.Context, db *sql.DB, rate *model.Rate) (*model.Rate, error) {

	if rate == nil {
		return nil, fmt.Errorf("Rate is nil")
	}
	r := &model.Rate{
		ID:           rate.ID,
		CurrencyPair: rate.CurrencyPair,
		ExchangeRate: rate.ExchangeRate,
		CreatedAt:    rate.CreatedAt,
	}

	const stmt = `INSERT INTO rate (id, currency_pair, exchange_rate, created_at) VALUES (?, ?, ?, ?)`
	if _, err := db.ExecContext(ctx, stmt, r.ID, r.CurrencyPair, r.ExchangeRate, r.CreatedAt); err != nil {
		return nil, fmt.Errorf("Create error %w", err)
	}
	return r, nil

}

func Read(ctx context.Context, db *sql.DB, currencyPair string, createdAt *time.Time) (*model.Rate, error) {
	var id string
	var currency_pair string
	var exchange_rate float64
	var created_at time.Time
	var stmt string
	var row *sql.Rows
	var err error

	if createdAt != nil {
		stmt = `SELECT id, currency_pair, exchange_rate, created_at FROM rate WHERE currency_pair = ? AND created_at <= ? ORDER BY created_at ASC`
		row, err = db.QueryContext(ctx, stmt, currencyPair, *createdAt)
	} else {
		stmt = `SELECT id, currency_pair, exchange_rate, created_at FROM rate WHERE currency_pair = ? ORDER BY created_at DESC`
		row, err = db.QueryContext(ctx, stmt, currencyPair)
	}

	if err != nil {
		return nil, fmt.Errorf("Read error %w", err)
	}

	defer row.Close()
	if row.Next() {
		err = row.Scan(&id, &currency_pair, &exchange_rate, &created_at)
		if err != nil {
			return nil, fmt.Errorf("Scan error %w", err)
		}
		return &model.Rate{
			ID:           uuid.MustParse(id),
			CurrencyPair: currency_pair,
			ExchangeRate: exchange_rate,
			CreatedAt:    created_at,
		}, nil
	}
	return nil, nil
}
