package main

import "fmt"

type contact struct {
	email   string
	contact int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {

	    vishnu := person{
		firstName: "Vishnu",
		lastName:  "Prakash",
		contact: contact{
			email:   "hi@hello.com",
			contact: 9349838412,
		},
	}
	vishnu.printDetails()
	vishnu.updateFirstName("Jishnu")
	vishnu.printDetails()
}

func (p person) printDetails() {
	fmt.Printf("%+v\n", p)
}

// *person means the type is a pointer pointing to person
func (pointerToPerson *person) updateFirstName(newName string) {
	// *pointerToPerson will be used to reference object stored in the pointer
	(*pointerToPerson).firstName = newName
}
