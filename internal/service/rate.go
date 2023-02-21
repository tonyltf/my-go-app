package service

import (
	"context"
	"encoding/json"
	"fmt"
	"my-go-app/internal/database"
	"my-go-app/internal/database/dal"
	"strconv"
	"time"
)

func GetRate(ctx context.Context, currencyPair string, timestamp *string) (*string, error) {

	db, err := database.Open(ctx, nil)
	if err != nil {
		fmt.Printf("Error in db %v\n", err)
		return nil, err
	}
	if currencyPair == "" {
		return nil, fmt.Errorf("Missing exchange pair")
	}

	pair := currencyPair[0:6]
	base := pair[0:3]
	target := pair[3:]
	var i int64
	var timestampValue time.Time
	var timestampPointer *time.Time
	timestampPointer = nil
	fmt.Printf("From %s - %s at %s\n", base, target, *timestamp)
	if *timestamp != "" {
		i, err = strconv.ParseInt(*timestamp, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		timestampValue = time.Unix(i, 0)
		timestampPointer = &timestampValue
	}
	rate, err := dal.Read(ctx, db, base+target, timestampPointer)
	if err != nil {
		fmt.Printf("Error in reading rate %v\n", err)
		return nil, fmt.Errorf("Error in reading rate %v\n", err)
	}
	if rate == nil {
		// TODO backup plan for fallback lookup
		return nil, fmt.Errorf("Exchange rate is missing")
	}
	fmt.Println(rate.ExchangeRate)
	if rate.ExchangeRate == 0 {
		return nil, fmt.Errorf("Exchange rate is missing")
	}

	jsonStr, err := json.Marshal(rate)
	if err != nil {
		fmt.Printf("Error in reading rate %v\n", err)
		return nil, err
	}
	str := string(jsonStr)
	return &str, nil
}
