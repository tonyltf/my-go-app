package cron

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.")
}

func RunCron() {
	// gocron.Every(1).Second().Do(task)
	gocron.Every(1).Minute().Do(task)
	<-gocron.Start()
}
