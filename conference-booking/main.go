package main

import (
	"fmt"
)

func main() {
	const conferenceName = "Golang Conference"
	const totalTickets = 100
	remainingTickets, bookings := getBookings(totalTickets)

	fmt.Printf("%s sold %d tickets, %d tickets remaining\n", conferenceName, len(bookings), remainingTickets)
}

func getBookings(remainingTickets uint) (remaining uint, bookings []string) {
	for remainingTickets > 0 {
		firstName, tickets := getUserInfo()

		bookings = append(bookings, firstName)
		remainingTickets -= tickets
	}

	return remainingTickets, bookings
}

func getUserInfo() (firstName string, tickets uint) {
	fmt.Print("Welcome to Luki's concert, please enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("How many tickets do you want? ")
	fmt.Scanln(&tickets)

	return firstName, tickets
}
