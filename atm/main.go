package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var names = []string{"Sridip", "John", "Alex", "Carry", "Emily", "Max", "Leela", "Sophia", "Jake", "Liam"}

func GenerateBanks(numBanks int) []*Bank {
	rand.Seed(time.Now().UnixNano())

	var banks []*Bank
	for i := 1; i <= numBanks; i++ {
		bank := InitBank("Bank_" + strconv.Itoa(i))
		numUsers := rand.Intn(100) + 1 // 1–10 users per bank

		for j := 0; j < numUsers; j++ {
			name := names[rand.Intn(len(names))] + " " + fmt.Sprintf("User%d", rand.Intn(1000))
			account := Account{
				Name: name,
				Card: GenerateCard(),
			}
			bank.AddAccount(account)
		}
		banks = append(banks, bank)
	}
	return banks
}

func GenerateCard() Card {
	return Card{
		CardNumber: fmt.Sprintf("%03d-%03d-%04d-%03d", rand.Intn(999), rand.Intn(999), rand.Intn(9999), rand.Intn(999)),
		CardPin:    fmt.Sprintf("%04d", rand.Intn(10000)),
	}
}

func main() {
	/*
		Requirements
		1. Bank
			Accounts (Card associated)
		2. Atm (Contain list of Banks)
			WithdrawAmount
			CheckBalance
			ChangePin
	*/

	bank := InitBank("SBI")
	bank.AddAccount(Account{
		Name: "Sridip Dutta",
	})
	bank.AddAccount(Account{
		Name: "John Doe",
	})
	bank.AddAccount(Account{
		Name: "Alex Leela",
	})
	bank.AddAccount(Account{
		Name: "Carry Hedge",
	})

	card := Card{
		CardNumber: "123-345-4561-123",
		CardPin:    "1234",
	}

	bank.addCard(2, card)

	bank.ListOfBankAccount()

	/*
		listOfBanks := GenerateBanks(10)
			for _, bank := range listOfBanks {
				bank.ListOfBankAccount()
			}
	*/

	var listOfBanks []*Bank
	listOfBanks = append(listOfBanks, bank)

	atm := InitATM(listOfBanks)

	var _ (AtmOperation) = (*ATM)(nil)

	fmt.Println(atm.ListOfBanksAvailable())

	/*
		card := Card{
			CardNumber: "825-879-8155-152",
			CardPin:    "8652",
		}
	*/

	pin := "1234"
	fmt.Println(atm.InsertCard(card, pin))
}
