package main

import "fmt"

type contactInfo struct {
	email 	string
	zipCode int
}

type person struct {
	firstName 	string
	lastName 	string
	contact 	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "party",
		contact: contactInfo{
			email: "jim@email.com",
			zipCode: 12345,
		},
	}

	fmt.Println(jim)
}