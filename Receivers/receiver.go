package main
//A function can receive something as a parameter. A method has a receiver that connects it to a particular type.
import "fmt"

type Person struct {
	Name string
	Age  int
}
type Dog struct{
	name string
}
 // Think:

// func (dog Dog) bark()

// as:

// Dog
//  │
//  ├── name
//  │
//  └── bark()

func (dog Dog) bark() { //dog Dog is the receiver 
	fmt.Println(dog.name, "says Woof!")
}

// So methods allow you to put behavior alongside the type it belongs to.
func (p Person) printPersonInfo() {
	fmt.Println(p.Name,"of age", p.Age)
}
// The person and p are not the same variable. They just contain the same Person value when the function is called.

func main() {
	person := Person{Name: "Alice", Age: 30}
	person.printPersonInfo()

	dog := Dog{name: "Lexy"}
	dog.bark()
}
