package main

import "fmt"

func main() {
	// var name = "Mike"
	var age = 27
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Mike"
	// email := "mike@gmail"

	name, email := "Mike", "mike@gmail"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}