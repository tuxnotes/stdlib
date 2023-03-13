package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("vim-go")
	pwd, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", pwd)

	shorter := filepath.Dir(pwd)
	fmt.Printf("cut path: %s\n", shorter)
	realPath, err := filepath.EvalSymlinks(pwd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("real path: %s\n", realPath)
}
