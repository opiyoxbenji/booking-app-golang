package main

import "fmt"

func main() {
	// := only applies to vars
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// user input
	fmt.Println("Enter your first & last name and email: ")
	fmt.Scan(&firstName, &lastName, &email)

	fmt.Println("How many tickets you want?: ")
	fmt.Scan(&userTickets)

	// remainingTickets
	remainingTickets = remainingTickets - userTickets

	// userTickets = 4
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName,lastName, userTickets, email)
	fmt.Printf("There are now %v tickets remaining for %v\n", remainingTickets, conferenceName)
}