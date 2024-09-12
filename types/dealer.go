package t

import "fmt"

type Dealer struct {
	deck          *Deck
	cards         []Card
	totalValue    int
	altTotalValue int
	revealed      bool
}

func CreateDealer() *Dealer {
	dealer := &Dealer{
		deck:          CreateDeck(),
		cards:         []Card{},
		totalValue:    0,
		altTotalValue: 0,
		revealed:      false,
	}
	return dealer
}

func (dealer *Dealer) Render() string {
	var r string
	if !dealer.revealed {
		r += dealer.cards[0].Colored()
		r += fmt.Sprintf("[ %s %s ]", "?", "?")
	} else {
		for _, card := range dealer.cards {
			r += card.Colored()
		}

	}

	return r
}

func (dealer *Dealer) DealSelf() {
	card := dealer.deck.DealCard()
	dealer.cards = append(dealer.cards, card)
	if card.GetValue() == 1 && dealer.totalValue+11 <= 21 {
		dealer.altTotalValue = dealer.totalValue + 11
	} else {
		dealer.altTotalValue = 0
	}
	dealer.totalValue += card.GetValue()
}

func (dealer *Dealer) Deal() Card {
	card := dealer.deck.DealCard()
	return card
}

func (dealer *Dealer) Reveal() {
	dealer.revealed = true
	for dealer.totalValue <= 16 {
		dealer.DealSelf()
	}
}

func (dealer *Dealer) GetCardValue() int {
	if dealer.revealed {
		if dealer.altTotalValue > 0 && dealer.totalValue <= 21 {
			return dealer.altTotalValue
		}
		return dealer.totalValue
	}
	if v := dealer.cards[0].GetValue(); v != 1 {
		return v
	}
	return 11
}

func (dealer *Dealer) IsRevealed() bool {
	return dealer.revealed
}

func (dealer *Dealer) GetTotalValue() int {
	return dealer.totalValue
}

func (dealer *Dealer) GetAltTotalValue() int {
	return dealer.altTotalValue
}

func (dealer *Dealer) ResetScore() {
	dealer.cards = []Card{}
	dealer.totalValue = 0
	dealer.altTotalValue = 0
	dealer.revealed = false
}
