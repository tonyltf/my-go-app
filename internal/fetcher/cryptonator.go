package fetcher

import (
	"encoding/json"
	"fmt"
	"my-go-app/internal/database/model"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Cryptonator struct {
	Fetcher
}

type Ticker struct {
	Ticker struct {
		Base   string `json:"base"`
		Target string `json:"target"`
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Change string `json:"change"`
	} `json:"ticker"`
	Timestamp int64  `json:"timestamp"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
}

func transformRate(f *Fetcher) (*model.Rate, error) {
	var m Ticker
	err := json.Unmarshal(f.myJson, &m)
	if err != nil {
		return nil, fmt.Errorf("Convert JSON error: %w", err)
	}
	CreatedAt := time.Unix(m.Timestamp, 0)
	ExchangeRate, err := strconv.ParseFloat(m.Ticker.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("Parse price error: %w", err)
	}

	NewId := uuid.New()

	rate := model.Rate{
		ID:           NewId,
		CurrencyPair: m.Ticker.Base + m.Ticker.Target,
		ExchangeRate: ExchangeRate,
		CreatedAt:    CreatedAt,
	}

	f.myRate = &rate

	return &rate, nil
}

func NewCryptonator() IFetcher {
	return &Cryptonator{
		Fetcher: Fetcher{
			apiPath:         "https://api.cryptonator.com/api/ticker/{base}-{target}",
			myTransformFunc: transformRate,
		},
	}
}
