package init

import (
	"context"
	"fmt"
	config "my-go-app/configs"
	"my-go-app/internal/database"
	"my-go-app/internal/database/dal"
)

func InitDb(ctx context.Context) {
	envConfig := config.InitConfig()
	dbFile := envConfig.DbConnection
	fmt.Println(dbFile)

	db, err := database.Open(ctx, []string{dal.RateTable})
	if err != nil {
		fmt.Printf("InitDb error %v", err)
	}

	defer db.Close()
}
