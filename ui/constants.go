package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	singleViewBreakpoint = 100
	sidebarWidth         = 40
)

var (
	// WindowSize store the size of the terminal window
	WindowSize tea.WindowSizeMsg
)

// HelpStyle styling for help context menu
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

var ExitKeybinding = key.NewBinding(
	key.WithKeys("ctrl+z"),
	key.WithHelp("ctrl + z", "exit"),
)

var EnterKeybinding = key.NewBinding(
	key.WithKeys("enter"),
)
