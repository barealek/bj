package main

import (
	"fmt"
	"os"

	t "github.com/barealek/bj/types"
	"github.com/barealek/bj/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var player *t.Player
var dealer *t.Dealer

func ccmodel() bool {
	shallContinue := true
	var m = ui.InitialCCModel(&shallContinue)
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return shallContinue
}

func blackjack() int {

}

func main() {
	if !ccmodel() {
		return
	}

	player = t.CreatePlayer()
	dealer = t.CreateDealer()
	dealer.DealSelf()
	player.Hit(dealer.Deal())
	dealer.DealSelf()
	player.Hit(dealer.Deal())

	fmt.Println(dealer.RenderDealer())
	fmt.Println(player.RenderPlayer())

}
