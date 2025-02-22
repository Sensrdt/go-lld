package main

func main() {

	p := NewProduct(1, 220, 4)
	m := CreateNewMachine()
	m.Inventory.AddProduct(p)

	m.ChooseProduct(p)

	m.InsertNote(note(TWO_HUNDERED))
	m.InsertNote(note(TWO_HUNDERED))
	// m.InsertCoin(coin(TEN))
	// m.InsertCoin(coin(TEN))

	m.DispenseProduct()

	m.DispenseChange()

}
