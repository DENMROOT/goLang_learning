package main

import (
	"fmt"
)

const i = 1
const appName = "My Application"

const (
	j = 1
	k = "K_CONSTANT"
	counter = iota
	counter_1
	counter_2
)

var (
	globalInt int
	globalBool bool
	globalString string = "Denys Makarov"
)

func main() {
    fmt.Printf("hello, world\n")

	var defaultInt int
	var defaultBool bool
	fmt.Println("default values: ", defaultInt, defaultBool)

	var v1, v2 int
	fmt.Println("multiple values: ", v1, v2)

	fmt.Println("Globals:", globalInt, globalBool, globalString, len(globalString))

	firstName := "Denys"
	lastName := "Makarov"

	fmt.Println("firstName, lastName: ", firstName, lastName)

	fmt.Println("Constants: ", i , appName, j , k)

	fmt.Println("IOTA: ", counter, counter_1, counter_2)
}