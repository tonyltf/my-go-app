package main

import (
	"my-go-app/api/v1/router"
	config "my-go-app/init"
	db "my-go-app/init"
	cron "my-go-app/internal"
)

//	@title			My Go App API
//	@version		1.0
//	@description	This is a simple app when learning Go

//	@contact.name	Tony Li
//	@contact.url	https://github.com/tonyltf
//	@contact.email	tingfung.tony@gmail.com

//	@BasePath	/v1

func main() {
	go config.InitConfig()
	go db.InitDb()
	go cron.InitCron()
	router.InitRouter()
}
