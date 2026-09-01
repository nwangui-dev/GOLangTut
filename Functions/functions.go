package main

import "fmt"
//A function is a reusable block of code that performs a specific task. 
// a 'return' is used to tell the function to return a value to the caller / exit the function
func myMessage() {
	fmt.Println("Hello, this is a message from the myMessage function!")
}
//Parameters and arguments: Parameters are variables that are defined in the function signature and act as placeholders for the values .
//  Arguments are the actual values that are provided to the function when it is invoked//called.
func greeting(name string){ //name string is the parameter 
	fmt.Println("Hello,", name, "Welcome to the Go programming language!")
}
// return
func AddNumbers(a, b int) int {
	return a + b
}
// for a function to return a value you need to define the datatype of the return value and use the return keyword inside the function
func myFunction(x , y int) int {
	return x + y
}
//Named return values
func myFunction1(m,n int) (result int) {
	result = m + n
	return // or you can also specify the variable name(result) to this return statement
}
//store the return value in a variable
func myFunction2(o,p int) (result int){
	result = o + p
	return result
}
//multiple return values
func myFunction3(o string, p int) (result int, txt string){
	result = p + p
	txt = o + "world"
	return
}
//storing two returned values in two returned variables
func myFunction4(p int, y string) (result int, txt string){
	result = p + p
	txt = y + "there"
	return
}
//if you want to omit any of the returned values use an underscore
func myFunction5(p int, y string)(result int, txt string){
	result = p+p
	txt = y + "there"
	return
}
//recursion functions
// A recursive function is a function that calls itself
// a function is recursive if it calls itself and reaches a stop condition
func myFactorial(n int) int {
	if n == 0 { //this is the base case, (tells the function stop calling urself)the condition that stops the recursion
		return 1 //if you make a mistake of returning 0 here the result wld be 0
	}
	//return 5 * myFactorial(4), 4 * myFactorial(3) ... downwards until it reaches the base case, then it will return 1 and multiply all the values together to get the final result
	return n * myFactorial(n-1) //the recursive function tells the function, call yourself but with a smaller value of n until you reach the base case (n == 0)
}
func openBox(box int) {
	if box == 0 {
		fmt.Println("All boxes are opened!")
		return //upon the condition being met 'return' will stop the recursion and exit the function
	}
	//return box * openBox(box-1)
	fmt.Println("Opening box number:", box)
	 openBox(box - 1) //Every time the function calls itself, box gets smaller. 
}

// func hello() {  //this here is an infinite recursive function because it calls itself without a base case to stop the recursion, this will cause a stack overflow error
// 	fmt.Println("Hello")
// 	hello()
// }

func main() {

	myMessage() // calling the function

	greeting("John") // calling the function with an argument

	result := AddNumbers(5, 10)
	fmt.Println("The sum is:", result)

	fmt.Println("here we were learning function return values and defining their type int;", myFunction(4,7))

	fmt.Println("here we were learning Named return values", myFunction1(9, 29))

	total := myFunction2(8 , 45)
	fmt.Println("executing return values(int) that were strored in variable total;", total)

	fmt.Println(myFunction3("Hello ... executing multiple return values hello as string and 9+9 as int ", 9))

	a, b := myFunction4(9, "Hello")
	fmt.Println("storing two returned values in two returned variables a int & b string;", a, b)

	_, b = myFunction5(9, "Hey!")
	fmt.Println("omitted 'a' of the returned values using an underscore now only executing b;", b)

	result = myFactorial(5)
	fmt.Println("The factorial of 5 is:", result)

	openBox(5) // calling the recursive function to open 5 boxes
}
