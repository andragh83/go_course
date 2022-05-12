package main

import (
	"fmt"
)

type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {

	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2, Boarded: false}
		ella = Passenger{"Ella", 3, false}
	)
	fmt.Println(bill, ella)

	heidy := Passenger{}
	heidy.Name = "Heidy"
	heidy.TicketNumber = 4
	heidy.Boarded = false

	bill.Boarded = true
	casey.Boarded = true
	

	if (bill.Boarded) {
		fmt.Println("Bill has boarded the bus")
	}

	if (casey.Boarded) {
		fmt.Println("Casey has boarded the bus")
	}

	if (heidy.Boarded) {
		fmt.Println("Heidy has boarded the bus")
	} else {
		fmt.Println("Heidy hasn't boarded the bus")
	}

	heidy.Boarded = true

	fmt.Println(bill, ella, heidy)

	bus := Bus{FrontSeat: heidy}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name)

}
