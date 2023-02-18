package db

import (
	"fmt"
	"os"
)

func RunDb() {
	os.Remove(("database.db"))
	file, err := os.Create("database.db")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}
