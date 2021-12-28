package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected an argument for filename.")
		os.Exit(1)
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
