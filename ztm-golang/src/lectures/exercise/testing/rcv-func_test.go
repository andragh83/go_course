//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestHealth (t *testing.T) {
	player := Player{
		health: 20,
		maxHealth: 100,
		energy: 10,
		maxEnergy: 100,
		name: "Joe",
	}

	player.addHealth(90)

	if player.health > player.maxHealth {
		t.Errorf("health should not exceed %v, got %v", player.maxHealth, player.health)
	}

	player.sufferDamage(player.maxHealth+1)

	if player.health != 0 {
		t.Errorf("health should not go below 0, got %v", player.health)
	}
}

func TestEnergy (t *testing.T) {
	player := Player{
		health: 20,
		maxHealth: 100,
		energy: 10,
		maxEnergy: 100,
		name: "Joe",
	}

	player.addEnergy(90)
	if player.energy > player.maxEnergy {
		t.Errorf("energy should not exceed %v, got %v", player.maxEnergy, player.energy)
	}
}