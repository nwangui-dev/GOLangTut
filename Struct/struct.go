package main

// struct (structure) collection of members of different data types into a single unit/variable
type Person struct {
	Name string
	Age int
	Address string
}
//Pass struct as function argument
func printPersonInfo(p Person) {
	println("Name:", p.Name)
	println("Age:", p.Age)
	println("Address:", p.Address)
}
//to access the members of the struct use the (.)dot operator
func main() {
	person1 := Person{Name: "Person1 ..John Doe", Age: 30, Address: "123 Main St"}
	//accessing the members of the struct using the dot operator
	printPersonInfo(person1)
	// println("Age:", person1.Age)
	// println("Address:", person1.Address)
	person2 := Person{Name: "Person2 ..Jay Joe", Age: 23, Address: "45556 Main Street"}
	printPersonInfo(person2)
}
