package main

//maps built-in data structure that stores key-value pairs, where each key is unique and maps to a specific value
// default value of a map is nil, length of a map is the number of elements and can be obtained by the ()len function
import "fmt"

func main() {
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	//b := map[KeyType]ValueType{key1:value1, key2:value2,...}

	//var a = map[string]string{"brand":"Ford", "model":"Mustang", "year":"1980"}
	//b := map[string]int{"pencils": 1, "biros": 4, "books": 9}

	//creating a map using the make function
	//var a = make(map[keyType]valueType)
	myMap := make(map[string]int) //map is empty then ..
	//adding key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 3
	myMap["orange"] = 7
	myMap["grape"] = 9
	myMap["apple"] = 2     //doing this changes the value of the key "apple" to 2, as keys are unique in a map
	delete(myMap, "grape") //deleting the key-value pair with the key "grape" from the map

	//accessing values in the map using keys
	fmt.Println("Number of apples:", myMap["apple"]) //go find the value associated with the key "apple"
	fmt.Println("Number of bananas:", myMap["banana"])
	fmt.Println("Number of oranges:", myMap["orange"])
	fmt.Println("Number of mangos:", myMap["mango"]) //this will return 0 as the key "mango" does not exist in the map

	//checking if a key exists in the map
	value, exists := myMap["grape"]
	if exists {
		fmt.Println("Number of grapes:", value)
	} else {
		fmt.Println("Grapes not found in the map.")
	}

	//could also check using ;  val, ok :=map_name[key]
	fruitx, ok := myMap["apple"]
	fmt.Println(fruitx, ok)

	//deleting a key-value pair from the map
	delete(myMap, "banana")
	fmt.Println("After deleting banana, number of bananas:", myMap["banana"])

	//looping through a map using range
	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}
	//getting the length of a map
	fmt.Println("Length of the map:", len(myMap))
	//checking if a map is empty
	if len(myMap) != 0 {
		fmt.Println("The map is not empty.")
	} else {
		fmt.Println("The map is empty.")
	}

}
