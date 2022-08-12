package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World")
	// something.SayGoodbye()
	// something.SayHello()

	// const isTrue bool = true
	// const numInt32 int32 = 2147483647
	// const numInt int = 50
	// const numShort int16 = 32767
	// const numByte byte = 255
	// const name = "Nick"
	// someNum := 15 //same as:  var name int32 = 15
	// someNum = 40
	// fmt.Println(someNum) //variables

	// fmt.Println(multiply(someNum, 2)) //func
	// totalLength, toUpperCase := toUpper(name) //func
	// fmt.Println(totalLength, toUpperCase) //func
	// fmt.Println(toUpper2(name)) //func
	// repeatMe("123", "234", "345") // func

	// fmt.Println(superAdd(1, 2, 3, 4, 5, 6)) cycles

	// fmt.Println(canIDrinkIf(17))  if/else
	// fmt.Println(canIDrinkSwitch(17)) switch/case
}

func multiply(a, b int) int {
	return a * b
}

func toUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func toUpper2(name string) (length int, upperCase string) {
	defer fmt.Println("im done")
	length = len(name)
	upperCase = strings.ToUpper(name)
	return //naked return (Go already understands what you mean here)
}

func repeatMe(words ...string) { //multiple arguments
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// for number := range numbers { //index counting
	// for index, number := range numbers {  //forEach
	// 	fmt.Println(index, number)
	// }

	// for _, number := range numbers { // _ means ignore
	// 	fmt.Println(number)
	// }

	// for i := 0; i < len(numbers); i++ { //typical for
	// 	fmt.Println(numbers[i])
	// }

	total := 0

	for _, number := range numbers {
		total += number
		fmt.Println(total)
	}

	return total
}

func canIDrinkIf(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { //creating variable inside
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return false
}
