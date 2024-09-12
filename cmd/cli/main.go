package main

import (
	"fmt"
	"os"

	"github.com/barealek/bj/client"
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
	bj := client.CreateBJModel()

	if _, err := tea.NewProgram(bj, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return 0
}

func main() {
	if !ccmodel() {
		return
	}

	blackjack()

}
