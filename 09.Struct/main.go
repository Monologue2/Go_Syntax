package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct for embedding
type Address struct {
	City, State string
}

// Define a struct
type Person struct {
	Name string
	Age  int
	// Basic imbedding
	Address
}

// Define methods on structs
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s, and I am %d years old.\n", p.Name, p.Age)
}

// Value receivers, p Person
// When a method has a value receiver, It operates on a copy of the receiver, not the original value
func (p Person) Getter_Name() string {
	return p.Name
}

func (p Person) Getter_Age() int {
	return p.Age
}

// Pointer receivers, p Person
// Define a method with a pointer receiver to modify the struct
func (p *Person) Setter_Age(age int) {
	p.Age = age
}

func (p *Person) Setter_Name(name string) {
	p.Name = name
}

func (p *Person) Happy_Birthday() {
	p.Age++
}

type Employee struct {
	Person  // Embedding person struct
	Company string
}

// Method overriding
func (e Employee) Greet() {
	fmt.Printf("Hello, my name is %s, and I am %d years old, and I work at %s\n", e.Name, e.Age, e.Company)
}

func main() {
	// Initialize a struct
	// var person1 Person = Person{Name: "Shogle", Age: 27}
	// person2 := Person{Name: "Uda", Age: 27}
	// person3 := Person{"Mum", 27}
	var person1 Person = Person{Name: "Shogle", Age: 27, Address: Address{City: "Busan", State: "BS"}}
	person2 := Person{Name: "Uda", Age: 27, Address: Address{City: "Busan", State: "BS"}}
	person3 := Person{"Mum", 27, Address{City: "Busan", State: "BS"}}
	fmt.Println(person1, person2, person3)

	// Accessing and mofifying struct fields
	fmt.Println("Name :", person3.Name)
	fmt.Println("Name :", person3.Age)

	// Modifying fields
	person3.Age = 24
	fmt.Println("Updated Age :", person3.Age)

	// Method overriding
	var employee Employee = Employee{Person: person1, Company: "Shogle Comp"}
	employee.Greet()

	// Using structs with JSON
	jsonData, err := json.Marshal(person3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("JSON data : ", string(jsonData))
	var decoded Person
	// Json Data를 받아 Structure 에 넣을 수 있다.
	// Structure 구조를 Json과 동일하게 만들어둔다면 재사용 가능
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Decoded Person:", decoded)

	// Using structs in map
	people := map[string]Person{
		// 타입 명을 붙이지 않아도 됨
		"Shogle": Person{Name: "Shogle", Age: 27, Address: Address{City: "Busan", State: "BS"}},
		"Mum":    {Name: "Mum", Age: 64272, Address: Address{City: "Deagu", State: "DG"}},
	}
	fmt.Println("Mum :", people["Mum"])
	fmt.Println("Shogle :", people["Shogle"])

	// Using structs in slice
	people_slice := []Person{
		{Name: "Shogle", Age: 25, Address: Address{City: "Busan", State: "BS"}},
		{Name: "Mum", Age: 36, Address: Address{City: "Seoul", State: "SE"}},
	}
	fmt.Println(people_slice[0], people_slice[1])
}
