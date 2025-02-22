package main

import "fmt"

type Product struct {
	ID       int
	Price    float64
	Quantity int
}

// initialize a new product
func NewProduct(id int, price float64, quantity int) *Product {
	fmt.Println("Creating new product", id, price, quantity)
	return &Product{id, price, quantity}
}
