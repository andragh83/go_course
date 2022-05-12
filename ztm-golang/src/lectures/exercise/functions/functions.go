//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func name (name string) {
	fmt.Println("Hello there", name)
}

func returnAnyMesage()string  {
	return "This is a message"
}

func addThreeNumbers(a,b,c int) int {
	return a+b+c
}

func returnAnyNumber () int {
	return 23
}

func returnAnyTwoNumbers ()(int, int) {
	return 19, 83
}
 
func main() {
name("Andra")
fmt.Println(returnAnyMesage())
fmt.Println("Add 1, 2 and 3", addThreeNumbers(1,2,3))
fmt.Println(returnAnyNumber())
a, b := returnAnyTwoNumbers()
fmt.Println("two numbers multireturned are", a, b)
fmt.Println("Add three numbers together using any combination of the existing functions", addThreeNumbers(a, b, returnAnyNumber()))
}
