package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 10

	// if
	if x > 5 {
		fmt.Println("X is greater than 5!")
	}

	// if-else
	if x > 5 {
		fmt.Println("if-else : ", "X is greater than 5!")
	} else {
		fmt.Println("if-else : ", "X is less than 5!")
	}

	// else if
	if x > 5 {
		fmt.Println("else if : ", "X is greater than 5!")
	} else if x > 8 {
		fmt.Println("else if : ", " X is greater than 15!")
	} else {
		fmt.Println("else if : ", " X is less than 5!")
	}

	// Nested if
	if x > 5 {
		fmt.Println("Nested if :", "X is greater than 5!")
		if x > 7 {
			fmt.Println("Nested if : ", "And X is greater than 7!")
		}
	}

	// error check
	str := "1998"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Println("Error check, Converted number : ", num)
	} else {
		fmt.Println("Error converting string to number :", err)
	}
}
