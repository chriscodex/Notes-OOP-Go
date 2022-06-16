package main

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

func main() {

}
