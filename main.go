package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	personA := person{
		firstName: "Jim",
		lastName:  "Part",
		contact: contactInfo{
			email:   "jiim@gmail.com",
			zipCode: 602148,
		},
	}

	fmt.Printf("%+v", personA)
}
