# Mediator Design Pattern

[Mediator Design Pattern in Go (Golang)](https://golangbyexample.com/mediator-design-pattern-golang/)

## Introduction

Mediator design pattern is a behavioral design pattern. This pattern suggests creating a mediator object to prevent direct communication among objects so that direct dependencies between them is avoided.

One very good example of a mediator patter is the railway system platform. Two trains never communicate between themselves for the availability of the platform. The **stationManager** acts as a mediator and makes the platform available to only one of the trains. The train connects with **stationManager** and acts accordingly. It maintains a queue of waiting trains. In case of any train leaving a platform, it notifies one of the train to arrive on the platform next.

Notice how **stationManger** acts as a mediator between the **trains** and the **platform** in the code below.

- passengerTrain and goodsTrain implement the train interface.
- stationManger implements the mediator interface.

## Practical Example

**train.go**

```go
package main

type train interface {
    requestArrival()
    departure()
    permitArrival()
}
```

**passengertrain.go**

```go
package main

import "fmt"

type passengerTrain struct {
    mediator mediator
}

func (g *passengerTrain) requestArrival() {
    if g.mediator.canLand(g) {
        fmt.Println("PassengerTrain: Landing")
    } else {
        fmt.Println("PassengerTrain: Waiting")
    }
}

func (g *passengerTrain) departure() {
    fmt.Println("PassengerTrain: Leaving")
    g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
    fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}
```

**goodstrain.go**

```go
package main

import "fmt"

type goodsTrain struct {
    mediator mediator
}

func (g *goodsTrain) requestArrival() {
    if g.mediator.canLand(g) {
        fmt.Println("GoodsTrain: Landing")
    } else {
        fmt.Println("GoodsTrain: Waiting")
    }
}

func (g *goodsTrain) departure() {
    g.mediator.notifyFree()
    fmt.Println("GoodsTrain: Leaving")
}

func (g *goodsTrain) permitArrival() {
    fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}
```

**mediator.go**

```go
package main

type mediator interface {
    canLand(train) bool
    notifyFree()
}
```

**stationmanager.go**

```go
package main

import "sync"

type stationManager struct {
    isPlatformFree bool
    lock           *sync.Mutex
    trainQueue     []train
}

func newStationManger() *stationManager {
    return &stationManager{
        isPlatformFree: true,
        lock:           &sync.Mutex{},
    }
}

func (s *stationManager) canLand(t train) bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.isPlatformFree {
        s.isPlatformFree = false
        return true
    }
    s.trainQueue = append(s.trainQueue, t)
    return false
}

func (s *stationManager) notifyFree() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if !s.isPlatformFree {
        s.isPlatformFree = true
    }
    if len(s.trainQueue) > 0 {
        firstTrainInQueue := s.trainQueue[0]
        s.trainQueue = s.trainQueue[1:]
        firstTrainInQueue.permitArrival()
    }
}
```

**main.go**

```go
package main

func main() {
    stationManager := newStationManger()
    passengerTrain := &passengerTrain{
        mediator: stationManager,
    }
    goodsTrain := &goodsTrain{
        mediator: stationManager,
    }
    passengerTrain.requestArrival()
    goodsTrain.requestArrival()
    passengerTrain.departure()
}
```

**Output:**

```go
PassengerTrain: Landing
GoodsTrain: Waiting
PassengerTrain: Leaving
GoodsTrain: Arrival Permitted. Landing
```

## Full Working Code

```go
package main

import (
    "fmt"
    "sync"
)

type train interface {
    requestArrival()
    departure()
    permitArrival()
}

type passengerTrain struct {
    mediator mediator
}

func (g *passengerTrain) requestArrival() {
    if g.mediator.canLand(g) {
        fmt.Println("PassengerTrain: Landing")
    } else {
        fmt.Println("PassengerTrain: Waiting")
    }
}

func (g *passengerTrain) departure() {
    fmt.Println("PassengerTrain: Leaving")
    g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
    fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}

type goodsTrain struct {
    mediator mediator
}

func (g *goodsTrain) requestArrival() {
    if g.mediator.canLand(g) {
        fmt.Println("GoodsTrain: Landing")
    } else {
        fmt.Println("GoodsTrain: Waiting")
    }
}

func (g *goodsTrain) departure() {
    g.mediator.notifyFree()
    fmt.Println("GoodsTrain: Leaving")
}

func (g *goodsTrain) permitArrival() {
    fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}

type mediator interface {
    canLand(train) bool
    notifyFree()
}

type stationManager struct {
    isPlatformFree bool
    lock           *sync.Mutex
    trainQueue     []train
}

func newStationManger() *stationManager {
    return &stationManager{
        isPlatformFree: true,
        lock:           &sync.Mutex{},
    }
}

func (s *stationManager) canLand(t train) bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.isPlatformFree {
        s.isPlatformFree = false
        return true
    }
    s.trainQueue = append(s.trainQueue, t)
    return false
}

func (s *stationManager) notifyFree() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if !s.isPlatformFree {
        s.isPlatformFree = true
    }
    if len(s.trainQueue) > 0 {
        firstTrainInQueue := s.trainQueue[0]
        s.trainQueue = s.trainQueue[1:]
        firstTrainInQueue.permitArrival()
    }
}

func main() {
    stationManager := newStationManger()
    passengerTrain := &passengerTrain{
        mediator: stationManager,
    }
    goodsTrain := &goodsTrain{
        mediator: stationManager,
    }
    passengerTrain.requestArrival()
    goodsTrain.requestArrival()
    passengerTrain.departure()
}
```

**Output:**

```go
PassengerTrain: Landing
GoodsTrain: Waiting
PassengerTrain: Leaving
GoodsTrain: Arrival Permitted. Landing
```
