package main

import "fmt"

func main() {
	myFirstName := "Andra"
	fmt.Println("My first name is", myFirstName)

	var myName string = "Andra Maria"
	fmt.Println("My name is", myName)
	
	
	var sum int
	fmt.Println("Value of sum is", sum)

	part1, other := 1, 5
	fmt.Println("Part 1 is", part1, "Other is", other)

	part2, other := 2, 0
	fmt.Println("Part 2 is", part2, "Other is now", 0)

	var (
		lessonOne = "Variables"
		lessonType = "Demo"
		)
	fmt.Println("Lesson one is ", lessonOne, "Lesson type is ", lessonType)

	word1, word2, _ := "hello", "world", '!'
	fmt.Println("Print:", word1, word2)

}
