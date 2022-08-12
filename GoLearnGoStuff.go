package main

import (
	"fmt"
)

func main() {
	// workWithPointers()
	// workWithArray()
	// workWithMap()
	workWithStruct()
}

func workWithPointers() {
	a := 2
	b := a
	a = 10
	fmt.Println(a, b) // diff adresses

	x := 4
	y := &x
	x = 100
	fmt.Println(x, *y) //same adresses  &- to take adress, *- to see throught address in memory(and updated the whole address)
	*y = 150
	fmt.Println(x, *y)
}

func workWithArray() {
	arr1 := [5]string{"123", "name", "blablabla"}
	arr1[3] = "something"
	fmt.Println(arr1)

	arr2 := []string{"123", "name", "blablabla"}
	// arr2[3] = "something" //exception out of bound
	arr2 = append(arr2, "newValue")
	fmt.Println(arr2)
}

func workWithMap() {
	nico := map[string]string{"name1": "alex", "name2": "someName"}

	for index, value := range nico {
		fmt.Println(index, value)
	}

}

func workWithStruct() {
	type user struct {
		id       int
		mail     string
		username string
		password string
		age      byte
		roles    []string
	}

	roles1 := []string{"user", "admin"}
	name1 := user{id: 100, mail: "name1@naver.com", username: "nameNick", password: "123123", age: 20, roles: roles1}
	fmt.Println(name1)
	fmt.Println(name1.mail)

}
