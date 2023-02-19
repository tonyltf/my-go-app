package database

import (
	"context"
	"database/sql"
	"fmt"
	config "my-go-app/configs"

	_ "github.com/mattn/go-sqlite3"
)

// Reference: https://github.com/g8rswimmer/go-data-access-example
func Open(ctx context.Context, stmts []string) (*sql.DB, error) {
	envConfig := config.InitConfig()
	dbSource := envConfig.DbConnection
	dbDrive := envConfig.DbDriver
	db, err := sql.Open(dbDrive, dbSource)
	if err != nil {
		return nil, fmt.Errorf("sqlite database open error %w", err)
	}

	for _, stmt := range stmts {
		if _, err := db.ExecContext(ctx, stmt); err != nil {
			db.Close()
			return nil, fmt.Errorf("sqlite database statment (%s) error %w", stmt, err)
		}
	}
	return db, nil
}
