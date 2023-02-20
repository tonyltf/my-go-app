package fetcher

import (
	"fmt"
	"io"
	"my-go-app/internal/database/model"
	"net/http"
	"strings"
)

type IFetcher interface {
	FetchRate(base string, target string) error
	TransformRate() (*model.Rate, error)
}

type fetchRateFunc func(base string, target string) error
type transformRateFunc func(*Fetcher) (*model.Rate, error)

type Fetcher struct {
	apiPath         string
	myTransformFunc transformRateFunc
	myFetchFunc     fetchRateFunc
	myJson          []byte
	myRate          *model.Rate
}

func (f *Fetcher) FetchRate(base string, target string) error {
	path := f.apiPath
	path = strings.Replace(path, "{base}", strings.ToLower(base), 1)
	path = strings.Replace(path, "{target}", strings.ToLower(target), 1)
	fmt.Println(path)

	if path != "" {
		res, err := http.Get(path)
		if err != nil {
			return fmt.Errorf("API error %w", err)
			// return nil, fmt.Errorf("API error %w", err)
		}

		if res.StatusCode >= 400 {
			return fmt.Errorf("Http status: " + res.Status)
			// return nil, fmt.Errorf("Http status: " + res.Status)
		}

		b, err := io.ReadAll(res.Body)
		f.myJson = b
		return nil
	}
	if f.myFetchFunc != nil {
		return f.myFetchFunc(base, target)
	}
	return fmt.Errorf("Missing API path")
}

func (f *Fetcher) TransformRate() (*model.Rate, error) {
	return f.myTransformFunc(f)
}
