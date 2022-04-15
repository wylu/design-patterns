package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h *hasMoneyState) requestItem() error {
	return fmt.Errorf("Can not requestItem when at hasMoneyState")
}

func (h *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Can not addItem when at hasMoneyState")
}

func (h *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Can not insertMoney when at hasMoneyState")
}

func (h *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}
