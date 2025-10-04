package main

import (
	"fmt"
)

func DataStrucutres() {
	// Array declaration
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Printf("Array[0]: %d.\n", numbers[0])

	// Declaration
	slice := []int{1, 2, 3}
	fmt.Printf("Slice: %v\n", slice) // [1 2 3]

	// Add element
	slice = append(slice, 4)
	fmt.Printf("Slice: %v\n", slice) // [1 2 3 4]

	// Slice from array
	array := [5]int{1, 2, 3, 4, 5}
	subSlice := array[1:4] // [2 3 4]
	fmt.Printf("Slice from Array: %v\n", subSlice)

	// Map
	m := make(map[string]int)
	m["age"] = 45
	m["price"] = 99
	fmt.Printf("Map: %v\n", m) // map[age:42 price:99]

	// Check existence
	value, exists := m["age"]
	if exists {
		fmt.Printf("Map['age']: %v\n", value)
	}

	// Delete
	delete(m, "price")
	fmt.Printf("Map without 'price' key: %v\n", m) // map[age:42]
}
