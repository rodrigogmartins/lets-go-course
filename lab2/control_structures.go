package main

import "fmt"

func ControlStructures() {
	// Use defer to close files, db conections or free mutexes
	defer fmt.Println("Executed at the end")
	fmt.Println("Executed first")

	number := 45

	// If with direct initialization
	if x := number % 2; x == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Traditional
	for i := 0; i < 3; i++ {
		fmt.Printf("Traditional for: %d.\n", i)
	}

	// As while loop
	count := 0
	for count < 3 {
		fmt.Printf("For as while loop: %d.\n", count)
		count++
	}

	// Infinity (com break)
	for {
		fmt.Println("Infinity")
		break
	}

	// Range
	numbers := []int{1, 2, 3}

	fmt.Println("Numbers:")
	for index, number := range numbers {
		fmt.Printf("For range. index: %d, value: %d\n", index, number)
	}

	// Switch with multi-conditions
	day := 3
	switch day {
	case 1, 2:
		fmt.Println("Start of the week")
	case 3, 4, 5:
		fmt.Println("Mid of the week")
	default:
		fmt.Println("Weekend")
	}

	// Switch with expression
	switch x := day * 2; x {
	case 6:
		fmt.Println("Day 3 double")
	default:
		fmt.Println("Other value")
	}

}
