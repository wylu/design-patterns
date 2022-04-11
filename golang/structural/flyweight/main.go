package main

import "fmt"

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 0),
		counterTerrorists: make([]*player, 0),
	}
}

func (g *game) addTerrorist(dressType string) {
	terrorist := newPlayer("Terrorist", dressType)
	g.terrorists = append(g.terrorists, terrorist)
}

func (g *game) addCounterTerrorist(dressType string) {
	counterTerrorist := newPlayer("CounterTerrorist", dressType)
	g.counterTerrorists = append(g.counterTerrorists, counterTerrorist)
}

func main() {
	game := newGame()
	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
