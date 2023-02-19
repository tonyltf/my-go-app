package init

import (
	"context"
	"fmt"
	config "my-go-app/configs"
	"my-go-app/internal/database"
	"my-go-app/internal/database/dal"
	"os"
)

func InitDb(ctx context.Context) {
	envConfig := config.InitConfig()
	dbFile := envConfig.DbConnection

	if _, err := os.Stat(dbFile); err == nil {
		fmt.Printf("Database File exists\n")
		// os.Remove(("database.db"))
	} else {
		file, err := os.Create(dbFile)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()

		db, err := database.Open(ctx, []string{dal.RateTable})
		if err != nil {
			fmt.Printf("InitDb error %v", err)
		}

		defer db.Close()
	}
}
