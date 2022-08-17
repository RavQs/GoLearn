package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequest = errors.New("error")
var results = make(map[string]string)
var errResults = make(map[string]string)

// var somethingddos = "https://amazon.com/"
// var ddos [15]string

func main() {
	// for ind := range ddos {
	// 	ddos[ind] = somethingddos
	// }

	urls := []string{
		"https://www.google.com/",
		"https://nastroyvse.ru/",
		"https://quizlet.com/",
		"https://app.milanote.com/",
		"https://reddit.com/",
		"https://amazon.com/",
		"https://amazon.com/",
	}

	for _, url := range urls {
		err := hitUrl(url)
		result := "OK"

		if err != nil {
			result = "FAILED"
			errResults[url] = result
		} else {
			results[url] = result
		}

	}

	for url, result := range results {
		fmt.Println(url, result)
	}
	fmt.Println("")
	for url, result := range errResults {
		fmt.Println(url, result)
	}

	// values1 := make(map[int]string)
	// values2 := make(map[int]string)

	// go goCounter("Name1", values1)
	// goCounter("Name2", values2)

	// fmt.Println(values1, "\n", values2)

}

func hitUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return errRequest
	}

	return nil
}

func goCounter(value string, values map[int]string) {
	for i := 0; i < 10; i++ {
		fmt.Println(value)
		values[i] = value
		time.Sleep(time.Second)
	}
}
