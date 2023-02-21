package fetcher

import (
	"encoding/json"
	"fmt"
	"my-go-app/internal/database/model"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CoinApi struct {
	Fetcher
}

type CoinApiResponse struct {
	Time   string  `json:"time"`
	Base   string  `json:"asset_id_base"`
	Target string  `json:"asset_id_quote"`
	Price  float64 `json:"rate"`
}

func NewCoinApi() IFetcher {
	return &CoinApi{
		Fetcher: Fetcher{
			apiPath:    "https://rest.coinapi.io/v1/exchangerate/{BASE}/{TARGET}",
			apiKeyName: "X-CoinAPI-Key",
			apiKey:     "43E37135-55C5-4675-8881-2F21932B1467",
			myTransformFunc: func(f *Fetcher) (*model.Rate, error) {
				var m CoinApiResponse
				err := json.Unmarshal(f.myJson, &m)
				if err != nil {
					return nil, fmt.Errorf("Convert JSON error: %w", err)
				}
				CreatedAt, err := time.Parse("2023-02-21T03:40:39.0000000Z", m.Time)
				if err != nil {
					return nil, fmt.Errorf("Parse time error: %w", err)
				}

				NewId := uuid.New()

				rate := model.Rate{
					ID:           NewId,
					CurrencyPair: strings.ToUpper(m.Base + m.Target),
					ExchangeRate: m.Price,
					CreatedAt:    CreatedAt,
				}

				f.myRate = &rate

				return &rate, nil
			},
		},
	}
}
