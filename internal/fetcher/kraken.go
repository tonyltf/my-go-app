package fetcher

import (
	"fmt"
	"my-go-app/internal/database/model"
	"strings"
	"time"

	krakenapi "github.com/beldur/kraken-go-api-client"
	"github.com/google/uuid"
)

type Kraken struct {
	Fetcher
}

func NewKraken() IFetcher {
	return &Kraken{
		Fetcher: Fetcher{
			apiPath: "",
			myFetchFunc: func(f *Fetcher, base, target string) error {
				api := krakenapi.New("", "")
				ticker, err := api.Ticker(base + target)
				if err != nil {
					return fmt.Errorf("Api Ticker error: %w", err)
				}

				newBase := strings.Replace(base, "BTC", "XXBTZ", 1)
				newTarget := strings.Replace(target, "BTC", "XXBTZ", 1)

				rate := &model.Rate{
					ID:           uuid.New(),
					CurrencyPair: strings.ToUpper(base + target),
					ExchangeRate: ticker.GetPairTickerInfo(newBase + newTarget).OpeningPrice,
					CreatedAt:    time.Now(),
				}
				f.myRate = rate
				return nil
			},
			myTransformFunc: func(f *Fetcher) (*model.Rate, error) {
				rate := f.myRate
				return rate, nil
			},
		},
	}
}
