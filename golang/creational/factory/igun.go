package main

type iGun interface {
	setName(name string)
	getName() string
	setPower(power int)
	getPower() int
}
