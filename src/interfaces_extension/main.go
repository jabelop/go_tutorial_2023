package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

// implements the Write method from the Writter interface which is called from io.copy(), so the console print behavior is overriden
func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// without overriding behavior
	//io.Copy(os.Stdout, resp.Body)
	fmt.Println("===================")
	// overriding behavior, must comment previous io.Copy line for working
	writer := customWriter{}
	io.Copy(writer, resp.Body)
	fmt.Println()
}
