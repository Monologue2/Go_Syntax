package main

import "fmt"

// 기본 함수 선언 및 정의
func add(a int, b int) int {
	return a + b
}

// Parameter Type 축약
func sub(a, b int) int {
	return a - b
}

// return 다수
func add_and_sub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	a, b := 20, 10

	fmt.Println("add : ", add(a, b))
	fmt.Println("sub : ", sub(a, b))

	add_result, sub_result := add_and_sub(a, b)
	fmt.Println("add_and_sub : ", add_result, sub_result)

	// 익명 함수
	func(a, b int) {
		fmt.Println("익명 함수 : ", a+b)
	}(a, b)
}
