package main

import (
	"fmt"
	"math"
)

// 타입 제한자와 인터페이스는 다른 개념이다.
// Define a basic interface
type Describer interface {
	Describe() string
	// Structure에 Method를 모두 구현하지 않으면 Implement 하지 않은 것으로 간주한다.
	// Greet()
}

// Implementing multiple interfaces
type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
	Age  int
}

// Implement the desctiber interface for person
func (p Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age : %d\n", p.Name, p.Age)
}

// Implementing multiple interfaces
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, %s!", p.Name)
}

// Structure에 Method를 모두 구현하지 않으면 Implement 하지 않은 것으로 간주한다.
// func (p Person) Greet() {
// 	fmt.Printf("Hello, My name is %s, I'm %d years old.", p.Name, p.Age)
// }

// Using interface as function parameters
func PrintDescription(d Describer) {
	fmt.Print(d.Describe())
}

// Empty interface
// type any = interface{}
// %T -> Type , %v -> value
func type_assertion_Description(a any) {
	fmt.Printf("Type: %T, Value: %v\n", a, a)
}

// Composing Interfaces
type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Document struct {
	Content string
}

func (d Document) Read() string {
	return d.Content
}

func (d *Document) Write(s string) {
	d.Content = s
}

// Shape interface
// Define as a interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct for circle
type Circle struct {
	Radius float64
}

// Implement the shape interface for circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Define a struct for rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the shape interface for rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	person := Person{Name: "Shogle", Age: 27}
	var d Describer = person
	var g Greeter = person
	fmt.Print(d.Describe())
	fmt.Print(g.Greet())

	//Using interface as function parameter
	PrintDescription(person)

	//Empty interface
	type_assertion_Description(1234)
	type_assertion_Description("HEII0")
	type_assertion_Description(3.1425)

	fmt.Println("-----Type Assertion-----")
	var i interface{} = "hello"

	// Type assertion.
	// 빈 인터페이스 또는 any 를 사용할 때 유용하게 사용합니다.
	// switch, case 문으로 타입 별 대응
	// if 문으로 원하는 타입인지 아닌지 체크
	// if the assertion fails,
	// a runtume pacin occurs unless you use the ", ok" idiom to handle the failure gracefully.
	s, ok := i.(string)
	if ok {
		fmt.Printf("string value : %s\n", s)
	} else {
		fmt.Println("Type assertion failed")
	}

	// Type switch
	switch v := i.(type) {
	case string:
		fmt.Println("String :", v)
	case int:
		fmt.Println("Integer :", v)
	default:
		fmt.Println("Unknown type")
	}

	// Composing Interface
	doc := &Document{}
	var rw ReadWriter = doc
	rw.Write("Hello, World")
	fmt.Println(rw.Read())

	// Test the shape interface
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 20},
	}

	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f, Perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}
}
