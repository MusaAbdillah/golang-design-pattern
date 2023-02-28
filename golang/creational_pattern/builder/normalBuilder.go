package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden door"
}

func (n *NormalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		doorType:   n.doorType,
		windowType: n.windowType,
		floor:      n.floor,
	}
}
