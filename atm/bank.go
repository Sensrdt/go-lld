package main

import (
	"fmt"
)

/*
*	Banks responsibility
	1. Manage accounts
	2. Manage Cards
	3. Get card details
	4. Set card details (if user provides new details)
*/

type Bank struct {
	Name           string
	ListOfAccounts []Account
}

func InitBank(Name string) *Bank {
	return &Bank{
		Name:           Name,
		ListOfAccounts: make([]Account, 0),
	}
}

func (b *Bank) GetBankName() string {
	return b.Name
}

func (b *Bank) AddAccount(accountDetails Account) {
	accountDetails.AccountId = len(b.ListOfAccounts) + 1
	b.ListOfAccounts = append(b.ListOfAccounts, accountDetails)
}

func (b *Bank) ListOfBankAccount() {
	for _, val := range b.ListOfAccounts {
		fmt.Printf("Name: %s -  Id: %d \n", val.getAccountName(), val.getAccountId())
		if val.Card.CardNumber != "" {
			fmt.Printf("Card Number: %s -  Card Pin: %s \n", val.getCardDetails().CardNumber, val.getCardDetails().CardPin)
		} else {
			fmt.Printf("No card present for this account \n")
		}
	}
}

func (b *Bank) addCard(accountId int, card Card) {
	for idx, account := range b.ListOfAccounts {
		if accountId == account.getAccountId() {
			newCard := account.initCard(card)
			b.ListOfAccounts[idx].Card = *newCard
		}
	}
}

type CheckBalanceResponse struct {
	CardNumber string
	Amount     float64
}

func (b *Bank) checkBalance(card Card) CheckBalanceResponse {
	for _, account := range b.ListOfAccounts {
		if card.CardNumber == account.getCardDetails().CardNumber && card.CardPin == account.getCardDetails().CardPin {
			return CheckBalanceResponse{
				CardNumber: account.getCardDetails().CardNumber,
				Amount:     account.Amount,
			}
		}
	}
	fmt.Println("No card found")
	return CheckBalanceResponse{}
}
