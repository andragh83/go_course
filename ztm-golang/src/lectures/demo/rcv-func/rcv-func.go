package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func ocupySpace(lot *ParkingLot, spaceNb int) {
	lot.spaces[spaceNb-1].occupied = true
}

func (lot *ParkingLot) ocupySpace2(spaceNb int) {
	lot.spaces[spaceNb-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNb int) {
	lot.spaces[spaceNb-1].occupied = false
}

func main() {

lot:=ParkingLot{spaces: make([]Space, 4)}
fmt.Println("Initial",lot)

lot.ocupySpace2(1)
fmt.Println("after first", lot)
ocupySpace(&lot, 2)
fmt.Println("after second occupy", lot)
lot.vacateSpace(2)
fmt.Println("After vacancy",lot)

}
