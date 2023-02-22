package service

import (
	"context"
	"encoding/json"
	"fmt"
	"my-go-app/internal/database"
	"my-go-app/internal/database/dal"
	"my-go-app/internal/messages"
	"strconv"
	"time"
)

func GetRate(ctx context.Context, currencyPair string, timestamp string) (*string, error) {

	db, err := database.Open(ctx, nil)
	defer db.Close()
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
	if timestamp != "" {
		fmt.Printf("From %s - %s at %s\n", base, target, timestamp)
		i, err = strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		timestampValue = time.Unix(i, 0)
		timestampPointer = &timestampValue
	}
	myDal := &dal.Dal{
		Ctx: ctx,
		Db:  db,
	}
	rate, err := myDal.Read(base+target, timestampPointer)
	if err != nil {
		fmt.Printf("Error in reading rate %v\n", err)
		return nil, fmt.Errorf("Error in reading rate %v\n", err)
	}
	if rate == nil {
		// TODO backup plan for fallback lookup
		return nil, fmt.Errorf(messages.Error.MISSING_EXCHANGE_RATE)
	}
	if rate.ExchangeRate == 0 {
		return nil, fmt.Errorf(messages.Error.MISSING_EXCHANGE_RATE)
	}

	jsonStr, err := json.Marshal(rate)
	if err != nil {
		fmt.Printf("Error in reading rate %v\n", err)
		return nil, err
	}
	str := string(jsonStr)
	return &str, nil
}

func GetAvgRate(ctx context.Context, currencyPair string, fromTimestring string, toTimestring string) (*string, error) {

	db, err := database.Open(ctx, nil)
	defer db.Close()
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
	if fromTimestring != "" && toTimestring != "" {
		fmt.Printf("From %s - %s from %s - %s\n", base, target, fromTimestring, toTimestring)
		i, err := strconv.ParseInt(fromTimestring, 10, 64)
		j, err := strconv.ParseInt(toTimestring, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fromTimestamp := time.Unix(i, 0)
		toTimestamp := time.Unix(j, 0)
		fmt.Println(fromTimestring)
		fmt.Println(toTimestring)
		fmt.Println(fromTimestamp)
		fmt.Println(toTimestamp)
		myDal := &dal.Dal{
			Ctx: ctx,
			Db:  db,
		}
		rate, err := myDal.ReadRange(base+target, fromTimestamp, toTimestamp)
		if err != nil {
			if err.Error() == messages.Error.MISSING_EXCHANGE_RATE {
				return nil, err
			}
			return nil, fmt.Errorf("Error in Reading range: %v\n", err)
		}

		jsonStr, err := json.Marshal(rate)
		if err != nil {
			fmt.Printf("Error in reading rate %v\n", err)
			return nil, err
		}
		str := string(jsonStr)
		return &str, nil
	}

	return nil, nil
}
