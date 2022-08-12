package main

import (
	"fmt"
	dict "simple/projects/dictionary/myDict"
)

func main() {
	dictionary := dict.Dictionary{}
	dictionary["sayHello"] = "Hello"
	fmt.Println(dictionary)
}
