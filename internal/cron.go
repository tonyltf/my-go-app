package cron

import (
	"context"
	"database/sql"
	"fmt"
	config "my-go-app/configs"
	"my-go-app/internal/database"
	"my-go-app/internal/database/dal"
	"my-go-app/internal/fetcher"

	"github.com/jasonlvhit/gocron"
)

func task(ctx context.Context, db *sql.DB) {
	fmt.Println("I am running task.")
	apiSource := config.InitConfig().ApiSource
	myFetcher, err := fetcher.NewFetcher(apiSource)
	if err != nil {
		fmt.Printf("Error in cron job (NewFetcher): %v\n", err)
		return
	}
	err = myFetcher.FetchRate("BTC", "USD")
	if err != nil {
		fmt.Printf("Error in cron job (FetchRate): %v\n", err)
		return
	}
	rate, err := myFetcher.TransformRate()
	if err != nil {
		fmt.Printf("Error in cron job (TransformRate): %v\n", err)
		return
	}
	r, err := dal.Create(ctx, db, rate)
	if err != nil {
		fmt.Printf("Insert rate error %v\n", err)
	}
	fmt.Println(r)
}

func InitCron(ctx context.Context) {
	// gocron.Every(1).Second().Do(task)
	db, err := database.Open(ctx, nil)
	if err != nil {
		fmt.Printf("InitCron open database error %v", err)
		return
	}
	task(ctx, db)
	gocron.Every(60).Seconds().Do(func() { task(ctx, db) })
	<-gocron.Start()
}
