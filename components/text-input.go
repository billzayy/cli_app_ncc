package components

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyleTextInput  = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type modelTextInput struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func initialModelTextInput(placeholders []string) modelTextInput {
	m := modelTextInput{
		inputs: make([]textinput.Model, len(placeholders)),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32
		t.Width = 20

		// Safely get placeholder (fallback if index out of range or empty)
		if i < len(placeholders) && placeholders[i] != "" {
			t.Placeholder = placeholders[i]
		} else {
			t.Placeholder = fmt.Sprintf("Input %d", i+1)
		}

		// Special handling only for the first input: focus it
		if i == 0 {
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		// Optional: customize specific fields based on placeholder name or index
		// Example: detect password field by placeholder text
		if strings.Contains(strings.ToLower(t.Placeholder), "password") {
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		// Example: longer limit for email-like fields
		if strings.Contains(strings.ToLower(t.Placeholder), "email") {
			t.CharLimit = 64
		}

		m.inputs[i] = t
	}

	return m
}

func (m modelTextInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m modelTextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *modelTextInput) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m modelTextInput) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyleTextInput.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyleTextInput.Render(" (ctrl+r to change style)"))

	return b.String()
}

func TextInput[T any](titles []string) T {
	finalModel, err := tea.NewProgram(initialModelTextInput(titles)).Run()

	if err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}

	m := finalModel.(modelTextInput)

	// Create a zero value of the generic type T
	var result T

	// Use reflection to fill the struct fields based on titles/order
	// This assumes fields are in the same order as inputs
	v := reflect.ValueOf(&result).Elem()

	numFields := v.NumField()
	if len(m.inputs) != numFields {
		panic(fmt.Sprintf("expected %d inputs, got %d", numFields, len(m.inputs)))
	}

	for i := 0; i < numFields; i++ {
		inputValue := m.inputs[i].Value()
		field := v.Field(i)
		if field.CanSet() {
			field.SetString(inputValue)
		}
	}

	return result
}
