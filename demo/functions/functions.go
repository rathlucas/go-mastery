package main

import "fmt"


func double(x int) int {
	return x + x
}

func add(arg1, arg2 int) int {
	return arg1 + arg2
}

func greet() {
	fmt.Println("Hello, greeeeeet!!")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("Dozen =", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("Baker's dozen =", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another baker's dozen", anotherBakersDozen)
}