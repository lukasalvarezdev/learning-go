package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// arrays()
	// maps()
	// loops()
	// sqrtExample()
	// makePerson()
	// memoryAddress()
	// formatTextExample()
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

type Person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson() Person {
	p := Person{
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

type width = float64
type height = float64

// METHODS
type Circle struct {
	x, y, radius float64
}

type Rect struct {
	w width
	h height
}

// area method for circle, passed by value
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// modifier method for circle, passed by reference
func (c *Circle) scale(scale float64) {
	// no need to use *c, it's already a pointer
	c.radius *= scale
}

// area method for rect
func (r Rect) area() float64 {
	return r.w * r.h
}

func shapes() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rect := Rect{w: 10, h: 5}

	fmt.Println("Circle area: ", circle.area())
	fmt.Println("Rect area: ", rect.area())
	circle.scale(2)
	fmt.Println("Circle area after scaling: ", circle.area())
}
