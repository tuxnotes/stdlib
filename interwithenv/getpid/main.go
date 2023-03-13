package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	pid := os.Getpid()
	fmt.Printf("Process id: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	output, err := prc.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
