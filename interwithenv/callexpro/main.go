package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("vim-go")
	// process execute result will
	// output to null device
	out := bytes.NewBuffer([]byte{})
	prc := exec.Command("ls")
	prc.Stdout = out
	// err := prc.Run()
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}
	prc.Wait()
	if prc.ProcessState.Success() {
		fmt.Println(out.String())
	}
}
