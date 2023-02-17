package main

import (
	"my-go-app/app/exchange-rate/api/v1/router"
	"my-go-app/app/exchange-rate/cron"
	"my-go-app/build/db"
)

//	@title			My Go App API
//	@version		1.0
//	@description	This is a simple app when learning Go

//	@contact.name	Tony Li
//	@contact.url	https://github.com/tonyltf
//	@contact.email	tingfung.tony@gmail.com

//	@BasePath	/v1

func main() {
	go db.RunDb()
	go cron.RunCron()
	router.RunRouter()
}
