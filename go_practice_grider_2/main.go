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

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "party@partyhouse.com",
			zipCode: 12233,
		},
	}

	//& operator retrieves the memory address of the given variable. & turns a value into an address

	jim.updateName("Joshua")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//* type descriptor tells us we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	//* operator retrieves the memory value of the given memory address. * turns an address into a value
	(*pointerToPerson).firstName = newFirstName
}

//alex := person{"Alex", "Anderson"}
//var alex = person

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

// jim := person{
// 	firstName: "Jim",
// 	lastName:  "Party",
// 	contact: contactInfo{
// 		email:   "party@partyhouse.com",
// 		zipCode: 12233,
// 	},
// }
