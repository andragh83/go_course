//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
)

type Player struct {
	health, maxHealth uint
	energy, maxEnergy uint
	name string
}

func (player *Player) addHealth (healthUnits uint) {
	if player.health + healthUnits >= player.maxHealth { player.health = player.maxHealth } else {
		player.health = player.health + healthUnits
	} 
}

func (player *Player) sufferDamage (healthUnits uint) {
	if player.health <= healthUnits { player.health = 0; fmt.Println("you're dead") } else {
		player.health = player.health - healthUnits
	} 
}

func (player *Player) addEnergy (energyUnits uint) {
	if player.energy + energyUnits <= player.maxEnergy { player.energy = player.maxEnergy } else {
		player.energy = player.energy + energyUnits
	} 
}

func main() {

	player := Player{
		health: 20,
		maxHealth: 100,
		energy: 10,
		maxEnergy: 100,
		name: "Joe",
	}
	fmt.Println("Initial", player)
	player.addHealth(30)
	fmt.Println("Adding 30 health", player)
	player.addEnergy(80)
	fmt.Println("Adding 80 enegry", player)
	player.addHealth(80)
	fmt.Println("Adding 80 health", player)
	player.sufferDamage(120)
	fmt.Println("Last", player)

}
