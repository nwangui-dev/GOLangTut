package main

import "fmt"

func main() {
	age := 22
	hasID := true

	if age >= 18 {
		fmt.Println("You are old enough.")

		if hasID {
			fmt.Println("You may enter the Party.")
		} else {
			fmt.Println("You need an ID to enter.")
		}

	}else if age >= 56 {
		fmt.Println("You are a senior citizen. Go find something to do.")
	} else {
		fmt.Println("Sorry, you are not old enough.")
	}
}