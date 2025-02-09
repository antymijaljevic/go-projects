package main

import (
	"fmt"
)

type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactinfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Hendrix",
		contactinfo: contactinfo{
			email:   "jim.hendrix@gmail.com",
			zipCode: 23000,
		},
	}

	// jimPointer := &jim // give me the memory address of the value
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFistName string) { // give me the value of the memory address
	(*pointerToPerson).firstName = newFistName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func main() { // doesn't need pointer
//     mySlice := []string{"Hi", "There", "How", "Are", "You"}

//     updateSlice(mySlice)

//     fmt.Println(mySlice)
// }

// func updateSlice(s []string) {
//     s[0] = "Bye"
// }
