package main

import "fmt"

type Operation interface {
	SelectProduct(p *Product)
	InsertNote(note note)
	InsertCoint(coin coin)
	DispenseProduct()
	DispenseChange()
}

// Idle states
type IdleState struct {
	machine *Machine
}

func (i *IdleState) DispenseChange() {
	println("select product")
}
func (i *IdleState) DispenseProduct() {
	println("select product")
}
func (i *IdleState) InsertCoint(coin coin) {
	println("select product")
}
func (i *IdleState) InsertNote(note note) {
	println("select product")
}
func (i *IdleState) SelectProduct(p *Product) {
	fmt.Println("Selecting a product", p.ID)
	if i.machine.Inventory.IsProductAvailable(p) {
		i.machine.SelectProduct = p
		i.machine.SetState(i.machine.ReadyState)
	}
}

// Ready states
type ReadyState struct {
	machine *Machine
}

func (r *ReadyState) DispenseChange() {
	println("product selected")
}
func (r *ReadyState) DispenseProduct() {
	println("product selected")
}
func (r *ReadyState) InsertCoint(coin coin) {
	fmt.Println("Inserting coin", coin)
	r.machine.TotalPayment += float64(coin)
	r.checkPaymentStatus()
}
func (r *ReadyState) InsertNote(note note) {
	fmt.Println("Inserting note", note)
	r.machine.TotalPayment += float64(note)
	r.checkPaymentStatus()
}
func (r *ReadyState) SelectProduct(p *Product) {
	println("product selected")
}

func (s *ReadyState) checkPaymentStatus() {
	fmt.Println("Checking payment", s.machine.TotalPayment)
	if s.machine.TotalPayment >= s.machine.SelectProduct.Price {
		s.machine.SetState(s.machine.DispenseState)
	}
}

// Dispense state
type DispenseState struct {
	machine *Machine
}

func (d *DispenseState) DispenseChange() {
	println("dispensing product")
}
func (d *DispenseState) DispenseProduct() {
	fmt.Println("Product: ", d.machine.SelectProduct.ID)
	d.machine.SetState(d.machine.DispenseChangeState)
	// d.machine.Inventory.ReduceProductCount(d.machine.SelectProduct.ID)
}
func (d *DispenseState) InsertCoint(coin coin) {
	println("dispensing product")
}
func (d *DispenseState) InsertNote(note note) {
	println("dispensing product")
}
func (d *DispenseState) SelectProduct(p *Product) {
	println("dispensing product")
}

// Dispense change state
type DispenseChangeState struct {
	machine *Machine
}

func (d *DispenseChangeState) DispenseChange() {
	totalPayment := d.machine.TotalPayment
	actualPrice := d.machine.SelectProduct.Price

	change := totalPayment - actualPrice

	if change > 0 {
		fmt.Println("Change ", change)
	} else {
		fmt.Println("No change to return")
	}

	d.machine.SetState(d.machine.IdleState)
	d.machine.TotalPayment = 0
	d.machine.SelectProduct = nil
}
func (d *DispenseChangeState) DispenseProduct() {
	println("dispensing change")
}
func (d *DispenseChangeState) InsertCoint(coin coin) {
	println("dispensing change")
}
func (d *DispenseChangeState) InsertNote(note note) {
	println("dispensing change")
}
func (d *DispenseChangeState) SelectProduct(p *Product) {
	println("dispensing change")
}
