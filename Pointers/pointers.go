package main

import "fmt"

//a variable that stores the memory address of another variable... symbols '&'  → "WHERE?"  '*'  → "WHAT is there?"

// age = 20 a pointer doesnt actually store 20. it hold the address where 20 is held
// its fundamental for modifying data across functions without making expensive memory copies

func changeAge(age *int) { //this function expects an address pointing to an integer
	*age = 25 // go to the address stored in age and access the value there
}

func main() {
	y := 32
	var z *int = &y
	fmt.Println("y=32, z then copies the memory address at y;", z)
	fmt.Println("y=*z now here we are getting the value stored at address y", *z)

	x := 10
	fmt.Println("the value of x is;", x)
	fmt.Println("the memory address at &x is;", &x) //0x2530377e00b0 is 'x' memory address

	// '&' symbol extracts the memory address of the variable doing p=&x we r storing the address of x in another variable
	p := &x // so p is *int, why, because p contains an address pointing of an int
	fmt.Println("after p copying the memory address of/at x;", p)
	fmt.Println("now lets go and check the value stored at that memory address at x;", *p) // if pointers contains an address using '*' ,is how we get the value stored at the address
	//changing the value at the pointer
	*p = 25
	fmt.Println("doing; *p=25 we are basically going to the memory address at x and changing the value there from 10 to;", x)

	age := 20
	changeAge(&age)
	fmt.Println("function changeAge, age initially was 25 but on doing age:=20 changeAge(&age), now;", age)
}
