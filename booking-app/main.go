package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	bookings := []string{}
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	// ask for user name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("please enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Println("how many tickets do you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for purchasing %v tickets!\nYou will receive a confirmation receipt to your %v email\n", firstName, lastName, userTickets, userEmail)

	fmt.Printf("there are %v tickets remaining\n", remainingTickets)
	fmt.Printf("these are all of our bookings: %v\n", bookings)


}
