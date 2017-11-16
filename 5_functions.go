package main

import "fmt"

func main() {
	fmt.Println("sum = ", sum(1,2))
	fmt.Println("sum = ", sumSimple(1,2))

	sl := [] int {1,2,3}
	fmt.Println("sum = ", sumMultiple(sl...))

	result, error := sumError(append(sl, -1) ...)

	fmt.Println("sum = ", result, "error = ",error)
}

func sum(i int,j int) int {
	return i + j
}


func sumSimple(i,j int) int {
	return i + j
}


func sumMultiple(arr ... int) (res int) {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func sumError(arr ... int) (res int, e error) {
	sum := 0
	for i := range arr {
		if arr[i] < 0 {
			return 0, fmt.Errorf("value is less then 0, %d", arr[i])
		}
		sum += arr[i]
	}
	return sum, nil
}
