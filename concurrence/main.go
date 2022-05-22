package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		balance         int
		expenseName     string
		expenseQuantity int
		sendToEmail     bool
	)

	fmt.Println("Enter your initial balance: ")
	fmt.Scanln(&balance)

	fmt.Println("Enter your expense name: ")
	fmt.Scanln(&expenseName)

	fmt.Println("Enter your expense amount: ")
	fmt.Scanln(&expenseQuantity)

	// We will update the balance but I don't need it to show the response to the user
	// So we will use a goroutine to update the balance and prevent the main thread from blocking
	// the keyword "go" does all the magic for us
	go expensiveUpdateBalance(&balance, expenseQuantity)

	fmt.Printf("Congrats!, your expense %s has been deducted from your account. You spent %v", expenseName, expenseQuantity)

	fmt.Println("Do you want this information in your email? (y/n)")
	fmt.Scanln(&sendToEmail)
}

func expensiveUpdateBalance(balance *int, expense int) {
	fmt.Println("--- Calculating your new balance ---")
	time.Sleep(time.Second * 2)

	*balance = *balance - expense

	fmt.Println("--- Your new balance is: ", *balance, " ---")
}
