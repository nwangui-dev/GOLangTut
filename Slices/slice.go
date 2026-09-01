package main

import "fmt"

// more flexible than arrays,
func main() {
	// slice_name := []datatype{values}
	mySlice := []string{"lynn", "dorcas", "macy"}
	fmt.Println("mySlice ;", mySlice)
	fmt.Println(len(mySlice), "is the length of mySlice")
	//  slice_name = append(slice_name, element1, element2, ...)
	mySlice = append(mySlice, "james", "jane")
	fmt.Printf("mySlice after appending/adding new elements ;%v\n", mySlice)
	fmt.Println(cap(mySlice), "is the capacity of mySlice after appending new elements")

	// slice_name := make([]type, length, capacity)
	mySlice1 := make([]string, 5, 10)
	fmt.Println("mySlice1 ;", mySlice1)
	fmt.Println(len(mySlice1), "is the length of mySlice1")
	fmt.Println(cap(mySlice1), "is the capacity of mySlice1")
	mySlice1 = append(mySlice1, "peter", "brooks", "brian")
	fmt.Printf("mySlice1 after appending/adding new elements ;%v\n", mySlice1)
	fmt.Println(len(mySlice1), "is the length of mySlice1 after appending new elements")
	fmt.Println(cap(mySlice1), "is the capacity of mySlice1 after appending new elements")

	myslice3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("myslice3 ;", myslice3)
	fmt.Println(len(myslice3), "is the length of myslice3")
	fmt.Println(cap(myslice3), "is the capacity of myslice3")
	myslice3 = append(myslice3, 7, 8, 9)
	fmt.Printf("myslice3 after appending/adding new elements ;%v\n", myslice3)
	fmt.Println(len(myslice3), "is the length of myslice3 after appending new elements")   //new length is 9
	fmt.Println(cap(myslice3), "is the capacity of myslice3 after appending new elements") // notice the capacity may become larger than the new length because Go often allocates more memory than needed to avoid frequent memory reallocation when the slice grows, this is called amortized allocation
	//removing elements from a slice can be done by using slicing and append() function
	myslice3 = append(myslice3[:2], myslice3[4:]...) //[:2]take everything from index 2 but not including index 2, [4:] then take everything from index 4 to the end and append it to the first part, this effectively removes the 3rd and 4th elements from the slice
	fmt.Printf("myslice3 after removing the 3rd and 4th elements ;%v\n", myslice3)
	fmt.Println(len(myslice3), "is the length of myslice3 after removing the 3rd and 4th elements")
	fmt.Println(cap(myslice3), "is the capacity of myslice3 after removing the 3rd and 4th elements")

	mySlice4 := []int{21, 34, 55, 89, 14, 23, 45, 67, 89}
	mySlice4 = mySlice4[1:5] // means give me the elements from index 1 up to (but not including) index 5.
	fmt.Println("mySlice4 ;", mySlice4)
	fmt.Println(len(mySlice4), "is the length of mySlice4")
	fmt.Println(cap(mySlice4), "is the capacity of mySlice4")
	mySlice4 = mySlice4[1:3] //change length by reslicing,
	fmt.Println("mySlice4 after reslicing ;", mySlice4)
	fmt.Println(len(mySlice4), "is the length of mySlice4 after reslicing")
	fmt.Println(cap(mySlice4), "is the capacity of mySlice4 after reslicing")
	mySlice4 = append(mySlice4, 100, 200, 300) //change length and capacity by appending new elements
	fmt.Println("mySlice4 after appending new elements ;", mySlice4)
	fmt.Println(len(mySlice4), "is the length of mySlice4 after appending new elements")
	fmt.Println(cap(mySlice4), "is the capacity of mySlice4 after appending new elements")

	//slices can be created from arrays, the slice will point to the same underlying array, so changes made to the slice will affect the array and vice versa
	//myslice := myarray[start:end] // A slice made from the array
	//a slice to an array is like a window into the array, it does not copy the elements of the array, it just points to the same underlying array elements, so changes made to the slice will affect the array and vice versa
	myArray := [5]string{"john", "ken", "doe", "joe", "ned"}
	mySlice2 := myArray[2:4] //starts at index 2 and stops before index 4
	fmt.Println("mySlice2 ;", mySlice2)
	mySlice2[0] = "peter" // now here slice changes the underlying array, so the first element of mySlice2 is changed to "peter", and the underlying array the index at 2 of myArray is also changed
	fmt.Println(cap(mySlice2), "is the capacity of mySlice2")
	mySlice2 = append(mySlice2, "james") // slice still has unused capacity cap =3 so so append() uses the existing underlying array instead of creating a new one "james" replaces "ned" in the array.
	//mySlice2 = append(mySlice2, "james", "mary", "cynthia", "collins", "linda", "grace", "paul") //if you do this slice will no longer be pointing to the same underlying array, because theres no enough capacity for all these new elements, so Go creates a new underlying array and copies the elements of the old array to the new one, so now changes made to the slice will not affect the original array
	fmt.Printf("mySlice2 after changing the first element ;%v\n", mySlice2)
	fmt.Println("myArray after changing the first element of mySlice2 ;", myArray)

	// The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	mySlice5 := []int{20, 32, 45, 67, 89, 37, 28, 27, 77, 90, 11, 21, 34, 56, 78, 92, 100}
	fmt.Printf("the length of mySlice5 is;%v\n", len(mySlice5))
	fmt.Printf("the capacity of mySlice5 is;%v\n", cap(mySlice5))
	neededNumbers := mySlice5[:len(mySlice5)-10] // take the first half of mySlice5
	fmt.Printf("the length of neededNumbers is;%v\n", len(neededNumbers))
	fmt.Printf("the capacity of neededNumbers is;%v\n", cap(neededNumbers))
	// create a new slice to hold the copied elements
	copiedNumbers := make([]int, len(neededNumbers))
	fmt.Printf("the length of copiedNumbers is;%v\n", len(copiedNumbers))
	fmt.Printf("the capacity of copiedNumbers is;%v\n", cap(copiedNumbers))
	// copy the elements from neededNumbers to copiedNumbers
	copy(copiedNumbers, neededNumbers)
	fmt.Printf("copiedNumbers after copying from neededNumbers;%v\n", copiedNumbers)
}
