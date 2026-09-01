package main
//array store value of the same type in a single variable, 
// the size of the array is fixed and cannot be changed, the index of an array starts from 0
import "fmt"

func main() {
	var persons = [3] string{"john", "ken", "doe"}
	fmt.Println("Persons array:", persons)
	people := [] string{"joe", "ned", "niit"}
	people = append(people, "james", "jane")
	people[3] = "peter"
	people = append(people, persons[0], persons[1], persons[2])
	// the above can be simplified as people = append(people, persons[:]...) / append(people, persons...) 
	// which will append all the elements of persons array to the people slice
	fmt.Println("People array:", people)
	//len() function returns the number of elements in an array, slice, or map
	fmt.Println(len(people), "number rep the number in people array")
	fmt.Println("the following is the second element in the persons array;", persons[2])
}