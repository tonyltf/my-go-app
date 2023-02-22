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
	myDal    *Dal
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

	myDal = &Dal{
		Ctx: context.Background(),
		Db:  db,
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Error when preparing test database: %v", err)
	}
}

func TestCreate(t *testing.T) {
	rate := &model.Rate{
		ID:           uuid.New(),
		CurrencyPair: "BTCUSD",
		ExchangeRate: 25000.0,
		CreatedAt:    time.Now(),
	}
	newRate, err := myDal.Create(rate)
	if err != nil {
		t.Errorf("Error when create new record: %v", err)
	}
	assert.Equal(t, rate, newRate)
}

func TestRead(t *testing.T) {
	prepareTestDatabase()
	createdAt, err := time.Parse("2006-01-02 15:04:05", "2023-02-20 00:00:00")
	if err != nil {
		t.Errorf("Error when setting CreatedAt: %v", err)
	}
	rate, err := myDal.Read("BTCUSD", &createdAt)
	if err != nil {
		t.Errorf("Error when reading database: %v", err)
	}
	assert.NotNil(t, rate)
	assert.Equal(t, 25000.0, rate.ExchangeRate)
}

// Test all 2 testdata sets
func TestReadRange(t *testing.T) {
	prepareTestDatabase()
	fromCreatedAt, err := time.Parse("2006-01-02 15:04:05", "2023-02-19 00:00:00")
	if err != nil {
		t.Errorf("Error when setting CreatedAt: %v", err)
	}
	toCreatedAt, err := time.Parse("2006-01-02 15:04:05", "2023-03-01 00:00:00")
	if err != nil {
		t.Errorf("Error when setting CreatedAt: %v", err)
	}
	rate, err := myDal.ReadRange("BTCUSD", fromCreatedAt, toCreatedAt)
	if err != nil {
		t.Errorf("Error when reading database: %v", err)
	}
	assert.NotNil(t, rate)
	assert.Equal(t, 26000.0, rate.ExchangeRate)
}

// Test only 1 testdata set
func TestReadRange2(t *testing.T) {
	prepareTestDatabase()
	fromCreatedAt, err := time.Parse("2006-01-02 15:04:05", "2023-02-20 00:00:00")
	if err != nil {
		t.Errorf("Error when setting CreatedAt: %v", err)
	}
	toCreatedAt, err := time.Parse("2006-01-02 15:04:05", "2023-03-01 00:00:00")
	if err != nil {
		t.Errorf("Error when setting CreatedAt: %v", err)
	}
	rate, err := myDal.ReadRange("BTCUSD", fromCreatedAt, toCreatedAt)
	if err != nil {
		t.Errorf("Error when reading database: %v", err)
	}
	assert.NotNil(t, rate)
	assert.Equal(t, 27000.0, rate.ExchangeRate)
}
