package main

import "fmt"

func main() {
	var myName = "Lucin"
	fmt.Println("My name is", myName)

	var name string = "Barto"
	fmt.Println("Name =", name)

	userName := "Admin"
	fmt.Println("Username:", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("Part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("Part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("Lesson Name =", lessonName)
	fmt.Println("Lesson Type =", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}