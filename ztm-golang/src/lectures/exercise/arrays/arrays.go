//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price int
}

func printStats (shoppingList [4]Product) {
	var lastItem Product
	nbOfItems, cost := 0, 0
	
	for i:=0; i<len(shoppingList); i++ {
		item := shoppingList[i]
		if (item.name != "") {
			fmt.Println(item.name, "costs", item.price)
			lastItem = item
			nbOfItems++
			cost +=item.price
		}
	}
	fmt.Println("last item in the list is", lastItem.name)
	fmt.Println("nb of items", nbOfItems)
	fmt.Println("total cost is",cost)
}



func main() {
	myShoppingList := [4]Product{
		{name: "bread", price: 3},
		{name: "milk", price: 5},
		{name: "eggs", price: 7},
	}

	printStats(myShoppingList)

	myShoppingList[3].name = "bottledwater"
	myShoppingList[3].price = 4

	printStats(myShoppingList)
	
}
