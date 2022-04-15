package main

import (
	"fmt"
)

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n *noItemState) addItem(count int) error {
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Can not insertMoney when at noItemState")
}

func (n *noItemState) dispenseItem() error {
	return fmt.Errorf("Can not dispenseItem when at noItemState")
}
