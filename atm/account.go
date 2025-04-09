package main

type Account struct {
	Name      string
	AccountId int
	Card      Card
	Amount    float64
}

func (a *Account) getCardDetails() Card {
	return a.Card
}

func (a *Account) getAccountId() int {
	return a.AccountId
}

func (a *Account) getAccountName() string {
	return a.Name
}

func (a *Account) getAmount() float64 {
	return a.Amount
}

type Card struct {
	CardNumber string
	CardPin    string
}

func (a *Account) initCard(card Card) *Card {
	return &Card{
		CardNumber: card.CardNumber,
		CardPin:    card.CardPin,
	}
}
