package main

import "fmt"

// constants can be declared both outside and inside of functions and cannot be changed
const pi = 3.14
//typed constants are declared with a defined type
const myName string = "Kenneth Joe"
//untyped constants are declared without a type
const isActive = true


func main() {
	//var can be used inside and outside of functions
	var StudentName string = "John Doe"
	var StudentAge = 20
	// := can only be used inside functions
	year := 2026
	var a bool
	// multiple variable declaration
	//int- signed (+ve and -ve) unsigned (non-negative) float64 is the default of floating numbers and take largest number of bits than float32, complex64 and complex128 are used to store complex numbers,
	var b, c, d int = 1, 2, 3
	var e, f, g = 4, 5, "Hi Go, Multiple varible declaration but with different data types"
	fmt.Println(" testing constants, ...The value of pi is", pi)
	fmt.Println("Testing Typed constant Value of myName is", myName)
	fmt.Println("Testing Untyped constant Value of isActive is", isActive)
	fmt.Println(a)
	fmt.Println(b, c, d)
	fmt.Println(e, f, g)
	fmt.Println("Hello there "+StudentName+"! of age", StudentAge, "class of", year)
}
