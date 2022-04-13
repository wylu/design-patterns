# Chain of Responsibility Design Pattern

[Chain of Responsibility Design Pattern in Go](https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/)

## Definition

Chain of Responsibility Design Pattern is a behavioral design pattern. It lets you create a chain of request handlers. For every incoming request, it is passed through the chain and each of the handler:

1. Processes the request or skips the processing.
2. Decides whether to pass the request to the next handler in the chain or not

Chain of Responsibility Design pattern will be best understood with an example. Let’s take the case of a hospital. A hospital has multiple departments such as:

1. Reception
2. Doctor
3. Medicine Room
4. Cashier

Whenever any patient arrives he first goes to **Reception** then to **Doctor** then to **Medicine Room** and then to **Cashier** and so on. In a way, a patient is sent into a chain of departments which when done, sends the patient to further departments. This is where the Chain of Responsibility pattern comes into the picture.

## When to Use

- The pattern is applicable when there are multiple candidates to process the same request.

- When you don’t want the client to choose the receiver as multiple objects can handle the request. Also, you want to decouple the client from receivers. The Client only needs to know the first element in the chain.

  As in the example of the hospital, a patient first goes to the reception and then reception based upon a patient’s current status sends up to the next handler in the chain.

## UML Diagram

![Chain-of-Responsibility-Design-Pattern-1](images/Chain-of-Responsibility-Design-Pattern-1.webp)

Below is the corresponding mapping UML diagram with the practical example given below

![Chain-of-Responsibility-Design-Pattern-2](images/Chain-of-Responsibility-Design-Pattern-2.webp)

## Mapping

| Role               | File          |
| ------------------ | ------------- |
| handler            | department.go |
| Concrete Handler 1 | account.go    |
| Concrete Handler 2 | doctor.go     |
| Concrete Handler 3 | medical.go    |
| Concrete Handler 4 | cashier.go    |
| Client             | main.go       |

## Practical Example

**department.go**

```go
package main

type department interface {
    execute(*patient)
    setNext(department)
}
```

**reception.go**

```go
package main

import "fmt"

type reception struct {
    next department
}

func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

func (r *reception) setNext(next department) {
    r.next = next
}
```

**doctor.go**

```go
package main

import "fmt"

type doctor struct {
    next department
}

func (d *doctor) execute(p *patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

func (d *doctor) setNext(next department) {
    d.next = next
}
```

**medical.go**

```go
package main

import "fmt"

type medical struct {
    next department
}

func (m *medical) execute(p *patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to patient")
        m.next.execute(p)
        return
    }
    fmt.Println("Medical giving medicine to patient")
    p.medicineDone = true
    m.next.execute(p)
}

func (m *medical) setNext(next department) {
    m.next = next
}
```

**cashier.go**

```go
package main

import "fmt"

type cashier struct {
    next department
}

func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
    c.next = next
}
```

**patient.go**

```go
package main

type patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}
```

**main.go**

```go
package main

func main() {
    cashier := &cashier{}
    //Set next for medical department
    medical := &medical{}
    medical.setNext(cashier)
    //Set next for doctor department
    doctor := &doctor{}
    doctor.setNext(medical)
    //Set next for reception department
    reception := &reception{}
    reception.setNext(doctor)
    patient := &patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}
```

**Output:**

```go
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
```

## Full Working Code

```go
package main

import "fmt"

type department interface {
    execute(*patient)
    setNext(department)
}

type reception struct {
    next department
}

func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

func (r *reception) setNext(next department) {
    r.next = next
}

type doctor struct {
    next department
}

func (d *doctor) execute(p *patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

func (d *doctor) setNext(next department) {
    d.next = next
}

type medical struct {
    next department
}

func (m *medical) execute(p *patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to patient")
        m.next.execute(p)
        return
    }
    fmt.Println("Medical giving medicine to patient")
    p.medicineDone = true
    m.next.execute(p)
}

func (m *medical) setNext(next department) {
    m.next = next
}

type cashier struct {
    next department
}

func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
    c.next = next
}

type patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

func main() {
    cashier := &cashier{}
   
    //Set next for medical department
    medical := &medical{}
    medical.setNext(cashier)
   
    //Set next for doctor department
    doctor := &doctor{}
    doctor.setNext(medical)
   
    //Set next for reception department
    reception := &reception{}
    reception.setNext(doctor)
   
    patient := &patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}
```

**Output:**

```go
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
```
