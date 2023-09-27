package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
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

var SwitchModeKeyBinding = key.NewBinding(
	key.WithKeys("esc"),
	key.WithHelp("escape", "Switch Mode"),
)

var EnterKeybinding = key.NewBinding(
	key.WithKeys("enter"),
)

func NewItemStyles() (s list.DefaultItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2)

	s.NormalDesc = s.NormalTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#fcb905", Dark: "#fcb905"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#e8b425", Dark: "#e8b425"}).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#c4981f", Dark: "#c4981f"})

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
		Padding(0, 0, 0, 2)

	s.DimmedDesc = s.DimmedTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
