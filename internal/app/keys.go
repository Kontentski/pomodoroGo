package pomodoro

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Enter   key.Binding
	Quit    key.Binding
	Pause   key.Binding
	Esc     key.Binding
	Mute       key.Binding
	Numbers []key.Binding
}

func DefaultKeyMap() KeyMap {
	numbers := make([]key.Binding, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = key.NewBinding(
			key.WithKeys(string(rune('0'+i))),
			key.WithHelp(string(rune('0'+i)), "number input"),
		)
	}
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k", "K"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j", "J"),
			key.WithHelp("↓/j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("Enter", "select"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c", "Q"),
			key.WithHelp("q", "quit"),
		),
		Pause: key.NewBinding(
			key.WithKeys(" ", "p", "P"),
			key.WithHelp("space/p", "toggle pause/resume"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel custom time input"),
		),
		Mute: key.NewBinding(
			key.WithKeys("m", "M"),
            key.WithHelp("m", "mute the sound"),
		),
		Numbers: numbers,
	}
}
