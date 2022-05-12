//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

type Coordinate struct {
	x,y int
}

type Rectangle struct {
	topLeft, bottomRight Coordinate
}

func width(rectangle Rectangle)int {
	return  rectangle.bottomRight.x-rectangle.topLeft.x
}

func length(rectangle Rectangle)int {
	return rectangle.topLeft.y - rectangle.bottomRight.y
}

func area (rectangle Rectangle) int {
	return width(rectangle) * length(rectangle)
}

func perimeter (rectangle Rectangle) int {
	return 2*width(rectangle) + 2*length(rectangle)
}

func printInfo (area int, perimeter int) {
	fmt.Println("Area is", area)
	fmt.Println("Perimeter is", perimeter)
}

func main() {
	rect := Rectangle{Coordinate{0,7},Coordinate{10,0}}
	printInfo(area(rect), perimeter(rect))
}
