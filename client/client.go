package client

import (
	"strconv"

	"github.com/barealek/bj/consts"
	t "github.com/barealek/bj/types"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BJModel struct {

	// Blackjack
	player *t.Player
	dealer *t.Dealer

	// Bubbletea
	viewport viewport.Model

	winner int
}

func CreateBJModel() *BJModel {
	const width = 50
	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#888800"))

	bjModel := &BJModel{
		player: t.CreatePlayer(),
		dealer: t.CreateDealer(),

		viewport: vp,
	}

	bjModel.dealer.DealSelf()
	bjModel.player.Hit(bjModel.dealer.Deal())
	bjModel.dealer.DealSelf()
	bjModel.player.Hit(bjModel.dealer.Deal())

	return bjModel
}

func (m *BJModel) Init() tea.Cmd {
	return nil
}

func (m *BJModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			if m.winner != 0 {
				return m, nil
			}

			m.player.Hit(m.dealer.Deal())
			m.checkWinner()

			m.View()
			return m, nil

		case "s":
			if m.winner != 0 {
				return m, nil
			}

			m.dealer.Reveal()
			m.checkWinner()
			m.View()
			return m, nil

		case "r":
			if m.winner == 0 {
				return m, nil
			}
			m.player.ResetScore()
			m.dealer.ResetScore()
			m.winner = consts.NoWinner
			m.dealer.DealSelf()
			m.player.Hit(m.dealer.Deal())
			m.dealer.DealSelf()
			m.player.Hit(m.dealer.Deal())

			return m, nil

		case "esc":
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		return m, nil
	}
}

func (m *BJModel) View() string {
	var s string

	s += lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0")).Render("					Blackjack") // ;)

	s += "\n\n\n Dealers hånd (" + strconv.Itoa(m.dealer.GetCardValue()) + ")\n" + "	" + m.dealer.Render()
	s += "\n\n\n\n Spillers hånd (" + strconv.Itoa(m.player.GetCardValue()) + ")\n	" + m.player.Render()

	n := "h‣Hit, s‣Stand, Esc‣Forlad"
	if m.winner != 0 {
		n = "r‣Restart, Esc‣Forlad"
	}

	w := "\n\n\n\n\n "
	if m.winner == consts.PlayerWins {
		w += "Du vandt! 🤑💸💸💸"
	} else if m.winner == consts.DealerWins {
		w += "Dealeren vandt.. 😒"
	} else if m.winner == consts.Push {
		w += "Push! "
	}
	s += w

	s += "\n\n\n\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#777")).Render(n)

	m.viewport.SetContent(s)
	return m.viewport.View()
}

func (m *BJModel) checkWinner() {
	if m.player.GetTotalValue() > 21 {
		m.winner = consts.DealerWins
		return
	}

	if m.dealer.IsRevealed() {
		if m.player.GetAltTotalValue() == 21 || m.dealer.GetTotalValue() > 21 || m.player.GetTotalValue() > m.dealer.GetTotalValue() {
			m.winner = consts.PlayerWins
			return
		} else if m.dealer.GetTotalValue() == m.player.GetTotalValue() {
			m.winner = consts.Push
			return
		} else {
			m.winner = consts.DealerWins
			return
		}
	}

	m.winner = consts.NoWinner
}
