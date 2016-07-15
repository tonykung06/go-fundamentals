package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 1, 4)
	fmt.Printf("Length is %d.Capacity is: %d", len(slice1), cap(slice1))
	for i := 1; i < 17; i++ {
		slice1 = append(slice1, i)
		fmt.Printf("\nCapacity is: %d", cap(slice1))
	}

	slice2 := []string{"OMG", "OMG2", "OMG3"}
	fmt.Printf("\nLength is %d.\nCapacity is: %d", len(slice2), cap(slice2))

	slice3 := []string{"hi1", "hi2"}
	slice2 = append(slice2, slice3...)
	fmt.Printf("\nLength is %d.\nCapacity is: %d", len(slice2), cap(slice2))

	sliceOfSlice2 := slice2[1:2]
	fmt.Printf("\n%sLength is %d.\nCapacity is: %d", sliceOfSlice2, len(sliceOfSlice2), cap(sliceOfSlice2))
}
