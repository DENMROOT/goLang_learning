package main

import (
	"fmt"
)

func main() {

	var myMap map[string]string = map[string]string{}
	myMap["Denys"] = "Makarov"
	myMap["Vladimir"] = "Dema"

	myMap_2 := map[string]string{}
	myMap_2["Denys"] = "Makarov"
	var myMap_3 = make(map[string]string)
	myMap_3["Denys"] = "Makarov"

    fmt.Println("Map: ", myMap)
	fmt.Println("Map: ", myMap_2)
	fmt.Println("Map: ", myMap_3)

	lastName, ok := myMap_3["Denys"]
	lastName_1, ok_1 := myMap_3["Petya"]
	_, exist := myMap_3["Petya"]


	fmt.Println("Map Value \"Denys\": ", lastName, ok)
	fmt.Println("Map Value \"Petya\": ", lastName_1, ok_1)
	fmt.Println("Map Value \"Petya\": ", exist)

	delete(myMap, "Denys")
	fmt.Println("Map: ", myMap)
}