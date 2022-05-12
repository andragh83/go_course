package main

import "fmt"

func double(x int) int {
	return x*2
}

func add(lhs, rhs int) int {
	return lhs+rhs
}

func greet(){
	fmt.Println("Hello from greet function")
}

func main() {
	greet()
	dozen := double(10)
	fmt.Println("A dozen is", dozen)
	bakersDozen := add(dozen, 1)
	fmt.Println("Bakers dozen is", bakersDozen)
	anotherBakesDozen := add(double(10), 1)
	fmt.Println("Another bakers dozen is", anotherBakesDozen)

}
