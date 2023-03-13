package main

import (
	"log"
	"runtime"
)

func main() {
	const info = `
	App %s is starting.
	The binary was build by %s.
	`
	log.Printf(info, "Example", runtime.Version())
}
