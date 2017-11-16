package main

import (
	"fmt"
)

func main() {

	a := true

	if (a) {
		fmt.Println("Control: ", a)
	}

	b := false
	if (a && b) {
		fmt.Println("Control: ", a)
	}

	myMap := map[string]string {
		"Denys": "Makarov",
		"Vladimir": "Dema",
	}

	if lastName, ok := myMap["Vladimir"]; ok && lastName != "Dema" {
		fmt.Println("LastName: ", lastName)
	} else {
		fmt.Println("LastName: ", "No Last Name")
	}

	// cycles

	for {
		fmt.Println("In cycle")
		break
	}

	myArray := [6]int {6, 5, 4, 3, 2, 1}
	idx := 0

	for idx < 3 {
		fmt.Println("Array [", idx, "] =", myArray[idx])
		idx++
	}

	mySlice := []int {1, 2, 3, 4, 5, 6}

	for idx, value := range mySlice {
		fmt.Println("Slice [", idx, "] =", value)
	}

	for _, value := range mySlice {
		fmt.Println("Slice =", value)
	}

	myNewMap := map[string]string {
		"Denys" : "Makarov",
		"Vladimir" : "Dema",
	}

	for key, value := range myNewMap {
		fmt.Println("Map key =", key, "value =", value)
	}


	/*
		Switch

		fallthrough - pass to next step
	 */
	switch myNewMap["Denys"] {
	case "Petrov", "Ivanov":
		fmt.Println("Petrov or Ivanov ???")
	case "Makarov":
		fmt.Println("Makarov !!!")
		fallthrough
	default:
		fmt.Println("Default !!!")
	}

	first := true
	last := false

	// multiple if else substitution
	switch {
	case first && last:
		fmt.Println("first AND last")
	case first || last:
		fmt.Println("first OR last")
	}

}