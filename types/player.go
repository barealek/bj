package t

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

func (player *Player) Render() string {
	r := ""
	for _, card := range player.cards {
		r += card.Colored()
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

func (p *Player) GetCardValue() int {
	if p.altTotalValue > 0 && p.totalValue <= 21 {
		return p.altTotalValue
	}
	return p.totalValue
}

func (p *Player) GetTotalValue() int {
	return p.totalValue
}

func (p *Player) GetAltTotalValue() int {
	return p.altTotalValue
}
