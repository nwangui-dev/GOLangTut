package main
import "fmt"

func main() {
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}

	// while loop (using for loop)
	j := 1
	for j <= 5 {
		fmt.Println("While Loop Iteration:", j)
		j++
	}

	// infinite loop (commented out to prevent execution)
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	// break statement
	for k := 1; k <= 10; k++ {
		if k == 4 {
			break // exit the loop when k is 4
		}
		fmt.Println("Break Loop Iteration:", k)
	}

	// continue statement
	for l := 1; l <= 10; l++ {
		if l%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("Continue Loop Iteration (Odd Numbers):", l)
	} 
	// nested loops
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 3; n++ {
			fmt.Printf("Nested Loop Iteration: (%d, %d)\n", m, n)
		}
	}
	// range loop ->range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value
// 	 for index, value := range array|slice|map {
//    // code to be executed for each iteration
// }
	mySlice := []string{"apple", "banana", "cherry"}
	for index, value := range mySlice { //Tip: To only show the value or the index, you can omit the other output using an underscore (_).
		fmt.Printf("Range Loop Iteration: Index %d, Value %s\n", index, value)
	}
	
}