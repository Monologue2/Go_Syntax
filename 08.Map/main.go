package main

import (
	"fmt"
)

func main() {
	// Basic map operations
	var m map[string]int = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m)
	// adding elements to the map
	m["three"] = 3
	m["four"] = 4
	fmt.Println("Modified Map :", m)

	// Accessing and modifying elements
	element01 := m["one"]
	fmt.Println("accessing element01 :", element01)

	// modifying elements
	m["one"] = 10001
	fmt.Println("Update one to 10 thousand :", m["one"])

	//accessing a non-exsitent key
	value, exists := m["five"]
	if exists {
		fmt.Println("five is ... :", value)
	} else {
		fmt.Println("Five not found in the map")
	}

	// Deleting elements
	delete(m, "four")
	fmt.Println("Map after deleing four : ", m)

	// Iterating over a map
	fmt.Println("Iterating over a mapm using range")
	for key, value := range m {
		fmt.Printf("Key : %s, Value : %d\n", key, value)
	}

	// Map of slice
	s_m := map[string][]string{
		"fruits":  {"apple", "banana", "mango"},
		"animals": {"dog", "cat", "elephant"},
	}
	for key, items := range s_m {
		fmt.Printf("%s: %v\n", key, items)
	}

	//Nested maps
	nestedMap := map[string]map[string]int{
		"TeamA": {
			"Alice": 1,
			"Tom":   2,
		},
		"TeamB": {
			"Shogle": 1,
			"Mum":    2,
		},
	}
	for team, members := range nestedMap {
		fmt.Printf("Nested Map, Team : %s\n", team)
		for name, number := range members {
			fmt.Printf("name : %s, Number : %d\n", name, number)
		}
	}
	fmt.Println(nestedMap["TeamA"]["Alice"])

	//Using structure as a map value
	people := map[string]Person{
		"Shogle": {Age: 27, City: "Busan"},
		"Mum":    {Age: 27, City: "Daejun"},
	}
	for name, person := range people {
		fmt.Printf("Name : %s, Age : %d, City : %s\n", name, person.Age, person.City)
	}

	//Using Interfaces as map keys or values
	// 인터페이스 타입
	entities := map[string]Describer{
		"p1": Person{City: "Busan", Age: 27},
		"a1": Animal{Species: "Dog"},
	}

	for key, entity := range entities {
		fmt.Printf("key : %s, Description; %s\n", key, entity.Describe())
	}

	// Using maps for frequency count
	text := "Hello world"
	freqMap := make(map[rune]int)
	fmt.Println(freqMap)
	for _, char := range text {
		// freqMap[char]++
		freqMap[char] = freqMap[char] + 1
	}
	//printing frequency map
	fmt.Println("Frequency of characters : ")
	for char, count := range freqMap {
		fmt.Printf("Characrter: %c, Count: %d\n", char, count)
	}
}

// Using structure as a map value
type Person struct {
	Age  int
	City string
}

// Using interfases as map keys or values
// define an interface
type Describer interface {
	Describe() string
}

// Person type 에 대한 Decriber Interface 구현
func (p Person) Describe() string {
	return "City: " + p.City
}

// define another struct
type Animal struct {
	Species string
}

// Animal type 에 대한 Desciber Interface 구현
func (a Animal) Describe() string {
	return "Animal: " + a.Species
}
