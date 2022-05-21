package main

import (
	"fmt"
)

func main() {
	const conferenceName = "Golang Conference"
	const totalTickets = 100
	remainingTickets, bookings := getBookings(totalTickets)

	fmt.Printf("%s sold %d tickets, %d tickets remaining\n", conferenceName, len(bookings), remainingTickets)

	for _, booking := range bookings {
		fmt.Printf("%s booked %d tickets\n", booking.name, booking.tickets)
	}
}

type person struct {
	name    string
	tickets int
}

func getBookings(remainingTickets uint) (remaining uint, bookings []person) {
	for remainingTickets > 0 {
		name, tickets := getUserInfo(remainingTickets)

		if tickets < 0 {
			fmt.Println("Invalid number of tickets, must be greater than 0")
			continue
		}

		if uint(tickets) > remainingTickets {
			fmt.Println("Invalid number of tickets, must be less than or equal to", remainingTickets)
			continue
		}

		bookings = append(bookings, person{name, tickets})
		remainingTickets -= uint(tickets)
	}

	return remainingTickets, bookings
}

func getUserInfo(remainingTickets uint) (name string, tickets int) {
	fmt.Printf("Welcome to Luki's concert, there are %v tickets remaining, please enter your name: \n", remainingTickets)
	fmt.Scanln(&name)

	fmt.Print("How many tickets do you want? \n")
	fmt.Scanln(&tickets)

	return name, tickets
}
