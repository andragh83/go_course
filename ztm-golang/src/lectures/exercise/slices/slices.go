//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(title string, slice []Part) {
	fmt.Println()
	fmt.Println(title)
	for i:=0; i<len(slice); i++ {
		el:=slice[i]
		fmt.Println(el)
	}
}

func main() {
	assemblyLine := []Part{"part1", "part2", "part3"}
	// assemblyLine := make([]Part, 3)
	// asign each part
	printAssemblyLine("First step", assemblyLine)
	additionals := []Part{"part4", "part5"}
	assemblyLine = append(assemblyLine, additionals...)
	printAssemblyLine("Second step", assemblyLine)
	assemblyLine = assemblyLine[3:]
	printAssemblyLine("Third step", assemblyLine)
}
