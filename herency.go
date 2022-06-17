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
// Struct FullTimeEmployee inherits from Employee and Person
type Teacher struct {
	Person
	Profession
}

// Test Functions
func NewTeacher(id int, name string, age int) {
	t := Teacher{}
	t.id = id
	t.name = name
	t.age = age
	fmt.Printf("%v\n", t)
}

func main() {
	NewTeacher(1, "Christian", 24)
}
