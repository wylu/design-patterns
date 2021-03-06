# Abstract Factory Design Pattern

[Abstract Factory Design Pattern in Go](https://golangbyexample.com/abstract-factory-design-pattern-go/)

## Definition

Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects. It is an abstraction over the factory pattern. It is best explained with an example. Let’s say we have two factories

- **nike**
- **adidas**

Imagine you need to buy a sports kit which has a **shoe** and **short**. Preferably most of the time you would want to buy a full sports kit of a similar factory i.e either **nike** or **adidas.** This is where the abstract factory comes into the picture as concrete products that you want is **shoe** and a **short** and these products will be created by the abstract factory of **nike** and **adidas.**
Both these two factories – **nike** and **adidas** implement **iSportsFactory** interface.
We have two product interfaces.

- **iShoe** – this interface is implemented by **nikeShoe** and **adidasShoe** concrete product.
- **iShort** – this interface is implemented by **nikeShort** and **adidasShort** concrete product.

Now let’s look at code

## Code

**isportsfactory.go**

```go
package main

import "fmt"

type iSportsFactory interface {
    makeShoe() iShoe
    makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
    if brand == "adidas" {
        return &adidas{}, nil
    }
    if brand == "nike" {
        return &nike{}, nil
    }
    return nil, fmt.Errorf("Wrong brand type passed")
}
```

**adidas.go**

```go
package main

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
    return &adidasShoe{
        shoe: shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a *adidas) makeShort() iShort {
    return &adidasShort{
        short: short{
            logo: "adidas",
            size: 14,
        },
    }
}
```

**nike.go**

```go
package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
    return &nikeShoe{
        shoe: shoe{
            logo: "nike",
            size: 14,
        },
    }
}

func (n *nike) makeShort() iShort {
    return &nikeShort{
        short: short{
            logo: "nike",
            size: 14,
        },
    }
}
```

**ishoe.go**

```go
package main

type iShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type shoe struct {
    logo string
    size int
}

func (s *shoe) setLogo(logo string) {
    s.logo = logo
}

func (s *shoe) getLogo() string {
    return s.logo
}

func (s *shoe) setSize(size int) {
    s.size = size
}

func (s *shoe) getSize() int {
    return s.size
}
```

**adidasshoe.go**

```go
package main

type adidasShoe struct {
	shoe
}
```

**nikeshoe.go**

```go
package main

type nikeShoe struct {
    shoe
}
```

**ishort.go**

```go
package main

type iShort interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type short struct {
    logo string
    size int
}

func (s *short) setLogo(logo string) {
    s.logo = logo
}

func (s *short) getLogo() string {
    return s.logo
}

func (s *short) setSize(size int) {
    s.size = size
}

func (s *short) getSize() int {
    return s.size
}
```

**adidasshort.go**

```go
package main

type adidasShort struct {
    short
}
```

**nikeshort.go**

```go
package main

type nikeShort struct {
    short
}
```

**main.go**

```go
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
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func printShortDetails(s iShort) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}
```

**Output:**

```go
Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
```
