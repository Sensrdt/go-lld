package main

import "fmt"

// add new product to inventory
// get product by id
// remove product from inventory
// update product quantity
// get all products

type Inventory struct {
	Products map[*Product]int
}

// intialize a new inventory
func CraeteNewInventory() *Inventory {
	return &Inventory{Products: make(map[*Product]int)}
}

func (i *Inventory) AddProduct(p *Product) {
	fmt.Println("Adding new product to machine", p.ID, p.Price, p.Quantity)
	i.Products[p] = p.Quantity
}

func (i *Inventory) IsProductAvailable(p *Product) bool {
	return i.Products[p] > 0
}

// func (i *Inventory) ReduceProductCount(id int) bool {
// 	i.Products[p] = i.Products[p] - 1
// 	return true
// }
