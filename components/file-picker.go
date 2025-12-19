package components

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type modelFile struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
	err          error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m modelFile) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m modelFile) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}

	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
		m.quitting = true
		return m, tea.Quit
	}

	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		m.err = errors.New(path + " is not valid")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m modelFile) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder
	s.WriteString("\n ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a CSV file (press Enter to select):")
	} else {
		s.WriteString("Selected: " + m.filepicker.Styles.Selected.Render(filepath.Base(m.selectedFile)))
	}
	s.WriteString("\n\n" + m.filepicker.View())
	s.WriteString("\n\nPress q or Ctrl+C to cancel.\n")
	return s.String()
}

func FilePicker() string {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".csv"}
	fp.ShowPermissions = false
	fp.ShowHidden = false

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	fp.CurrentDirectory = cwd

	csvDir := filepath.Join(cwd, "csv")
	if info, err := os.Stat(csvDir); err == nil && info.IsDir() {
		fp.CurrentDirectory = csvDir
	}

	m := modelFile{filepicker: fp}

	program := tea.NewProgram(m, tea.WithAltScreen())

	if tm, err := program.Run(); err != nil {
		return ""
	} else {
		finalModel := tm.(modelFile)
		if finalModel.selectedFile == "" {
			return ""
		}
		return filepath.Base(finalModel.selectedFile)
	}
}
