package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var personA person

	personA.firstName = "Alex"
	personA.lastName = "Anderson"

	fmt.Printf("%+v", personA)
}
