package cron

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.")
}

func InitCron() {
	// gocron.Every(1).Second().Do(task)
	gocron.Every(60).Seconds().Do(task)
	<-gocron.Start()
}
