package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	var alex person
	alex.firstName = "alex"
	alex.lastName = "anderson"
	fmt.Println(alex)
}