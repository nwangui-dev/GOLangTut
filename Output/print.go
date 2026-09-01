package main
import "fmt"

func main() {
	var i, j string = "Hello", "World"
	//Print(), Println(), Printf() are the three functions used to print output to the console
	fmt.Print(i, j)
	fmt.Println(i,j)
	//Printf() is used to format the output, based on the given formatting verb
	fmt.Printf(" %s %s ", i, j)
	//%v prints the value in a default format, %#v prints the value in Go syntax format, %T prints the type of the variable, %s prints string values and %t prints the boolean value
	var isActive bool = true
	fmt.Printf(" %v %T %t ", isActive, isActive, isActive)
}