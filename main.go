package main

import "fmt"

// Structs (Equivalent to class)
type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "name"
	fmt.Println(e)
}
