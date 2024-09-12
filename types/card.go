package t

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

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

func (card Card) Colored() string {
	if card.suit == "♠" || card.suit == "♥" {
		return "[ " + lipgloss.NewStyle().Foreground(lipgloss.Color("#c55")).Render(card.suit+" "+card.value) + " ]"
	} else {
		return "[ " + lipgloss.NewStyle().Foreground(lipgloss.Color("#fff")).Render(card.suit+" "+card.value) + " ]"
	}
}
