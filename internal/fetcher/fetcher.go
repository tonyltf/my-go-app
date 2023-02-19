package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	config "my-go-app/init"
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

type MyTime struct {
	time.Time
}

func (m *MyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	fmt.Println(data)
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = MyTime{tt}
	return err
}

func FetchRate(base string, target string) (*model.Rate, error) {
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
	// str := string(b)

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
