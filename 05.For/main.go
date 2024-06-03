package main

import "fmt"

func main() {
	//Basic for
	for i := 0; i < 5; i++ {
		fmt.Println("Basic for loop : ", i)
	}

	//Infinite for
	i := 0
	for {
		fmt.Println("Infinite for : ", i)
		i++
		if i > 5 {
			break
		}
	}

	//For as a while
	for i < 5 {
		fmt.Println("For as a while", i)
		i++
	}

	//Iterating over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Iterating over a slice", numbers[i])
	}

	//Using `range` to iterate over a slice or array
	for index, value := range numbers {
		fmt.Println("Using `range` to iterate over a slice or array")
		fmt.Printf("Index : %d, Value : %d\n", index, value)
	}

	//Iterating over a map
	capitals := map[string]string{
		"South Korea": "Seoul",
		"France":      "Paris",
		"USA":         "Washingthon D.C",
	}
	for country, capital := range capitals {
		fmt.Printf("Map Iterater : The capital of %s is %s\n", country, capital)
	}

	//Iterating over a string
	str := "My name is Shogle/Seungun"
	for index, char := range str { //rune
		fmt.Printf("Index %d, Character : %c\n", index, char)
	}

	// Nested for loop
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Nested for loop")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}

	//Using `continue` and `break`
	fmt.Println("Using continue and break")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	//Using labels with for loop
	fmt.Println("Labels with for loop")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Printf("i : %d, j : %d\n", i, j)
		}
	}
}
