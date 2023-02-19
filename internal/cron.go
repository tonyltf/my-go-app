package cron

import (
	"fmt"
	"my-go-app/internal/fetcher"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.")
	rate, err := fetcher.FetchRate("BTC", "USD")
	if err != nil {
		fmt.Printf("Error in cron job: %v\n", err)
	}

	fmt.Println(rate)
}

func InitCron() {
	// gocron.Every(1).Second().Do(task)
	task()
	gocron.Every(60).Seconds().Do(task)
	<-gocron.Start()
}
