package t

import "fmt"

type Player struct {
	cards         []Card
	totalValue    int
	altTotalValue int
}

func CreatePlayer() *Player {
	player := Player{
		cards:         []Card{},
		totalValue:    0,
		altTotalValue: 0,
	}

	return &player
}

func (player *Player) RenderPlayer() string {
	r := ""
	r += fmt.Sprintf("\r\n===== SPILLER =====\n")
	for _, card := range player.cards {
		r += fmt.Sprintf("[ %s %s ]", card.suit, card.value)
	}

	if player.altTotalValue > 0 {
		r += fmt.Sprintf("\r\nVærdi: %d\r\n", player.altTotalValue)
	} else {
		r += fmt.Sprintf("\r\nVærdi: %d\r\n", player.totalValue)
	}

	return r
}

func (player *Player) Hit(card Card) {
	player.cards = append(player.cards, card)
	if card.GetValue() == 1 && player.totalValue+11 < 21 {
		player.altTotalValue = player.totalValue + 11
	} else {
		player.altTotalValue = 0
	}
	player.totalValue += card.GetValue()
}

func (player *Player) ResetScore() {
	player.cards = []Card{}
	player.totalValue = 0
	player.altTotalValue = 0
}
