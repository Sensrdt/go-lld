package main

import (
	"errors"
	"fmt"
)

type Card struct {
	CardNumber string
	CardPin    string
}

type Account struct {
	Name      string
	AccountId int
	Card      Card
	Balance   float64
}

type AccountOperations interface {
	GetCardDetails() Card
	GetAccountId() int
	GetAccountName() string
	GetBalance() float64
	UpdateBalance(amount float64) error
	ValidatePin(pin string) bool
}

func (a *Account) GetCardDetails() Card {
	return a.Card
}

func (a *Account) GetAccountId() int {
	return a.AccountId
}

func (a *Account) GetAccountName() string {
	return a.Name
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) UpdateBalance(amount float64) error {
	newBalance := a.Balance + amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	a.Balance = newBalance
	return nil
}

func (a *Account) ValidatePin(pin string) bool {
	return a.Card.CardPin == pin
}

func (a *Account) InitCard(card Card) error {
	if card.CardNumber == "" || card.CardPin == "" {
		return errors.New("invalid card details")
	}
	a.Card = card
	return nil
}

// String implements the Stringer interface for pretty printing
func (a *Account) String() string {
	return fmt.Sprintf("Account{Name: %s, ID: %d, Balance: %.2f}", a.Name, a.AccountId, a.Balance)
}
