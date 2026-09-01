package main

import "fmt"

//collection of method signatures that defines the behaviour of an object
// an interface is also a set of methods that a type must have in order to be considered that interface
type Speaker interface{
	Speak()
}
type student struct{
	name string
}
type teacher struct{
	name string
}

// Anything that has a Speak() method can be a Speaker interface.
func (s student) Speak(){
	fmt.Println("Hey guys!")
}
func (t teacher) Speak(){
	fmt.Println("Goodmorning people")
}
//makeTalk accepts anything that satisfies Speaker type interface
func makeTalk(s Speaker) { //this is the function that accpet that interface type
	s.Speak()
}
func main() {
	s := student{name: "mariam"}
	makeTalk(s)
	t := teacher{name: "david"}
	makeTalk(t)
	
}

