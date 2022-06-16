package main

import "fmt"

// Struct (Equivalent to classes)
// Struct Employee
type Employee struct {
	id       int
	name     string
	vacation bool
}

// Constructor
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

// Receiver Function (Equivalent to methods)
// Setters
func (e *Employee) SetId(id int) {
	e.id = id
}
func (e *Employee) SetName(name string) {
	e.name = name
}

// Getters
func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) GetName() string {
	return e.name
}

// Test Functions
// Constructor Test
func ConstructorTest() {
	// Way 1
	e := NewEmployee(1, "Employee 1", true)
	fmt.Printf("%v\n", *e)
	// Way 2
	e2 := new(Employee)
	e2.id = 2
	e2.name = "Employee 2"
	e2.vacation = true
	fmt.Printf("%v\n", *e2)
	// Way 3
	e3 := Employee{
		id:       3,
		name:     "Employee 3",
		vacation: true,
	}
	fmt.Printf("%v\n", e3)
}

func main() {
	ConstructorTest()
}
