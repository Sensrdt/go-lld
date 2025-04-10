package main

import (
	"errors"
	"sync"
)

type ATMOperations interface {
	InsertCard(card Card, pin string) error
	CheckBalance() (float64, error)
	WithdrawMoney(amount float64) error
	ChangePin(oldPin, newPin string) error
	ListBanks() []string
}

type ATM struct {
	Banks          []BankOperations
	currentCard    *Card
	currentBank    BankOperations
	currentAccount *Account
	mu             sync.Mutex
}

func InitATM(banks []BankOperations) *ATM {
	return &ATM{
		Banks: banks,
	}
}

func (a *ATM) ListBanks() []string {
	bankNames := make([]string, len(a.Banks))
	for i, bank := range a.Banks {
		bankNames[i] = bank.GetBankName()
	}
	return bankNames
}

func (a *ATM) InsertCard(card Card, pin string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.currentCard = nil
	a.currentBank = nil
	a.currentAccount = nil

	for _, bank := range a.Banks {
		if account, err := bank.GetAccountByCard(card.CardNumber); err == nil {
			if !account.ValidatePin(pin) {
				return errors.New("invalid PIN")
			}
			a.currentCard = &card
			a.currentBank = bank
			// Type Assertion: Converting interface -> concrete type
			// account is of type AccountOperations (interface) defined in account.go
			// The Account struct implements all AccountOperations interface methods:
			// - GetCardDetails() Card
			// - GetAccountId() int
			// - GetBalance() float64
			// - ValidatePin(pin string) bool
			// - UpdateBalance(amount float64) error
			// Because Account implements all methods, Go allows converting the interface
			// to the concrete *Account type using type assertion account.(*Account)
			// If account doesn't contain an *Account, this will panic at runtime
			a.currentAccount = account.(*Account)
			return nil
		}
	}
	return errors.New("card not recognized")
}

func (a *ATM) CheckBalance() (float64, error) {
	if err := a.validateSession(); err != nil {
		return 0, err
	}
	return a.currentAccount.GetBalance(), nil
}

func (a *ATM) WithdrawMoney(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if err := a.validateSession(); err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}

	return a.currentAccount.UpdateBalance(-amount)
}

func (a *ATM) ChangePin(oldPin, newPin string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if err := a.validateSession(); err != nil {
		return err
	}

	if !a.currentAccount.ValidatePin(oldPin) {
		return errors.New("invalid current PIN")
	}

	if len(newPin) != 4 {
		return errors.New("PIN must be 4 digits")
	}

	a.currentAccount.Card.CardPin = newPin
	return nil
}

func (a *ATM) validateSession() error {
	if a.currentCard == nil || a.currentBank == nil || a.currentAccount == nil {
		return errors.New("no card inserted or invalid session")
	}
	return nil
}

func (a *ATM) EndSession() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.currentCard = nil
	a.currentBank = nil
	a.currentAccount = nil
}
