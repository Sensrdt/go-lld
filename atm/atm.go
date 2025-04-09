package main

import "sync"

type AtmOperation interface {
	CheckBalance() float64
	WithdrawMoney() float64
	ChangePin(int32)
	InsertCard(card Card, pin string) bool
}

type ATM struct {
	AtmOperation
	Banks []*Bank
}

func InitATM(banks []*Bank) *ATM {
	return &ATM{
		Banks: banks,
	}
}

func (a *ATM) ListOfBanksAvailable() []string {
	response := []string{}
	for _, val := range a.Banks {
		response = append(response, val.GetBankName())
	}
	return response
}

func (a *ATM) InsertCard(card Card, pin string) bool {
	var wg sync.WaitGroup
	resultChan := make(chan bool)
	var once sync.Once
	found := false
	for _, banks := range a.Banks {
		wg.Add(1)
		go func(banks *Bank) {
			defer wg.Done()
			for _, accounts := range banks.ListOfAccounts {
				if accounts.getCardDetails().CardNumber == card.CardNumber && accounts.getCardDetails().CardPin == pin {

					once.Do(func() {
						resultChan <- true
					})
					return
				}
			}
		}(banks)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for val := range resultChan {
		found = val
		break
	}

	return found

}

func (a *ATM) CheckBalance() float64 {

	return 0.0
}

func (a *ATM) WithdrawMoney() float64 {
	return 0.0
}

func (a *ATM) ChangePin(newPin int32) {
}
