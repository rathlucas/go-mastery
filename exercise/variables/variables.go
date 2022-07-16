package main

import "fmt"

func main() {
	var myFavoriteColor = "Red"
	fmt.Println("My favorite color is", myFavoriteColor)

	birthYear, myAgeInYears := 2000, 21
	fmt.Printf("I was born in the year %v, and i'm %v years old\n", birthYear, myAgeInYears)

	var (
		firstInitial = "L"
		lastInitial = "S"
	)
	fmt.Printf("My first initial is %v, and my last initial is %v\n", firstInitial, lastInitial)

	var myAgeInDays = myAgeInYears * 365

	fmt.Println("My age in days is", myAgeInDays)
}