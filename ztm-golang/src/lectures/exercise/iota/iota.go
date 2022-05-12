//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operator int

const (
	Add Operator = iota+1
	Subtract 
	Multiply
	Divide
)

func (operator Operator) calculate (x float64, y float64) float64 {
	switch operator {
	case Add: return x+y
	case Subtract: return x-y
	case Multiply: return x*y
	case Divide: return x/y
	}
	panic ("Unhandled operation")
}

func main() {

	add := Add
	sub := Subtract
	mul := Multiply
	div := Divide

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
