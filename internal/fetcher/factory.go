package fetcher

import "fmt"

func NewFetcher(fetcherName string) (IFetcher, error) {
	if fetcherName == "cryptonator" {
		return NewCryptonator(), nil
	}
	if fetcherName == "random" {
		return NewRandom(), nil
	}
	if fetcherName == "coinapi" {
		return NewCoinApi(), nil
	}
	if fetcherName == "kraken" {
		return NewKraken(), nil
	}

	return nil, fmt.Errorf("Fetcher undefined")
}
