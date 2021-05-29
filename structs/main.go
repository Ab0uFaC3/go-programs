package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber int
}

// Defining a struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	robert := person{
		firstName: "Robert",
		lastName:  "Downey",
		contact: contactInfo{
			email:       "robert@gmail.com",
			phoneNumber: 9181039290,
		},
	}

	robert.updateFirstName("Rony")
	robert.print()

}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
