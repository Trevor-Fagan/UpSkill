package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
    firstName string
	lastName string
	contact contactInfo
}

func main () {
	jim := person{
		firstName: "Jim",
		lastName: "Parson",
		contact: contactInfo{
			email: "jimparson@gmail.com",
			zipCode: 12345,
		},
	}

	jimPointer := &jim
	jimPointer.print()
	jimPointer.updateName("John")
	jimPointer.print()
}

func (p *person) updateName (firstName string) {
	(*p).firstName = firstName
}

func (p person) print () {
	fmt.Println(p.firstName, p.lastName)
	fmt.Println(p.contact.email)
	fmt.Println(p.contact.zipCode)
}