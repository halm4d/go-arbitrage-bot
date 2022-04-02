package app

// A simple example that shows how to send messages to a Bubble Tea program
// from outside the program using Program.Send(Msg).

import (
	"fmt"
	"github.com/halm4d/go-arbitrage-bot/src/arb"
	"github.com/halm4d/go-arbitrage-bot/src/client"
	"github.com/halm4d/go-arbitrage-bot/src/constants"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	spinnerStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	helpStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(1, 0)
	dotStyle          = helpStyle.Copy().UnsetMargins()
	profitableStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("36"))
	unProfitableStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	durationStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(0, 0, 0, 2)
	appStyle          = lipgloss.NewStyle().Margin(1, 2, 0, 2)
)

type ResultMsg struct {
	Duration          time.Duration
	MostProfitableArb *arb.Arbitrage
}

func (r ResultMsg) String() string {
	if r.Duration == 0 {
		return dotStyle.Render(strings.Repeat(".", 30))
	}
	if r.MostProfitableArb.ProfitPercentage > 0 {
		return fmt.Sprintf("%s %s", profitableStyle.Render(r.MostProfitableArb.GetRouteString()),
			durationStyle.Render(r.Duration.String()))
	} else {
		return fmt.Sprintf("%s %s", unProfitableStyle.Render(r.MostProfitableArb.GetRouteString()),
			durationStyle.Render(r.Duration.String()))
	}
}

type model struct {
	spinner  spinner.Model
	results  []ResultMsg
	quitting bool
}

func newModel() model {
	const numLastResults = 15
	s := spinner.New()
	s.Style = spinnerStyle
	return model{
		spinner: s,
		results: make([]ResultMsg, numLastResults),
	}
}

func (m model) Init() tea.Cmd {
	return spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	case ResultMsg:
		m.results = append(m.results[1:], msg)
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m model) View() string {
	var s string

	if m.quitting {
		s += "That’s all for today!"
	} else {
		s += m.spinner.View() + " Calculating most profitable arbs..."
	}

	s += "\n\n"

	for _, res := range m.results {
		s += res.String() + "\n"
	}

	if !m.quitting {
		s += helpStyle.Render("Press any key to exit")
	}

	if m.quitting {
		s += "\n"
	}

	return appStyle.Render(s)
}

func RunTermUI() {
	constants.BasePrice = 100
	constants.Fee = .75
	rand.Seed(time.Now().UTC().UnixNano())

	p := tea.NewProgram(newModel())

	go func() {
		symbols := arb.NewSymbols()
		symbols.Init(client.GetExchangeInfo())

		arbs := arb.New(symbols)
		fmt.Printf("Found arbs: %v\n", len(*arbs))
		client.RunWebSocket(symbols, func(bt *arb.BookTickers) {
			go func() {
				for {
					time.Sleep(time.Millisecond * 500)
					startOfCalculation := time.Now()
					bt.MU.Lock()
					var cbt = make(arb.BookTickerMap)
					for key, value := range bt.CryptoBookTickers {
						cbt[key] = value
					}
					var ubt = make(arb.BookTickerMap)
					for key, value := range bt.USDTBookTickers {
						ubt[key] = value
					}
					bt.MU.Unlock()
					pr := arbs.CalculateProfits(&cbt, &ubt)
					p.Send(ResultMsg{MostProfitableArb: pr.GetBestRoute(), Duration: time.Now().Sub(startOfCalculation)})
				}
			}()
		})
	}()

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
