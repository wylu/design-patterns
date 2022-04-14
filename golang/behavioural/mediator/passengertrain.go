package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) requestArrival() {
	if p.mediator.canLand(p) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (p *passengerTrain) departure() {
	fmt.Println("PassengerTrain: Leaving")
	p.mediator.notifyFree()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}
