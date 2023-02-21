package fetcher

import (
	"math/rand"
	"my-go-app/internal/database/model"
	"time"

	"github.com/google/uuid"
)

type Random struct {
	Fetcher
}

func NewRandom() IFetcher {
	return &Random{
		Fetcher: Fetcher{
			apiPath: "",
			myFetchFunc: func(f *Fetcher, base, target string) error {
				return nil
			},
			myTransformFunc: func(f *Fetcher) (*model.Rate, error) {
				return &model.Rate{
					ID:           uuid.New(),
					CurrencyPair: "BTCUSD",
					ExchangeRate: 20000.0 + rand.Float64()*10000,
					CreatedAt:    time.Now(),
				}, nil
			},
		},
	}
}
