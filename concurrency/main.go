package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var withSync string

	fmt.Println("Do you want to use the async feature? (y/n)")
	fmt.Scanln(&withSync)

	if withSync != "y" && withSync != "n" {
		fmt.Println("Come on, is it so hard to choose? (y/n)")
		os.Exit(0)
	}

	if withSync == "y" {
		syncRoutine()
	} else {
		asyncRoutine()
	}

}

func getUserData() (
	balance int,
	expenseName string,
	expenseQuantity int,
) {
	fmt.Println("Enter your initial balance: ")
	fmt.Scanln(&balance)

	fmt.Println("Enter your expense name: ")
	fmt.Scanln(&expenseName)

	fmt.Println("Enter your expense amount: ")
	fmt.Scanln(&expenseQuantity)

	return
}

// This will have a bug, the program will exit before the balance is updated
func asyncRoutine() {
	balance, expenseName, expenseQuantity := getUserData()

	// We will update the balance but I don't need it to show the response to the user
	// So we will use a goroutine to update the balance and prevent the main thread from blocking
	// the keyword "go" creates a new thread and does all the magic for us
	go expensiveUpdateBalance(&balance, expenseQuantity, false)

	fmt.Printf("Congrats!, your expense %s has been deducted from your account. You spent %v \n", expenseName, expenseQuantity)
}

var wg = sync.WaitGroup{}

func syncRoutine() {
	balance, expenseName, expenseQuantity := getUserData()

	// wg is basically a number of threads, Add() adds threads to the wait group
	wg.Add(1)
	go expensiveUpdateBalance(&balance, expenseQuantity, true)

	fmt.Printf("Congrats!, your expense %s has been deducted from your account. You spent %v", expenseName, expenseQuantity)

	// Will wait for the thread counter to reach 0 and exits the program
	wg.Wait()
}

func expensiveUpdateBalance(balance *int, expense int, withWg bool) {
	fmt.Println("--- Calculating your new balance ---")
	time.Sleep(time.Second * 2)

	*balance = *balance - expense

	fmt.Println("--- Your new balance is: ", *balance, " ---")

	if withWg {
		// This will decrement the thread counter
		wg.Done()
	}
}
