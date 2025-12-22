package components

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ItemRenderer[T any] func(T) string

func ViewPaginationList[T any](
	items []T,
	render ItemRenderer[T],
	options ...Option,
) {
	cfg := defaultConfig()
	for _, opt := range options {
		opt(cfg)
	}

	m := newModel(items, render, cfg)

	program := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}

// Configurable options
type config struct {
	perPage     int
	title       string
	showHelp    bool
	activeDot   string
	inactiveDot string
}

type Option func(*config)

func WithPerPage(n int) Option {
	return func(c *config) { c.perPage = n }
}

func WithTitle(title string) Option {
	return func(c *config) { c.title = title }
}

func WithHelp(show bool) Option {
	return func(c *config) { c.showHelp = show }
}

func defaultConfig() *config {
	return &config{
		perPage:  10,
		title:    "List",
		showHelp: true,
		activeDot: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).
			Render("•"),
		inactiveDot: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).
			Render("•"),
	}
}

// Internal model
type model[T any] struct {
	items     []T
	render    ItemRenderer[T]
	paginator paginator.Model
	config    *config
}

func newModel[T any](items []T, render ItemRenderer[T], cfg *config) model[T] {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = cfg.perPage
	p.ActiveDot = cfg.activeDot
	p.InactiveDot = cfg.inactiveDot
	p.SetTotalPages(len(items))

	return model[T]{
		items:     items,
		render:    render,
		paginator: p,
		config:    cfg,
	}
}

func (m model[T]) Init() tea.Cmd {
	return nil
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model[T]) View() string {
	var b strings.Builder

	if m.config.title != "" {
		b.WriteString("\n " + m.config.title + "\n\n")
	}

	if len(m.items) == 0 {
		b.WriteString(" (No items to display)\n")
	} else {
		start, end := m.paginator.GetSliceBounds(len(m.items))
		for i := start; i < end && i < len(m.items); i++ {
			b.WriteString(" • " + m.render(m.items[i]) + "\n\n")
		}
	}

	b.WriteString("\n " + m.paginator.View() + "\n")

	if m.config.showHelp {
		b.WriteString("\n ←/→ or h/l: navigate pages • q: quit\n")
	}

	return b.String()
}
