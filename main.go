package main

import "fmt"

// Struct (Equivalent to classes)
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

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "name"
	fmt.Println(e)
}
