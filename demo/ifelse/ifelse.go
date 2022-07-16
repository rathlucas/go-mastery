package main

import "fmt"


func average(a, b, c int) float32 {
	return float32(a + b + c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 10, 5, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 scored higher than Quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 scored higer than Quiz 1")
	} else {
		fmt.Println("The scores are equal")
	}

	if average(quiz1, quiz2, quiz3) > 5 {
		fmt.Println("Acceptable average!")
	} else {
		fmt.Println("Unacceptable average!")
	}
}