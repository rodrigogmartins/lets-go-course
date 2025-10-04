package main

import "fmt"

func Pointers() {
	x := 45
	p := &x
	fmt.Printf("Pointer to X: %v\n", *p) // 45

	// Modify through pointer
	*p = 100
	fmt.Printf("Pointer modification directly: %v\n", x) // 100

	// Modify as collateral effect
	fmt.Printf("Pointer before modification: %v\n", *p)
	fmt.Println("incr(p)")
	incr(p)
	fmt.Printf("Pointer after modification: %v\n", *p)
}

func incr(p *int) {
	*p++
}
