package fetcher

import "fmt"

func NewFetcher(fetcherName string) (IFetcher, error) {
	if fetcherName == "cryptonator" {
		return NewCryptonator(), nil
	}
	if fetcherName == "random" {
		return NewRandom(), nil
	}

	return nil, fmt.Errorf("Fetcher undefined")
}
