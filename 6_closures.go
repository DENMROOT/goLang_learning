package main

import "fmt"

func main() {
	counter := intCounter()
	fmt.Println("", counter())
	fmt.Println("", counter())
	fmt.Println("", counter())

	counter2 := intCounter()
	fmt.Println("", counter2())
	fmt.Println("", counter2())
	fmt.Println("", counter2())
}

func intCounter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
