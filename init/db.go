package init

import (
	"fmt"
	"os"
)

func InitDb() {
	os.Remove(("database.db"))
	file, err := os.Create("database.db")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}
