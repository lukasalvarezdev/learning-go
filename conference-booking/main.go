package main

import (
	"fmt"
)

func main() {
	const conferenceName = "Golang Conference"
	const totalTickets = 100
	remainingTickets, bookings := getBookings(totalTickets)

	fmt.Printf("%s sold %d tickets, %d tickets remaining\n", conferenceName, len(bookings), remainingTickets)
	fmt.Println("Bookings:", bookings)
}

func getBookings(remainingTickets uint) (remaining uint, bookings []string) {
	for remainingTickets > 0 {
		firstName, tickets := getUserInfo(remainingTickets)

		bookings = append(bookings, firstName)

		if tickets < 0 {
			fmt.Println("Invalid number of tickets, must be greater than 0")
			continue
		}

		if uint(tickets) > remainingTickets {
			fmt.Println("Invalid number of tickets, must be less than or equal to", remainingTickets)
			continue
		}

		remainingTickets -= uint(tickets)
	}

	return remainingTickets, bookings
}

func getUserInfo(remainingTickets uint) (firstName string, tickets int) {
	fmt.Printf("Welcome to Luki's concert, there are %v tickets remaining, please enter your first name: \n", remainingTickets)
	fmt.Scanln(&firstName)

	fmt.Print("How many tickets do you want? \n")
	fmt.Scanln(&tickets)

	return firstName, tickets
}
