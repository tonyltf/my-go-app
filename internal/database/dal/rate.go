package dal

import (
	"context"
	"database/sql"
	"fmt"
	"my-go-app/internal/database/model"
	"time"

	"github.com/google/uuid"
)

type Dal struct {
	Ctx context.Context
	Db  *sql.DB
}

func (d *Dal) Create(rate *model.Rate) (*model.Rate, error) {

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
	if _, err := d.Db.ExecContext(d.Ctx, stmt, r.ID, r.CurrencyPair, r.ExchangeRate, r.CreatedAt); err != nil {
		return nil, fmt.Errorf("Create error %w", err)
	}
	return r, nil

}

func (d *Dal) Read(currencyPair string, createdAt *time.Time) (*model.Rate, error) {
	var id string
	var currency_pair string
	var exchange_rate float64
	var created_at time.Time
	var stmt string
	var row *sql.Rows
	var err error

	if createdAt != nil {
		stmt = `SELECT id, currency_pair, exchange_rate, created_at FROM rate WHERE currency_pair = ? AND created_at <= ? ORDER BY created_at ASC`
		row, err = d.Db.QueryContext(d.Ctx, stmt, currencyPair, *createdAt)
	} else {
		stmt = `SELECT id, currency_pair, exchange_rate, created_at FROM rate WHERE currency_pair = ? ORDER BY created_at DESC`
		row, err = d.Db.QueryContext(d.Ctx, stmt, currencyPair)
	}

	if err != nil {
		return nil, fmt.Errorf("Query error %w", err)
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

func (d *Dal) ReadRange(currencyPair string, startTimestamp time.Time, endTimestamp time.Time) (*model.AvgRate, error) {
	const stmt = `SELECT SUM(exchange_rate)/COUNT(1) AS avg FROM rate WHERE currency_pair = ? AND created_at >= ? AND created_at <= ?`
	row, err := d.Db.QueryContext(d.Ctx, stmt, currencyPair, startTimestamp, endTimestamp)
	if err != nil {
		return nil, fmt.Errorf("Query error %w", err)
	}
	defer row.Close()
	var avg float64
	if row.Next() {
		err = row.Scan(&avg)
		if err != nil {
			return nil, fmt.Errorf("Scan error %w", err)
		}
		return &model.AvgRate{
			CurrencyPair:  currencyPair,
			ExchangeRate:  avg,
			FromCreatedAt: startTimestamp,
			ToCreatedAt:   endTimestamp,
		}, nil
	}

	return nil, nil
}
