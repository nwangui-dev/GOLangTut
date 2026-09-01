package main

import "fmt"
func main() {
	// arithmetic operators
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	mod := a % b //modulus operator returns the remainder of the division

	// comparison operators
	isEqual := a == b
	isNotEqual := a != b
	isGreater := a > b
	isLess := a < b
	isGreaterOrEqual := a >= b
	isLessOrEqual := a <= b

	// logical operators
	x := true
	y := false
	andResult := x && y // logical AND ...Returns true if both statements are true
	orResult := x || y  // logical OR ...Returns true if one of the statements is true
	notResult := !x     // logical NOT ...Returns true if the statement is false, and false if the statement is true

	// assignment operators
	c := 10
	c += 5 // equivalent to c = c + 5
	c -= 3 // equivalent to c = c - 3
	c *= 2 // equivalent to c = c * 2
	c /= 4 // equivalent to c = c / 4
	c %= 3 // equivalent to c = c % 3

	// bitwise operators
	m := 6  // binary: 110
	n := 3  // binary: 011
	andBitwise := m & n // binary: 010 (decimal: 2) ...Returns 1 if both bits are 1, otherwise returns 0
	orBitwise := m | n  // binary: 111 (decimal: 7) ...Returns 1 if at least one bit is 1
	xorBitwise := m ^ n // binary: 101 (decimal: 5) ...Returns 1 if bits are different
	leftShift := m << 1 // binary: 1100 (decimal: 12) ...Shifts bits to the left
	rightShift := m >> 1 // binary: 011 (decimal: 3) ...Shifts bits to the right

	// print results
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
	fmt.Println("Modulus:", mod)

	fmt.Println("\nComparison Operators:")
	fmt.Println("Is Equal:", isEqual)
	fmt.Println("Is Not Equal:", isNotEqual)
	fmt.Println("Is Greater:", isGreater)
	fmt.Println("Is Less:", isLess)
	fmt.Println("Is Greater or Equal:", isGreaterOrEqual)
	fmt.Println("Is Less or Equal:", isLessOrEqual)

	fmt.Println("\nLogical Operators:")
	fmt.Println("AND Result:", andResult)
	fmt.Println("OR Result:", orResult)
	fmt.Println("NOT Result:", notResult)

	fmt.Println("\nAssignment Operators:")
	fmt.Println("Final Value of c:", c)

	fmt.Println("\nBitwise Operators:")
	fmt.Println("AND Bitwise:", andBitwise)
	fmt.Println("OR Bitwise:", orBitwise)
	fmt.Println("XOR Bitwise:", xorBitwise)
	fmt.Println("Left Shift:", leftShift)
	fmt.Println("Right Shift:", rightShift)

	//lets look at even/odd numbers using modulus operator
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is an even number")
	} else {
		fmt.Println(number, "is an odd number")
	}

}