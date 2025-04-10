package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var names = []string{"Sridip", "John", "Alex", "Carry", "Emily", "Max", "Leela", "Sophia", "Jake", "Liam"}

// GenerateCard creates a new card with random number and PIN
func GenerateCard() Card {
	return Card{
		CardNumber: fmt.Sprintf("%03d-%03d-%04d-%03d", rand.Intn(999), rand.Intn(999), rand.Intn(9999), rand.Intn(999)),
		CardPin:    fmt.Sprintf("%04d", rand.Intn(10000)),
	}
}

// GenerateBanks creates a specified number of banks with random accounts
func GenerateBanks(numBanks int) []*Bank {
	rand.Seed(time.Now().UnixNano())

	var banks []*Bank
	for i := 1; i <= numBanks; i++ {
		bank := InitBank("Bank_" + strconv.Itoa(i))
		numUsers := rand.Intn(5) + 1 // 1-5 users per bank

		for j := 0; j < numUsers; j++ {
			name := names[rand.Intn(len(names))] + " " + fmt.Sprintf("User%d", rand.Intn(1000))
			account := Account{
				Name:    name,
				Card:    GenerateCard(),
				Balance: float64(rand.Intn(10000)), // Random initial balance
			}
			if err := bank.AddAccount(account); err != nil {
				log.Printf("Failed to add account: %v", err)
			}
		}
		banks = append(banks, bank)
	}
	return banks
}

func main() {
	// Create some sample banks and accounts
	banks := GenerateBanks(3)

	// Convert []*Bank to []BankOperations
	bankOperations := make([]BankOperations, len(banks))
	for i, bank := range banks {
		bankOperations[i] = bank
	}

	// Initialize ATM with the banks
	atm := InitATM(bankOperations)

	// Print available banks
	fmt.Println("Available Banks:")
	for _, bankName := range atm.ListBanks() {
		fmt.Printf("- %s\n", bankName)
	}
	fmt.Println()

	// Print all accounts for demonstration
	fmt.Println("All Bank Accounts:")
	for _, bank := range banks {
		bank.PrintAccountDetails()
	}

	// Demonstrate ATM operations with the first account
	firstBank := banks[0]
	if len(firstBank.ListAccounts()) == 0 {
		log.Fatal("No accounts available for demonstration")
	}

	firstAccount := firstBank.ListAccounts()[0]
	card := firstAccount.GetCardDetails()

	// ATM Session Example
	fmt.Println("\nATM Session Example:")
	fmt.Printf("Using card: %s\n", card.CardNumber)

	// 1. Insert Card
	if err := atm.InsertCard(card, card.CardPin); err != nil {
		log.Fatalf("Failed to insert card: %v", err)
	}
	fmt.Println("Card accepted!")

	// 2. Check Balance
	if balance, err := atm.CheckBalance(); err != nil {
		log.Printf("Failed to check balance: %v", err)
	} else {
		fmt.Printf("Current balance: INR %.2f\n", balance)
	}

	// 3. Withdraw Money
	withdrawAmount := 50.0
	if err := atm.WithdrawMoney(withdrawAmount); err != nil {
		fmt.Printf("Withdrawal failed: %v\n", err)
	} else {
		fmt.Printf("Successfully withdrew INR %.2f\n", withdrawAmount)
		if balance, err := atm.CheckBalance(); err == nil {
			fmt.Printf("New balance: INR %.2f\n", balance)
		}
	}

	// 4. Change PIN
	newPin := "5678"
	if err := atm.ChangePin(card.CardPin, newPin); err != nil {
		fmt.Printf("Failed to change PIN: %v\n", err)
	} else {
		fmt.Println("PIN successfully changed!")
	}

	// 5. End Session
	atm.EndSession()
	fmt.Println("Session ended")
}
