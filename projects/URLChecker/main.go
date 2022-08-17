package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequest = errors.New("error")
var errResults = make(map[string]string)

type results struct {
	url    string
	status string
}

func main() {

	c := make(chan results)

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
		fmt.Println(<-c)
	}

}

func hitUrl(url string, ch chan<- results) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- results{url: url, status: status}
}
