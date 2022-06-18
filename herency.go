package main

import "fmt"

// Structs
type Person struct {
	name string
	age  int
}

type Profession struct {
	id int
}

// Herency
// Struct FullTime inherits from Profession and Person
type FullTimeEmployee struct {
	Person
	Profession
	endDate string
}

// Struct Temporary inherits from Profession and Person
type PartTimeEmployee struct {
	Person
	Profession
	taxRate int
}

// Interfaces
type PrintInfo interface {
	getMessage() string
}

// Method of interface
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

// Methods needed for interfaces
// Method for FullTime
func (FullTimeEmployee) getMessage() string {
	return "This employee works full time"
}

// Method for Temporary
func (PartTimeEmployee) getMessage() string {
	return "This employee works part time"
}

// Function to create new full time employee
func NewFullTimeEmployee(id int, name string, age int, endDate string) *FullTimeEmployee {
	t := FullTimeEmployee{}
	t.id = id
	t.name = name
	t.age = age
	t.endDate = endDate
	fmt.Printf("%v\n", t)
	return &t
}

// Function to create new part time employee
func NewPartTimeEmployee(id int, name string, age int, taxRate int) *PartTimeEmployee {
	p := PartTimeEmployee{}
	p.id = id
	p.name = name
	p.age = age
	p.taxRate = taxRate
	return &p
}

func main() {
	t := NewFullTimeEmployee(1, "Christian", 24, "25/12")
	p := NewPartTimeEmployee(1, "Eduardo", 25, 5)
	getMessage(t)
	getMessage(p)
}
