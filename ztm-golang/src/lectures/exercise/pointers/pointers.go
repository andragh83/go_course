//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (Active = true
Inactive = false)

type SecurityTag bool

type Item struct {
	name string
	securityTag SecurityTag
}

func deactivateTag(securityTag *SecurityTag) {
	*securityTag = Inactive
}

func activateTag(securityTag *SecurityTag) {
	*securityTag = Active
}

func checkout(slice []Item) {
	for i := range slice {
		deactivateTag(&slice[i].securityTag)
	}
	// for i:=0; i<len(slice); i++ {
	// 	deactivateTag(&slice[i].securityTag)
	// }
}

func printSlice(slice []Item){
	fmt.Println()
	for _, item :=range slice{
		fmt.Println(item.name, item.securityTag)
	}
}

func main() {

	itemsSlice := []Item{
		{
			name: "item1",
			securityTag: Active,
		},
		{
			name: "item2",
			securityTag: Active,
		},
		{
			name: "item3",
			securityTag: Active,
		},
		{
			name: "item4",
			securityTag: Active,
		},

	} 
	printSlice(itemsSlice)
	deactivateTag(&itemsSlice[1].securityTag)
	printSlice(itemsSlice)
	checkout(itemsSlice)
	printSlice(itemsSlice)
	activateTag(&itemsSlice[3].securityTag)
	printSlice(itemsSlice)
}
