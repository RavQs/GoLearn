package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequest = errors.New("error")
var errResults = make(map[string]string)

type resultRequest struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan resultRequest)

	urls := []string{
		"https://www.google.com/",
		"https://nastroyvse.ru/",
		"https://app.milanote.com/",
		"https://amazon.com/",
		"https://bbc.com/",
		"https://stackoverflow.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for key, value := range results {
		fmt.Println(key, value)
	}

}

func hitUrl(url string, ch chan<- resultRequest) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- resultRequest{url: url, status: status}
}
