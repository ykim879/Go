package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(resp.Status)
			bytes, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading:%s\n", err)
				os.Exit(1)
			}
			fmt.Println(string(bytes))
		} else {
			fmt.Println("Invalid url. Please enter url starts with http://")
		}
	}
}
