package service

import (
	"context"
	"encoding/json"
	"fmt"
	"my-go-app/internal/database"
	"my-go-app/internal/database/dal"
)

func GetLastRate(ctx context.Context, currencyPair string) *string {

	db, err := database.Open(ctx, nil)
	if err != nil {
		fmt.Printf("Error in db %v\n", err)
		return nil
	}
	if currencyPair == "" {
		return nil
	}

	fmt.Println(currencyPair)
	pair := currencyPair[0:6]
	base := pair[0:3]
	target := pair[3:]
	fmt.Printf("From %s - %s\n", base, target)
	rate, err := dal.Read(ctx, db, base+target, nil)
	if err != nil {
		fmt.Printf("Error in reading rate %v\n", err)
		return nil
	}
	fmt.Println(rate)
	if rate == nil {
		return nil
	}

	jsonStr, err := json.Marshal(rate)
	if err != nil {
		fmt.Printf("Error in reading rate %v\n", err)
		return nil
	}
	str := string(jsonStr)
	return &str
}
