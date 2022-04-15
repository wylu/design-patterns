package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h *hasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Can not insertMoney when at hasItemState")
}

func (h *hasItemState) dispenseItem() error {
	return fmt.Errorf("Can not dispenseItem when at hasItemState")
}
