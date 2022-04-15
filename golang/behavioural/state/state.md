# State Design pattern

[State Design pattern in Go](https://golangbyexample.com/state-design-pattern-go/)

## Definition

State design pattern is a behavioral design pattern that is based on Finite State Machine. We will explain the State Design Pattern in the context of an example of a Vending Machine. For simplicity, let’s assume that vending machine only has one type of item or product. Also for simplicity lets assume that a Vending Machine can be in 4 different states

1. **hasItem**
2. **noItem**
3. **itemRequested**
4. **hasMoney**

A vending machine will also have different actions. Again for simplicity lets assume that there are only four actions:

1. **Select the item**
2. **Add the item**
3. **Insert Money**
4. **Dispense Item**

## When To Use

- Use the State design pattern when the object can be in many different states. Depending upon current request the object needs to change its current state
  - In the above example, Vending Machine can be in many different states. A Vending Machine will move from one state to another. Let’s say Vending Machine is in **itemRequested** then it will move to **hasMoney** state once the action “Insert Money” is done
- Use when an object will have different responses to the same request depending upon the current state. Using state design pattern here will prevent a lot of conditional statements
  - For example in the case of Vending Machine, if a user is want to purchase an item then the machine will proceed if it is **hasItemState** or it will reject if it is in **noItemState**. If you notice here that the Vending Machine on the request of purchase of an item gives two different responses depending upon whether it is in **hasItemState** or **noItemState.**  Do notice the vendingmachine.go file below, it doesn’t have any kind of conditional statement. All the logic is being handled by concrete state implementations.

## UML Diagram

![State-Design-Pattern-1](images/State-Design-Pattern-1.webp)

## Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in code.

| Role             | File                  |
| ---------------- | --------------------- |
| Context          | vendingmachine.go     |
| State Interface  | state.go              |
| Concrete State 1 | noitemstate.go        |
| Concrete State 2 | hasitemstate.go       |
| Concrete State 3 | itemrequestedstate.go |
| Concrete State 4 | hasmoneystate.go      |

## Explanation

- We have an interface “State” which defines signatures of functions that represents action in the context of Vending Machine. Below are the actions function signatures
  1. addItem(int) error
  2. requestItem() error
  3. insertMoney(money int) error
  4. dispenseItem() error
- Each of the concrete state implementations implements all 4 above function and either move to another state on these actions or gives some response.
- Each of the concrete state also embeds a pointer to current Vending Machine object so that state transition can happen on that object.

Now lets look at code

## Practical Example

**vendingmachine.go**

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    vendingMachine := newVendingMachine(1, 10)
    err := vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()
    err = vendingMachine.addItem(2)
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()
    err = vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
}
```

**state.go**

```go
package main

type state interface {
    addItem(int) error
    requestItem() error
    insertMoney(money int) error
    dispenseItem() error
}
```

**noitemstate.go**

```go
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
```

**hasitemstate.go**

```go
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
```

**itemrequestedstate.go**

```go
package main

import "fmt"

type itemRequestedState struct {
    vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
    return fmt.Errorf("Can not requestItem when at itemRequestedState")
}

func (i *itemRequestedState) addItem(count int) error {
    return fmt.Errorf("Can not addItem when at itemRequestedState")
}

func (i *itemRequestedState) insertMoney(money int) error {
    if money < i.vendingMachine.itemPrice {
        return fmt.Errorf("Inserted money is less. Please insert %d",
            i.vendingMachine.itemPrice)
    }
    fmt.Println("Money entered is ok")
    i.vendingMachine.setState(i.vendingMachine.hasMoney)
    return nil
}

func (i *itemRequestedState) dispenseItem() error {
    return fmt.Errorf("Can not dispenseItem when at itemRequestedState")
}
```

**hasmoneystate.go**

```go
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
```

**main.go**

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    vendingMachine := newVendingMachine(1, 10)
    err := vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()
    err = vendingMachine.addItem(2)
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()

    err = vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
}
```

**Output:**

```go
Item requestd
Money entered is ok
Dispensing Item

Adding 2 items

Item requestd
Money entered is ok
Dispensing Item
```
