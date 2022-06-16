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
func testNewTeacher(id int, name string, age int) {
	t := Teacher{}
	t.id = id
	t.name = name
	t.age = age
	fmt.Printf("%v\n", t)
}

func main() {
	testNewTeacher(1, "Christian", 24)
}
