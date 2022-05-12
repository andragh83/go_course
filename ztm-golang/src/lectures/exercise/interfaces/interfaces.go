//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type VehicleDirecter interface {
	DirectVehicles() Lift
}

type Motorcycles string
type Cars string
type Trucks string

func (m Motorcycles) DirectVehicles() Lift {
	return SmallLift
}

func (c Cars) DirectVehicles() Lift {
	return StandardLift
}

func (t Trucks) DirectVehicles() Lift {
	return LargeLift
}

func sendToLift(v VehicleDirecter) {
	switch v.DirectVehicles() {
		case SmallLift: fmt.Println("direct to small lift")
		case StandardLift: fmt.Println("direct to standards lift")
		case LargeLift: fmt.Println("direct to large lift")
		}
}

func DirectAllVehicles(vehicles[] VehicleDirecter) {
	for _, vehicle := range(vehicles) {
		fmt.Printf("where to direct %v ?\n", vehicle)
		sendToLift(vehicle)
	}
}

func main() {
	vehicles := []VehicleDirecter{Motorcycles("BMW motocycle"), Trucks("Mercedes truck"), Cars("Mini cooper")}
	DirectAllVehicles(vehicles)
}
