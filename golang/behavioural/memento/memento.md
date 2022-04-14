# Memento Design Pattern

[Memento Design Pattern in Go (Golang)](https://golangbyexample.com/memento-design-pattern-go/)

## Introduction

Memento design pattern is a behavioral design pattern. It allows us to save checkpoints for an object and thus allow an object to revert to its previous state. Basically it helps in undo-redo operations on an object. Below are the design components of the Memento Design Pattern.

- **Originator**: It is the actual object whose state is saved as a memento.
- **Memento**: This is the object which saves the state of the originator
- **Caretaker**: This is the object that saves multiple mementos. Given an index, it returns the corresponding memento.

The originator defines two methods. **savememento()** and **restorememento()**

- **savememento()-** in this method the originator saves its internal state into a memento object.
- **restorememento()-** this method takes input as a memento object. The originator restores itself to the pass memento. Hence a previous state is restored.

## Practical Example

**originator.go**

```go
package main

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}
```

**memento.go**

```go
package main

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}
```

**caretaker.go**

Notice that the caretaker contains the mementoArray which holds all the memento.

```go
package main

type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}
```

**main.go**

```go
package main

import "fmt"

func main() {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())

	caretaker.addMemento(originator.createMemento())
	originator.setState("C")

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
```

**Output:**

```none
originator Current State: A
originator Current State: B
originator Current State: C
Restored to State: B
Restored to State: A
```

## Full Working Code

```go
package main

import "fmt"

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}

func main() {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())

	caretaker.addMemento(originator.createMemento())
	originator.setState("C")

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
```

**Output:**

```none
originator Current State: A
originator Current State: B
originator Current State: C
Restored to State: B
Restored to State: A
```
