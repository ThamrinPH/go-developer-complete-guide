package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	personA := person{
		firstName: "Jim",
		lastName:  "Part",
		contactInfo: contactInfo{
			email:   "jiim@gmail.com",
			zipCode: 602148,
		},
	}

	personA.updateName("Oscar")
	personA.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
