package t

import (
	"math/rand"
)

type Deck []Card

func (deck *Deck) Shuffle() {
	rand.Shuffle(len(*deck), func(i, j int) {
		dereffedDeck := (*deck)
		dereffedDeck[i], dereffedDeck[j] = dereffedDeck[j], dereffedDeck[i]
	})
}

func CreateDeck() *Deck {
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"Es", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Knægt", "Dame", "Konge"}
	var deck Deck = []Card{}

	for _, value := range values {
		for _, suit := range suits {
			deck = append(deck, NewCard(suit, value))
		}
	}
	deck.Shuffle()
	return &deck
}

func (deck *Deck) DealCard() Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	if len(*deck) == 1 {
		*deck = append(*deck, *CreateDeck()...)
	}
	return card
}
