package main

import (
	"fmt"
)

func main() {

	var mySlice [] int
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)

    fmt.Println("Slice: ", mySlice, ", size", len(mySlice), ",cap", cap(mySlice))

    mySlice_2 := [] int {5, 4 , 3, 2, 1}
	fmt.Println("Slice: ", mySlice_2, ", size", len(mySlice_2), ",cap", cap(mySlice_2))

	mySlice = append(mySlice, mySlice_2...)

	fmt.Println("Slice: ", mySlice, ", size", len(mySlice), ",cap", cap(mySlice))

	var sliceOfSlices [][] int
	sliceOfSlices = append(sliceOfSlices, mySlice, mySlice_2)

	fmt.Println("Slice Of Slices: ", sliceOfSlices, ", size", len(sliceOfSlices), ",cap", cap(sliceOfSlices))

	// slice with length and capacity
	var sliceWithSize = make([] int, 7, 12)
	sliceWithSize = append(sliceWithSize, []int {1, 2, 3}...)
	fmt.Println("Slice with Size: ", sliceWithSize, ", size", len(sliceWithSize), ",cap", cap(sliceWithSize))

	newSliceWithSize := sliceWithSize
	fmt.Println("New slice with Size: ", newSliceWithSize, ", size", len(newSliceWithSize), ",cap", cap(newSliceWithSize))

	// slice copy

	copiedSlice := make([] int, len(mySlice), cap(mySlice))
	copy(copiedSlice, mySlice)
	fmt.Println("Copied slice : ", copiedSlice, ", size", len(copiedSlice), ",cap", cap(copiedSlice))
	fmt.Println("Part of slice 1:3 : ", copiedSlice[1:3], ", part ... 4", copiedSlice[:4] , ", part 10 ...", copiedSlice[6:])

	myArray := [] int {1,2,3}
	newSlice := myArray[:]
	newSlice[2] = 10
	fmt.Println("Slice from array ", newSlice)

}