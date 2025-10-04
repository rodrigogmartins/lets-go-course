package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

func Structs() {
	// Structs with tags
	p := Product{ID: 1, Name: "Laptop"}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData)) // {"id":1,"name":"Laptop"}
}
