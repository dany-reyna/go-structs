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

func (p person) String() string {
	return fmt.Sprintf(
		"firstName: %s, lastName: %s, email: %s, zipCode: %d",
		p.firstName, p.lastName, p.email, p.zipCode,
	)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var pepe person

	pepe.firstName = "Pepe"
	pepe.lastName = "Anderson"

	fmt.Println(pepe)
	fmt.Printf("%+v\n", pepe)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}
