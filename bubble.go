package main

/* they are shit, dont use again
"github.com/charmbracelet/bubbles/cursor"
"github.com/charmbracelet/bubbles/textinput"
tea "github.com/charmbracelet/bubbletea"
*/

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/cursor"
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var (
// 	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))
// 	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
// 	noStyle             = lipgloss.NewStyle()
// 	helpStyle           = blurredStyle.Copy()
// 	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

// 	focusedButton = focusedStyle.Copy().Render("[submit]")
// 	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("submit"))
// )

// var initialModel model

// type model struct {
// 	focusIndex int
// 	cursorMode cursor.Mode
// 	inputs     []textinput.Model
// }

// func (m model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.Type {
// 		case tea.KeyCtrlC:
// 			return m, tea.Quit

// 		case tea.KeyUp, tea.KeyDown, tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter:
// 			s := msg.String()

// 			if s == "enter" && m.focusIndex == len(m.inputs) {
// 				return m, tea.Quit
// 			}

// 			if s == "up" || s == "shift+tab" {
// 				m.focusIndex--
// 			} else {
// 				m.focusIndex++
// 			}

// 			if m.focusIndex > len(m.inputs) {
// 				m.focusIndex = 0
// 			} else if m.focusIndex < 0 {
// 				m.focusIndex = len(m.inputs)
// 			}

// 			cmds := make([]tea.Cmd, len(m.inputs))
// 			for i := 0; i <= len(m.inputs)-1; i++ {
// 				if i == m.focusIndex {
// 					cmds[i] = m.inputs[i].Focus()
// 					m.inputs[i].PromptStyle = focusedStyle
// 					m.inputs[i].TextStyle = focusedStyle
// 				} else {
// 					m.inputs[i].Blur()
// 					m.inputs[i].PromptStyle = noStyle
// 					m.inputs[i].TextStyle = noStyle
// 				}
// 			}
// 			return m, tea.Batch(cmds...)

// 		case tea.KeyCtrlR:
// 			m.cursorMode++
// 			if m.cursorMode > cursor.CursorHide {
// 				m.cursorMode = cursor.CursorBlink
// 			}
// 			cmds := make([]tea.Cmd, len(m.inputs))
// 			for i := range cmds {
// 				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
// 			}
// 			return m, tea.Batch(cmds...)

// 		}
// 	}
// 	cmd := m.updateInputs(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	var b strings.Builder

// 	for i := range m.inputs {
// 		b.WriteString(m.inputs[i].View())
// 		if i < len(m.inputs)-1 {
// 			b.WriteRune('\n')
// 		}
// 	}

// 	button := &blurredButton
// 	if m.focusIndex == len(m.inputs) {
// 		button = &focusedButton
// 	}
// 	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

// 	b.WriteString(helpStyle.Render("cursor mode is "))
// 	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
// 	b.WriteString(helpStyle.Render(" ctrl+r to change style"))

// 	return b.String()
// }

// func (m model) updateInputs(msg tea.Msg) tea.Cmd {
// 	cmds := make([]tea.Cmd, len(m.inputs))

// 	for i := range m.inputs {
// 		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
// 	}

// 	return tea.Batch(cmds...)
// }

// func init() {
// 	initialModel = model{
// 		inputs: make([]textinput.Model, 6),
// 	}

// 	var t textinput.Model
// 	for i := range initialModel.inputs {
// 		t = textinput.New()
// 		t.CharLimit = 32

// 		switch i {
// 		case 0:
// 			t.Placeholder = "domain"
// 			t.Focus()
// 			t.PromptStyle = focusedStyle
// 			t.TextStyle = focusedStyle
// 		case 1:
// 			t.Placeholder = "has MX"
// 		case 2:
// 			t.Placeholder = "has SPF"
// 			t.EchoMode = textinput.EchoPassword
// 			t.EchoCharacter = '*'
// 		case 3:
// 			t.Placeholder = "spr Record"
// 		case 4:
// 			t.Placeholder = "has DMARC"
// 		case 5:
// 			t.Placeholder = "dmarc Record"
// 		}

// 		initialModel.inputs[i] = t
// 	}
// }

// func init() {
// 	initialModel = model{
// 		inputs: make([]textinput.Model, 6),
// 	}

// 	var t textinput.Model
// 	for i := range initialModel.inputs {
// 		t = textinput.New()
// 		t.CharLimit = 32

// 		switch i {
// 		case 0:
// 			t.Placeholder = "domain"
// 			t.Focus()
// 			t.PromptStyle = focusedStyle
// 			t.TextStyle = focusedStyle
// 			t.Validate = func(s string) error {
// 				_, err := mail.ParseAddress(s)
// 				return err
// 			}
// 		case 1:
// 			t.Placeholder = "has MX"
// 		case 2:
// 			t.Placeholder = "has SPF"
// 			t.EchoMode = textinput.EchoPassword
// 			t.EchoCharacter = '*'
// 		case 3:
// 			t.Placeholder = "spf Record"
// 		case 4:
// 			t.Placeholder = "has DMARC"
// 		case 5:
// 			t.Placeholder = "dmarc Record"
// 		}

// 		initialModel.inputs[i] = t
// 	}
// }
