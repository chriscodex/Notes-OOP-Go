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
type TemporaryEmployee struct {
	Person
	Profession
	taxRate int
}

// Interfaces
type PrintInfo interface {
	getMessage() string
}

// Methods needed for interfaces
// Method for FullTime
func (FullTimeEmployee) getMessage() string {
	return "This person works full time employee"
}

// Method for Temporary
func (TemporaryEmployee) getMessage() string {
	return "This person works temporary employee"
}

// Function to create new teacher without pointers
func NewFullTimeEmployee(id int, name string, age int, endDate string) *FullTimeEmployee {
	t := FullTimeEmployee{}
	t.id = id
	t.name = name
	t.age = age
	t.endDate = endDate
	fmt.Printf("%v\n", t)
	return &t
}

func main() {
	t := NewFullTimeEmployee(1, "Christian", 24, "25/12")
	fmt.Println(*t)
}
