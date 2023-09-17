package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	nav state = iota
	write
)

type model struct {
	state         state
	fileViewModel navModel
	writeModel    writeModel
	windowWidth   int
}

// View implements tea.Model.
func (m model) View() string {
	// TODO Put conditional styling breakpoints in here
	// if WindowSize.Width > xx

	if m.state == write {
		if m.fileViewModel.width == 0 {
			return m.writeModel.View()
		}
		return lipgloss.JoinHorizontal(0, m.fileViewModel.View(), m.writeModel.View())
	} else if m.state == nav {
		if m.writeModel.width == 0 {
			return m.fileViewModel.View()
		}
		return lipgloss.JoinHorizontal(0, m.fileViewModel.View(), m.writeModel.View())
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
		// TODO: UPDATE WIDTH AND HEIGHT OF EACH COMPONENT HERE
		m.windowWidth = msg.Width
		if msg.Width > singleViewBreakpoint {
			m.writeModel.width = msg.Width - sidebarWidth
			m.fileViewModel.width = sidebarWidth
		} else {
			if m.state == nav {
				m.writeModel.width = 0
				m.fileViewModel.width = msg.Width
			} else if m.state == write {
				m.writeModel.width = msg.Width
				m.fileViewModel.width = 0
			}
		}
	case tea.KeyMsg:
		if key.Matches(msg, ExitKeybinding) {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	if m.state == write {
		var newWriteModel tea.Model
		newWriteModel, cmd = m.writeModel.Update(msg)
		m.writeModel = newWriteModel.(writeModel)
		if m.windowWidth > singleViewBreakpoint {
			var newNavModel tea.Model
			newNavModel, cmd = m.fileViewModel.Update(msg)
			m.fileViewModel = newNavModel.(navModel)
		}
	} else if m.state == nav {
		// TODO Add nav functionality
	}
	return m, cmd
}

func CreateUI() tea.Model {
	return model{
		state:         write,
		writeModel:    CreateWriteModel(),
		fileViewModel: CreateNavModel(),
	}
}
