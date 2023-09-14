package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	writemodel "github.com/charmbracelet/bubbletea-app-template/ui/writeModel"
)

type state int

const (
	nav state = iota
	write
)

type model struct {
	state         state
	fileViewModel tea.Model
	writeModel    tea.Model
}

// View implements tea.Model.
func (m model) View() string {
	// TODO Put conditional styling breakpoints in here
	// if WindowSize.Width > xx

	if m.state == write {
		return m.writeModel.View()
	} else if m.state == nav {
		return "Nav bar state"
	}
	return "No state selected"
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		WindowSize = msg
	case tea.KeyMsg:
		if key.Matches(msg, ExitKeybinding) {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	if m.state == write {
		m.writeModel, cmd = m.writeModel.Update(msg)
	} else if m.state == nav {
		// TODO Add nav functionality
	}
	return m, cmd
}

func CreateModel() tea.Model {
	return model{
		state:      write,
		writeModel: writemodel.CreateModel(),
	}
}
