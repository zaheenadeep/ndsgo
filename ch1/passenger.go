package main

import "fmt"

type Passenger struct {
	Name		string
	Address		string
	PhoneNumber	int
	SeatNumber	int
	Smoking		bool
	Destination	string
	SameFlight	*Passenger
	SameDest	*Passenger
}

func AddPassenger(firstFlight, firstDest **Passenger) *Passenger {
	p := new(Passenger)
	p.SameFlight, p.SameDest = *firstFlight, *firstDest
	*firstFlight, *firstDest = p, p
	return p
}

func CountSameDestination(firstDest *Passenger) int {
	count := 0
	for p := firstDest; p != nil; p = p.SameDest {
		count++
	}
	return count
}

var (
	first451	*Passenger
	firstDallas *Passenger
)

func main() {
	p := AddPassenger(&first451, &firstDallas)
	p.Name = "Hanan"
	p = AddPassenger(&first451, &firstDallas)
	firstDallas.Name = "Samet"
	AddPassenger(&first451, &firstDallas)
	fmt.Println(CountSameDestination(firstDallas))
}