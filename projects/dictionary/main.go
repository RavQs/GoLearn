package main

import (
	"fmt"
	dict "simple/projects/dictionary/myDict"
)

func main() {
	dictionary := dict.Dictionary{"first": "firstNum"}

	_, err := dictionary.Search("first")
	err1 := dictionary.Add("first", "firstNum")
	_, err2 := dictionary.Search("second")
	err3 := dictionary.Add("second", "secondNum")
	_, err4 := dictionary.Search("second")
	err5 := dictionary.Update("second", "fourth")
	fmt.Println(dictionary)
	err6 := dictionary.Delete("first")
	check(err, err1, err2, err3, err4, err5, err6)
	fmt.Println(dictionary)
}

func check(err ...error) {
	if err != nil {
		fmt.Println(err, ",")
	}
}
