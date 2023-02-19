package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	config "my-go-app/configs"
	"my-go-app/internal/database/model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

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

func FetchRate(ctx context.Context, base string, target string) (*model.Rate, error) {
	// fmt.Println(ctx.Value("apiSource"))
	// apiSource := ctx.Value("apiSource").(string)
	envConfig := config.InitConfig()
	apiSource := envConfig.ApiSource
	apiSource = strings.Replace(apiSource, "{base}", strings.ToLower(base), 1)
	apiSource = strings.Replace(apiSource, "{target}", strings.ToLower(target), 1)
	fmt.Println(apiSource)

	res, err := http.Get(apiSource)
	if err != nil {
		return nil, fmt.Errorf("API error %w", err)
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Http status: " + res.Status)
	}

	b, err := io.ReadAll(res.Body)

	var m Ticker
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, fmt.Errorf("Convert JSON error: %w", err)
	}

	CreatedAt := time.Unix(m.Timestamp, 0)
	ExchangeRate, err := strconv.ParseFloat(m.Ticker.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("Parse price error: %w", err)
	}

	rate := model.Rate{
		ID:           uuid.New(),
		CurrencyPair: m.Ticker.Base + m.Ticker.Target,
		ExchangeRate: ExchangeRate,
		CreatedAt:    CreatedAt,
	}

	return &rate, nil
}
