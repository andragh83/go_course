//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func oneDiceRoll( sidesNb int ) int {
	return rand.Intn(sidesNb)+1
}


func main() {
	rand.Seed(time.Now().UnixNano())
	rollNb, diceNb, sidesNb, diceRollSum := 2, 2, 6, 0
	
	for i:=1; i<=rollNb; i++ {
		for j:=1; j<=diceNb; j++ {
			randomPick := oneDiceRoll(sidesNb)
			fmt.Println("dice", j, "is", randomPick)
			diceRollSum += randomPick
		}
		fmt.Println("sum at roll #", i, "sum is", diceRollSum)
	}

	switch sum:=diceRollSum; {
	case sum == 2 && diceNb == 2: fmt.Println("snake eyes")
	case sum == 7 : fmt.Println("lucky 7")
	case sum%2 == 0: fmt.Println(diceRollSum, "even")
	default: fmt.Println(diceRollSum,"odd")
	}
}

