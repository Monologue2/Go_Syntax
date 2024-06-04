package main

import (
	"fmt"
	"sort"
)

func main() {
	// Basic declare a slice
	var s []int
	fmt.Println("Uninitialized slice:", s)

	//Initialize a slice using a slice literal
	s = []int{1, 2, 3, 4, 5}
	fmt.Println("Initialized slice:", s)

	//declare and Initialize a slice using a slice literal
	var s01 []int = []int{7, 8, 9, 10}
	fmt.Println("Initalized slice 2:", s01)

	//Appending to slices
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println("After appending :", s)

	// Append another slice to a slice
	// ... -> Slice를 개별 함수로 전달 할 수 있습니다.
	//builtin : func append(slice []Type, elems ...Type) []Type
	s = append(s, s01...)
	fmt.Println("After appending another slice:", s)

	//Slicing slice
	s02 := s[1:3]
	fmt.Println("Slice from index 1 to 3 :", s02)
	s03 := s[:3]
	fmt.Println("Slice from start to index 3", s03)
	s04 := s[2:]
	fmt.Println("Slice from index 2 to end", s04)

	//Iterating over slice
	//using traditional for loop
	for i := 0; i < len(s); i++ {
		fmt.Printf("Index %d : %d\n", i, s[i])
	}
	//using range loop
	for index, value := range s {
		fmt.Printf("Range Iterating, Index %d : %d\n", index, value)
	}

	//Copying slice
	dst := make([]int, len(s))
	copy(dst, s) // copy(destination, source), copy function is a built-in function.
	fmt.Println("Source slice :", s)
	fmt.Println("Destiantion slice :", dst)

	// Using `make` to create slice
	make_s := make([]int, 5, 10)
	fmt.Println("Slice with lengh 5 and capacity 10", make_s)
	fmt.Printf("Length : %d, Capasicity : %d\n", len(make_s), cap(make_s))

	// Slices of slices (2D slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//iterate over a 2D slice
	fmt.Println("Iterate over a 2D slice")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	//Remove elements from a slice
	removed_elements := append(s[:2], s[3:]...)
	fmt.Println("After removing element at index 2:", removed_elements)
	//function version
	removed_elements = remove(removed_elements, 2)
	fmt.Println("After removing element at index 2:", removed_elements)

	//Sorting slice
	//import "sort"
	sort.Ints(s)
	fmt.Println("Sorted slice :", s)
	//sort the slice in descending orfer
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("Sorted slice in descending order :", s)

	//Finding elements in slice
	fmt.Println("Slice contains 3 :", contains(s, 3))
	fmt.Println("Slice contains 6 :", contains(s, 6))

}

// Remove elements from a slice
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

// Finding elements in slice
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
