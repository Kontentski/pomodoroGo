package pomodoro

import (
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kontentski/pomodoroGo/internal/helper"
)

//helper for Update when you input custom time
func (m Model) inputUpdate(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, m.keys.Esc):
		m.inputting = false
		m.input = ""
		return m, nil
	case key.Matches(msg, m.keys.Enter):
		// Convert input to minutes
		minutes, err := strconv.Atoi(m.input)
		if err == nil && minutes > 0 {
			m.minutes = minutes
			m.totalTime = minutes * 60
			m.selected = true
			m.inputting = false
			m.isRunning = true
			m.isPaused = false
			return m, tick()
		}
	case msg.Type == tea.KeyBackspace:
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
		}
	default:
		// Only accept numeric input
		if len(msg.String()) == 1 && msg.String() >= "0" && msg.String() <= "9" {
			m.input += msg.String()
		}
	}
	return m, nil
}

//helper for Update when ypu on start menu
func (m Model) menuUpdate(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit
	case key.Matches(msg, m.keys.Up):
		if !m.selected {
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	case key.Matches(msg, m.keys.Down):
		if !m.selected {
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		}
	case key.Matches(msg, m.keys.Enter):
		if m.timerComplete { //restart timer after completion
			m.selected = false
			m.cursor = 0
			m.minutes = 0
			m.seconds = 0
			m.isRunning = false
			m.isPaused = false
			m.timerComplete = false
			m.totalTime = 0
			m.inputting = false
			m.input = ""
			return m, nil
		}
		if !m.selected {
			if m.cursor == len(m.choices)-1 { // Custom duration selected
				m.inputting = true
				m.input = ""
				return m, nil
			} else {
				m.selected = true
				m.minutes = m.choices[m.cursor].duration
				m.totalTime = m.minutes * 60 //store total seconds
				m.isRunning = true
				m.isPaused = false
				return m, tick()
			}
		}
	case key.Matches(msg, m.keys.Pause):
		if m.selected {
			m.isPaused = !m.isPaused
			m.isRunning = !m.isPaused
			if m.isRunning {
				return m, tick()
			}
		}

	case key.Matches(msg, m.keys.Mute):
		m.muted =!m.muted
        return m, nil
	}
	return m, nil
}

//helper for Update, used for timer start and pause
func (m Model) timerUpdate() (tea.Model, tea.Cmd) {
	if m.isRunning && !m.isPaused {
		if m.seconds == 0 {
			if m.minutes == 0 {
				m.isRunning = false
				m.timerComplete = true
				helper.PlayNotification(m.muted)
				return m, nil
			}
			m.minutes--
			m.seconds = 59
		} else {
			m.seconds--
		}
		return m, tick()
	}
	return m, nil
}
