package main

import (
	"fmt"
	"something"
	"strings"
)

func main() {
	fmt.Println("Hello, World")
	something.SayGoodbye()
	something.SayHello()

	const isTrue bool = true
	const numInt32 int32 = 2147483647
	const numInt int = 50
	const numShort int16 = 32767
	const numByte byte = 255
	const name = "Nick"

	someNum := 15 //same as:  var name int32 = 15
	someNum = 40
	fmt.Println(someNum)
	fmt.Println(multiply(someNum, 2))

	totalLength, toUpperCase := toUpper(name)
	fmt.Println(totalLength, toUpperCase)

	repeatMe("enmgfklewmfg", "fdkwqeof", "dwikdfm")

	fmt.Println(toUpper2(name))
}

func multiply(a, b int) int {
	return a * b
}

func toUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func toUpper2(name string) (length int, upperCase string) {
	length := name.len()
	upperCase := strings.ToUpper(name)
	return
}
