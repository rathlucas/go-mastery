package main

import (
	"fmt"
)


func greeting(name string) {
	fmt.Println("Greetings", name, "!")
}

func returnMessage() string {
	return "This function returns a message !"
}

func addThreeNums(a, b, c int) int {
	return a + b + c
}

func returnsTen() int {
	return 10
}

func returnsTwoAndThree() (int, int) {
	return 2, 3
}

func main() {
	greeting("Lucas")

	fmt.Println(returnMessage())

	fmt.Println("addThreeNums returned:", addThreeNums(1, 2, 3))

	a, b := returnsTwoAndThree()
	result := addThreeNums(returnsTen(), a, b)

	fmt.Println("The result is", result)
}