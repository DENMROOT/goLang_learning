package main

import (
	"fmt"
)

func main() {

	var myArray [10] int

    fmt.Println("Array: ", myArray)

    myArray[0] = 10

	fmt.Println("Array: " , myArray, ",size ",len(myArray))

	arr := [] int{1,2,3}
	fmt.Println("Dafault Array: " , arr, ",size ",len(arr))

	var matrix [3][3] int
	matrix[0][1] = 1
	fmt.Println("Matirx: " , matrix, ",size ",len(matrix))
}