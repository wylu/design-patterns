package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
