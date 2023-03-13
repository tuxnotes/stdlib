package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("vim-go")
	// connStr := os.Getenv("DB_CONN")
	connStr, ex := os.LookupEnv("DD_CONN")
	if !ex {
		fmt.Printf("env var DB_CONN not found")
	} else {
		fmt.Printf("Connection string: %s\n", connStr)
	}
}
