package main

import (
	"fmt"
	dict "simple/projects/dictionary/myDict"
)

func main() {
	dictionary := dict.Dictionary{"first": "firstNum"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
