package t

import "strconv"

type Card struct {
	suit  string
	value string
}

func (card Card) GetValue() int {
	v := 0
	switch card.value {
	case "Es":
		v = 1
	case "Knægt":
		v = 10
	case "Dame":
		v = 10
	case "Konge":
		v = 10
	default:
		if i, err := strconv.Atoi(card.value); err == nil {
			v = i
		}
	}
	return v
}

func (card Card) GetSuit() string {
	return card.suit
}

func NewCard(suit, value string) Card {
	return Card{
		suit:  suit,
		value: value,
	}
}
