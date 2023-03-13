package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ArrayValue []string

func (a *ArrayValue) String() string {
	return fmt.Sprintf("%v", *a)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {
	retry := flag.Int("retry", -1, "count for retry")
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "prefix string for logging")
	var arr ArrayValue
	flag.Var(&arr, "array", "slice for string")
	flag.Parse()
	retryCount := 0
	logger := log.New(os.Stdout, logPrefix, log.Ldate)
	for retryCount < *retry {
		logger.Println("Retry connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}
