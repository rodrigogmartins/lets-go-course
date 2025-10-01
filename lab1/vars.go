package main

import "fmt"

const version = "1.0.0"
const ivaTax = 0.21

var name string = "Rodrigo"

func vars() {
	age := 25
	var price float64 = 99.00
	active := true

	fmt.Printf("Name: %s, Age: %d, Version: %s\n", name, age, version)
	fmt.Printf("Price: %.2f, Active: %t\n", price, active)
	fmt.Printf("Final price: %.2f\n", calculatePriceWithIvaTax(price))

	x, y := swap("Go", "Let's")
	fmt.Printf("Swap: %s, %s!\n", x, y)

	number := 12.5
	double := doubleNumber(number)
	fmt.Printf("Double of: %.2f is: %.2f!\n", number, double)
}

func swap(a, b string) (string, string) {
	return b, a
}

func doubleNumber(n float64) float64 {
	return n * 2
}

func calculatePriceWithIvaTax(price float64) float64 {
	return price * (1 + ivaTax)
}
