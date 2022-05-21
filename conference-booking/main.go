package main

import (
	"fmt"
	mail "net/mail"
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

type Person struct {
	name    string
	tickets int
}

func getBookings(remainingTickets uint) (remaining uint, bookings []Person) {
	for remainingTickets > 0 {
		name, email, tickets := getUserInfo(remainingTickets)

		isValidName, isValidEmail, isValidTickets := validateUserData(name, email, tickets)

		if !isValidName {
			fmt.Printf("%s is not a valid name\n", name)
			continue
		}

		if !isValidEmail {
			fmt.Printf("%s is not a valid email\n", email)
			continue
		}

		if !isValidTickets {
			fmt.Println("Invalid number of tickets, must be greater than 0")
			continue
		}

		if uint(tickets) > remainingTickets {
			fmt.Println("Invalid number of tickets, must be less than or equal to", remainingTickets)
			continue
		}

		var person = Person{name, tickets}
		bookings = append(bookings, person)
		remainingTickets -= uint(tickets)
	}

	return remainingTickets, bookings
}

func getUserInfo(remainingTickets uint) (name string, email string, tickets int) {
	fmt.Printf("Welcome to Luki's concert, there are %v tickets remaining, please enter your name: \n", remainingTickets)
	fmt.Scanln(&name)

	fmt.Println("Please enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Please enter the number of tickets you would like to purchase:")
	fmt.Scanln(&tickets)

	return
}

func validateUserData(name string, email string, tickets int) (bool, bool, bool) {
	return validateName(name), validateTickets(tickets), validateEmail(email)
}

func validateName(name string) bool {
	return len(name) > 2
}

func validateTickets(tickets int) bool {
	return tickets > 0
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
