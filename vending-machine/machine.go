package main

import "fmt"

type Machine struct {
	Inventory           *Inventory
	IdleState           Operation
	ReadyState          Operation
	DispenseState       Operation
	DispenseChangeState Operation
	CurrentState        Operation
	SelectProduct       *Product
	TotalPayment        float64
}

func CreateNewMachine() *Machine {
	fmt.Println("Creating new machine")
	machine := &Machine{}
	machine.Inventory = CraeteNewInventory()
	machine.IdleState = &IdleState{machine}
	machine.ReadyState = &ReadyState{machine}
	machine.DispenseState = &DispenseState{machine}
	machine.DispenseChangeState = &DispenseChangeState{machine}
	machine.CurrentState = machine.IdleState
	return machine
}

func (m *Machine) ChooseProduct(p *Product) {
	m.CurrentState.SelectProduct(p)
}

func (m *Machine) InsertNote(n note) {
	m.CurrentState.InsertNote(n)
}

func (m *Machine) InsertCoin(c coin) {
	m.CurrentState.InsertCoint(c)
}

func (m *Machine) DispenseProduct() {
	m.CurrentState.DispenseProduct()
}

func (m *Machine) DispenseChange() {
	m.CurrentState.DispenseChange()
}

func (m *Machine) SetState(state Operation) {
	m.CurrentState = state
}
