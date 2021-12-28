package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type writeToScreen struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	wts := writeToScreen{}
	_, err = io.Copy(wts, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func (writeToScreen) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
