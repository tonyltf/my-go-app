package fetcher

import (
	"fmt"
	config "my-go-app/init"
	"net/http"
	"strings"
)

func FetchRate(base string, target string) (*string, error) {
	envConfig := config.InitConfig()
	apiSource := envConfig.ApiSource
	apiSource = strings.Replace(apiSource, "{base}", strings.ToLower(base), 1)
	apiSource = strings.Replace(apiSource, "{target}", strings.ToLower(target), 1)
	fmt.Println(apiSource)
	response, err := http.Get(apiSource)
	if err != nil {
		return nil, fmt.Errorf("API error %w", err)
	}

	if response.StatusCode >= 400 {
		fmt.Println("Http status: " + response.Status)
	}

	str := fmt.Sprintf("%s", response.Body)
	return &str, nil

}
