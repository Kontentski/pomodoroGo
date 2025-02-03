package pomodoro

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kontentski/pomodoroGo/internal/helper"
)

type SessionType struct {
	duration int
	label    string
}

// Model represents the application state
type Model struct {
	choices       []SessionType
	keys          KeyMap
	cursor        int // cursor ">" position
	selected      bool
	minutes       int
	seconds       int
	isRunning     bool
	isPaused      bool
	progress      progress.Model // progress bar
	totalTime     int            // total time for progress bar
	timerType     string         // "work", "break", or "long-break" TODO: if was "work", then suggest a "break"
	inputting     bool           // are we currently inputting custom time?
	input         string         // store custom time input
	timerComplete bool
	muted         bool //for notification
}

type tickMsg time.Time

// builder
func InitialModel() Model {
	return Model{
		choices: []SessionType{
			{duration: 25, label: "25 minutes "},
			{duration: 5, label: " 5 minutes "},
			{duration: 15, label: "15 minutes "},
			{duration: 0, label: "Custom"},
		},
		keys:          DefaultKeyMap(),
		cursor:        0,
		selected:      false,
		minutes:       0,
		seconds:       0,
		isRunning:     false,
		isPaused:      false,
		progress:      progress.New(progress.WithDefaultGradient(), progress.WithoutPercentage()),
		totalTime:     0,
		timerType:     "work",
		inputting:     false,
		input:         "",
		timerComplete: false,
		muted:         false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//Update listens for updates in the app (keypress, timer updates)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputting {
			return m.inputUpdate(msg)
		}

		// start menu controls
		return m.menuUpdate(msg)

	case tickMsg:
		return m.timerUpdate()
	}
	return m, nil
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// View renders the application UI
func (m Model) View() string {
	hasSpace, width, height := helper.HasEnoughSpace()

	style := lipgloss.NewStyle().Width(width).Align(lipgloss.Center, lipgloss.Bottom)

	var content string

	if !m.selected {
		content = m.menuView(style)
	} else {
		content = m.timerView(style, hasSpace)
	}

	//style for the control panel at the bottom of the screen
	helpStyle := lipgloss.NewStyle().
		Width(width).Align(lipgloss.Center).
		Foreground(lipgloss.Color("#6a6a6a")).
		Padding(0, 1)

	muteStyle := lipgloss.NewStyle().AlignHorizontal(lipgloss.Right)
		muteHelp := " • m - mute"
	if m.muted {
        muteHelp = " • m - unmute"
    }

	//the logic fot the control panel to be at the bottom of the screen
	contentHeight := strings.Count(content, "\n") + 1
	paddingHeight := height - contentHeight
	if paddingHeight > 0 {
		content += strings.Repeat("\n", paddingHeight)
	}
	return content + helpStyle.Render(m.controlsView() + muteStyle.Render(muteHelp))
}
