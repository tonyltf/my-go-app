package dal

import (
	"context"
	"database/sql"
	"fmt"
	"my-go-app/internal/database/model"
	"os"
	"testing"
	"time"

	testfixtures "github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	if _, err = os.Stat("../../../database_test.db"); err == nil {
		err := os.Remove("../../../database_test.db")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		file, err := os.Create("../../../database_test.db")
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}

	db, err = sql.Open("sqlite3", "database_test.db")
	if err != nil {
		fmt.Println(err)
	}

	if _, err := db.Exec(RateTable); err != nil {
		db.Close()
		fmt.Println("sqlite database statment error %w", err)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("sqlite3"),
		testfixtures.Directory("../../../test/testdata"),
	)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	rate := &model.Rate{
		ID:           uuid.New(),
		CurrencyPair: "BTCUSD",
		ExchangeRate: 25000,
		CreatedAt:    time.Now(),
	}
	newRate, err := Create(context.Background(), db, rate)
	if err != nil {
		fmt.Println()
		t.Errorf("Error when create new record: %v", err)
	}
	assert.Equal(t, rate, newRate)
}

func TestRead(t *testing.T) {
	created_at, err := time.Parse("2006-01-02 15:04:05", "2023-02-20 00:00:00")
	if err != nil {
		t.Errorf("Error when setting CreatedAt: %v", err)
	}
	// r := &model.Rate{
	// 	CurrencyPair: "BTCUSD",
	// 	CreatedAt:    created_at,
	// }
	rate, err := Read(context.Background(), db, "BTCUSD", &created_at)
	if err != nil {
		t.Errorf("Error when reading database: %v", err)
	}
	assert.NotNil(t, rate)
}
