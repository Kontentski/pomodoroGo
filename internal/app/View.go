package pomodoro

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kontentski/pomodoroGo/internal/helper"
)

//helper for View when you rendering start menu
func (m Model) menuView(style lipgloss.Style) string {
	s := "\n"
	s += style.Render("POMODORO TIMER") + "\n\n"
	s += style.Render("Choose timer duration:") + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += style.Render(fmt.Sprintf("%s %s ", cursor, choice.label)) + "\n"
	}

	if m.inputting {
		s += style.Render(fmt.Sprintf("\nEnter minutes: %s█", m.input)) + "\n"
	}
	return s
}

//helper for View when you rendering the timer
func (m Model) timerView(style lipgloss.Style, hasSpace bool) string {
	s := "\n"
	s += style.Render("🍅 POMODORING TIME") + "\n"
	s += style.Render("════════════════════") + "\n\n"

	if hasSpace {
		// Big number display
		minutesStr := fmt.Sprintf("%02d", m.minutes)
		secondsStr := fmt.Sprintf("%02d", m.seconds)

		min2 := helper.GetBigNumber(int(minutesStr[1] - '0'))
		min1 := helper.GetBigNumber(int(minutesStr[0] - '0'))
		colonLines := strings.Split(helper.Colon, "\n")[1:]
		sec1 := helper.GetBigNumber(int(secondsStr[0] - '0'))
		sec2 := helper.GetBigNumber(int(secondsStr[1] - '0'))

		// Combine the numbers
		for i := 0; i < len(min1); i++ {
			line := fmt.Sprintf("%s  %s  %s  %s  %s",
				min1[i],
				min2[i],
				colonLines[i],
				sec1[i],
				sec2[i])
			s += style.Render(line) + "\n"
		}

	} else {//small number display if the terminal doesn't have enough space
		s += style.Render(fmt.Sprintf(helper.SmallStyle, m.minutes, m.seconds)) + "\n"
	}

	if m.totalTime > 0 {
		remainingTime := m.minutes*60 + m.seconds
		percent := 1.0 - float64(remainingTime)/float64(m.totalTime)
		s += "\n" + style.Render(m.progress.ViewAs(percent)) + "\n\n"
	}

	s += "\n"

	if !m.isRunning && m.minutes == 0 && m.seconds == 0 {
		s += style.Render("Timer completed!") + "\n"
	} else if m.isPaused {
		s += style.Render("Timer paused") + "\n"
	}

	return s
}

//helper for View, renders the controls
func (m Model) controlsView() string {
	var help string
	if m.inputting { //input view
		help = fmt.Sprintf("%s - confirm • %s - cancel • %s - quit",
			m.keys.Enter.Help().Key,
			m.keys.Esc.Help().Key,
			m.keys.Quit.Help().Key,
		)

	} else if !m.selected { //start menu
		help = fmt.Sprintf("%s - up • %s - down • %s - select • %s - quit",
			m.keys.Up.Help().Key,
			m.keys.Down.Help().Key,
			m.keys.Enter.Help().Key,
			m.keys.Quit.Help().Key,
		)
	} else {
		if m.timerComplete {
			help = fmt.Sprintf("%s - new timer • %s - quit",
				m.keys.Enter.Help().Key,
				m.keys.Quit.Help().Key,
			)
		} else {
			help = fmt.Sprintf("%s - pause • %s - quit",
				m.keys.Pause.Help().Key,
				m.keys.Quit.Help().Key,
			)
		}
	}

	return help
}
