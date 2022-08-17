package main

import (
	"errors"
	"net/http"
)

var errRequest = errors.New("error")

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://nastroyvse.ru/",
		"https://quizlet.com/",
		"https://app.milanote.com/",
		"https://reddit.com/",
	}

	for _, url := range urls {
		hitUrl(url)
	}

}

func hitUrl(url string) error {
	resp, err := http.Get(url)

	if err == nil || resp.StatusCode >= 400 {
		return errRequest
	}

	return nil
}
