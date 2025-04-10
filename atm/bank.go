package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
*	Banks responsibility
	1. GetBankName
	2. AddAccount
	3. GetAccount
	4. GetAccountByCard
	5. ListAccounts
*/

// interface for bank operations
type BankOperations interface {
	GetBankName() string
	AddAccount(accountDetails Account) error
	GetAccount(accountId int) (AccountOperations, error)
	GetAccountByCard(cardNumber string) (AccountOperations, error)
	ListAccounts() []AccountOperations
}

type Bank struct {
	Name           string
	ListOfAccounts []AccountOperations
	mu             sync.RWMutex // Read-Write lock
}

func InitBank(name string) *Bank {
	return &Bank{
		Name:           name,
		ListOfAccounts: make([]AccountOperations, 0),
	}
}

func (b *Bank) GetBankName() string {
	return b.Name
}

func (b *Bank) AddAccount(accountDetails Account) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if accountDetails.Name == "" {
		return errors.New("account name cannot be empty")
	}

	accountDetails.AccountId = len(b.ListOfAccounts) + 1
	b.ListOfAccounts = append(b.ListOfAccounts, &accountDetails)
	return nil
}

func (b *Bank) GetAccount(accountId int) (AccountOperations, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, account := range b.ListOfAccounts {
		if account.GetAccountId() == accountId {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

func (b *Bank) GetAccountByCard(cardNumber string) (AccountOperations, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, account := range b.ListOfAccounts {
		if account.GetCardDetails().CardNumber == cardNumber {
			return account, nil
		}
	}
	return nil, errors.New("card not found")
}

func (b *Bank) ListAccounts() []AccountOperations {
	b.mu.RLock()
	defer b.mu.RUnlock()

	accounts := make([]AccountOperations, len(b.ListOfAccounts))
	copy(accounts, b.ListOfAccounts)
	return accounts
}

func (b *Bank) PrintAccountDetails() {
	b.mu.RLock()
	defer b.mu.RUnlock()

	fmt.Printf("Bank: %s\n", b.Name)
	fmt.Println("Accounts:")
	for _, account := range b.ListOfAccounts {
		fmt.Printf("- %s\n", account)
		if account.GetCardDetails().CardNumber != "" {
			fmt.Printf("  Card: %s\n", account.GetCardDetails().CardNumber)
		}
	}
}
