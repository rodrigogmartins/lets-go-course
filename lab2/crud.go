package main

import (
	"errors"
	"fmt"
)

type ProductTest struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []ProductTest
var productsMap map[int]ProductTest
var idCounter = 1

func createProduct(name string, price float64) (ProductTest, error) {
	if price < 0 {
		return ProductTest{}, errors.New("Product price must be greather than 0")
	}

	product := ProductTest{ID: idCounter, Name: name, Price: price}
	products = append(products, product)
	productsMap[idCounter] = product
	idCounter++

	return product, nil
}

func listProducts() []ProductTest {
	return products
}

func findProduct(id int) (ProductTest, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return ProductTest{}, errors.New("Product not found")
}

func findProductInMap(id int) (ProductTest, error) {
	product, exists := productsMap[id]

	if !exists {
		return ProductTest{}, errors.New("Product not found")
	}

	return product, nil
}

func updateProduct(id int, name string, price float64) (ProductTest, error) {
	if price < 0 {
		return ProductTest{}, errors.New("Product price must be greather than 0")
	}

	for i, p := range products {
		if p.ID == id {
			products[i] = ProductTest{ID: id, Name: name, Price: price}
			return products[i], nil
		}
	}

	return ProductTest{}, errors.New("Product not found")
}

func deleteProduct(id int) error {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("Product not found")
}

func CrudProducts() {
	// Map initialization
	productsMap = make(map[int]ProductTest)

	// Creating products
	createProduct("Laptop", 999.99)
	createProduct("Mouse", 29.99)
	fmt.Println()

	fmt.Println("Creating product with price < 0:")
	_, err := createProduct("Mouse", -29.99)
	if err != nil {
		fmt.Println("Error creating product. Error: ", err)
	}
	fmt.Println()

	// Listing products
	fmt.Println("Listing products:")
	for _, p := range listProducts() {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
	fmt.Println()

	// Find product
	fmt.Println("Finding product:")
	if p, err := findProduct(1); err == nil {
		fmt.Printf("Product found: %+v\n", p)
	}

	fmt.Println("Finding product in Map:")
	if p, err := findProductInMap(1); err == nil {
		fmt.Printf("Product found: %+v\n", p)
	}
	fmt.Println()

	// Update product
	if p, err := updateProduct(1, "Laptop Pro", 1299.99); err == nil {
		fmt.Printf("Product updated: %+v\n", p)
	}

	fmt.Println("Updating product with price < 0:")
	_, err = updateProduct(1, "Laptop Pro", -299.99)
	if err != nil {
		fmt.Println("Error updating product. Error: ", err)
	}
	fmt.Println()

	// Delete product
	if err := deleteProduct(2); err == nil {
		fmt.Println("Product deleted with success")
	}
	fmt.Println()

	// Final list
	fmt.Println("Final list:")
	for _, p := range listProducts() {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}
