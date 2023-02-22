package fetcher

import (
	"fmt"
	"my-go-app/internal/database/model"
	"strconv"
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
					return fmt.Errorf("Api Ticker error: %v", err)
				}

				newBase := strings.Replace(base, "BTC", "XXBTZ", 1)
				newTarget := strings.Replace(target, "BTC", "XXBTZ", 1)
				exchangeRate, err := strconv.ParseFloat(ticker.GetPairTickerInfo(newBase + newTarget).Ask[0], 64)
				if err != nil {
					return fmt.Errorf("Parse asking price error: %v", err)
				}

				rate := &model.Rate{
					ID:           uuid.New(),
					CurrencyPair: strings.ToUpper(base + target),
					ExchangeRate: exchangeRate,
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
