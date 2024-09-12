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

func (dealer *Dealer) RenderDealer() string {
	r := "===== DEALER ======\n"
	if !dealer.revealed {
		r += fmt.Sprintf("[ %s %v ]", dealer.cards[0].GetSuit(), dealer.cards[0].GetValue())
		r += fmt.Sprintf("[ %s %s ]", "?", "?")
	} else {
		for _, card := range dealer.cards {
			r += fmt.Sprintf("[ %s %v ]", card.GetSuit(), card.GetValue())
		}

		r += fmt.Sprintf("\r\nTotal: %d\r\n", dealer.totalValue)

		if dealer.altTotalValue > 0 {
			r += fmt.Sprintf("Alternative Total: %d\r\n", dealer.altTotalValue)
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
