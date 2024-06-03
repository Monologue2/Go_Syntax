package main

import "fmt"

// Var 선언 블록
var (
	int_a    int
	int_b    int
	string_a string
)

// Const 선언 및 정의 블록
const (
	const_int    = 10000
	const_string = "상수 선언"
)

func main() {
	// 변수 선언
	var int_c int
	var string_b string

	// 변수 정의
	int_a = 10
	int_b, int_c = 20, 30

	// 축약형 :=, func 내부에서만 사용 가능, const 선언 불가능
	// 선언과 정의를 동시에 함
	int_d := 40
	// 축약형과 동일
	var int_e = 50

	fmt.Println(int_a, int_b, int_c, int_d, int_e)

	string_a = "Hello"
	string_b = "World"
	// 축약형, 타입 추론
	string_c := "!!!"

	fmt.Println(string_a, string_b, string_c)

	fmt.Println(const_int, const_string)
}
