package main

type nornamlBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *nornamlBuilder {
	return &nornamlBuilder{}
}

func (b *nornamlBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *nornamlBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *nornamlBuilder) setNumFloor() {
	b.floor = 2
}

func (b *nornamlBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
