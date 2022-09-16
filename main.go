package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	personA := person{firstName: "Fichri", lastName: "Ardianto", age: 25}
	fmt.Println(personA, personA.firstName)
}
