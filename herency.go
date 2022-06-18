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
type FullTime struct {
	Person
	Profession
	endDate string
}

// Struct Temporary inherits from Profession and Person
type Temporary struct {
	Person
	Profession
	taxRate int
}

// Function to create new teacher without pointers
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
