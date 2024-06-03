package main

import "fmt"

// Passing array to functions
// by value
func printArray00(arr [5]int) {
	for i, v := range arr {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}
}

// Passing array to functions (Using slice)
// by reference
func printArray01(arr []int) {
	for i, v := range arr {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}
}

func main() {
	//Basic array declaration and initialization
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Iintial array :", arr)

	// Initialize an array with values
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array :", arr2)

	//Accessing array elements
	arr3 := [5]int{10, 200, 3000, 40000, 500000}
	fmt.Println("First element :", arr3[0])
	fmt.Println("Fourth element :", arr3[3])

	arr3[2] = 600000
	fmt.Println("Modified array :", arr)

	//Interaing over an array
	//using for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println("Interating array :", arr[i])
	}

	//using range loop
	for i, v := range arr {
		fmt.Println("Interating array with range", i, " ", v)
	}

	// Multi-Dimensional array
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Multi-dimentional array : ", matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Passing array to functions
	// by value
	printArray00(arr)

	// by reference
	// if you want to avoid copying the array, you can pass a slice instead
	printArray01(arr[:]) // pass the slice of the array.

	//Array of pointers
	x, y := 100, 200
	pointer_arr := [2]*int{&x, &y}
	fmt.Println("First Element(pointer) :", *pointer_arr[0])
	fmt.Println("Second Element(pointer) :", *pointer_arr[1])
	// Modifying values through pointers
	*pointer_arr[0] = 300
	fmt.Println("Modified first element(pointer) :", *pointer_arr[0])

	//Advanced usage : Copying and Comparing arrays
	arr_advanced := [3]int{1, 2, 3}
	var arr_copyed = arr_advanced
	fmt.Println("advanced arr : ", arr_advanced)
	fmt.Println("copyed arr : ", arr_copyed)

	arr_comparing := [3]int{3, 4, 5}
	fmt.Println("arr_copyed == arr_comparing", arr_copyed == arr_comparing)
}
