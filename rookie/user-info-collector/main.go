package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	GetUserInfo()
	// arrays()
	// maps()
	// loops()
	// sqrtExample()
	// makePerson()
	// memoryAddress()
	// formatTextExample()
}

func preventNegaiveNumbers() {
	var remainingTickets uint = 50
	// will give an error, because unit cannot be negative
	// remainingTickets = -1

	fmt.Println("Remaining tickets: ", remainingTickets)
}

func subtract(a int, b int) int {
	return a - b
}

func formatTextExample() {
	const userName = "Luki"
	const userAge = 25

	fmt.Printf("User: %v is %v years old \n", userName, userAge)
}

func sqrtExample() {
	result, error := getSqrt(4)

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Sqrt: ", result)
	}
}

func getSqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("cannot get sqrt of negative number")
	}

	return math.Sqrt(num), nil
}

func arrays() {
	// array with fixed size
	arr := [5]int{1, 2, 3, 4, 5}

	// array with dynamic size
	arr2 := []int{1, 2, 3, 4, 5}

	// add element to array, returns new array
	newArr := append(arr2, 6)

	fmt.Println("Fixed: ", arr)
	fmt.Println("Non-fixed: ", arr2)
	fmt.Println("New element: ", newArr)
}

func maps() {
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	fmt.Println("Map with 2 keys: ", myMap)

	delete(myMap, "one")

	fmt.Println("Map with 1 key: ", myMap)
}

func loops() {
	arr := [5]int{1, 2, 3, 4, 5}
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	// for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println("for loop: ", arr[i])
	}

	// while loop
	i := 0
	for i < len(arr) {
		fmt.Println("while loop: ", arr[i])
		i++
	}

	// for range loop
	for index, value := range arr {
		fmt.Println("for range loop: ", index, value)
	}

	// for range loop with map
	for key, value := range myMap {
		fmt.Println("for range loop with map: ", key, value)
	}
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson() person {
	p := person{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
	}

	return p
}

func memoryAddress() {
	num := 10

	// useless, won't change the value of num
	incrementCopy(num)

	// will change the value of num, note the & to pass the pointer
	incrementReal(&num)

	fmt.Println("Memory address of num: ", &num)
	fmt.Println("Increased num: ", num)
}

func incrementCopy(num int) {
	// num is a copy of the original value, so the original value is not changed
	num++
}

// note the * to receive the pointer
func incrementReal(num *int) {
	// num is a pointer to the original value, so the original value is changed
	// the * is used to dereference the pointer
	*num++
}