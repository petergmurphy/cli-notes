package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type mode int

const (
	title mode = iota
	body
	view
)

type writeModel struct {
	mode          mode
	textInput     textinput.Model
	textArea      textarea.Model
	width, height int
}

// Init implements tea.Model.
func (m writeModel) Init() tea.Cmd {
	return nil
}

// View implements tea.Model.
func (m writeModel) View() string {
	titleTextbox := titleStyle.Width(m.width).Render(m.textInput.View())
	// if m.mode != title {
	// 	titleTextbox = titleUnselected.Render(m.textInput.View())
	// }

	bodyTextbox := m.textArea.View()
	return lipgloss.JoinVertical(0, titleTextbox, bodyTextbox)
}

func (m writeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textArea.SetHeight(msg.Height - 2)
		m.textArea.SetWidth(m.width)
	case tea.KeyMsg:
		key := msg.String()
		if m.mode == title && (key == "down" || key == "enter") {
			m.selectBody()
			return m, nil
		}
		if m.mode == body && key == "up" && m.textArea.Line() == 0 {
			m.selectTitle()
			return m, nil
		}
	}

	if m.mode == title {
		m.textInput, cmd = m.textInput.Update(msg)
	} else if m.mode == body {
		m.textArea, cmd = m.textArea.Update(msg)
	}

	return m, cmd
}

func (m *writeModel) selectTitle() {
	m.mode = title
	m.textInput.Focus()
	m.textArea.Blur()
}

func (m *writeModel) selectBody() {
	m.mode = body
	m.textArea.Focus()
	m.textInput.Blur()
}

func CreateWriteModel() writeModel {
	textInput := textinput.New()
	textInput.CursorStart()
	textInput.Placeholder = "Enter title..."
	textInput.Prompt = "  📝  "
	textInput.CharLimit = 40

	textArea := textarea.New()
	// textArea.FocusedStyle.CursorLine = lipgloss.NewStyle().Background(lipgloss.Color("7"))
	// textArea.ShowLineNumbers = false
	// textArea.FocusedStyle.Base = textArea.FocusedStyle.Base.Border(lipgloss.Border{Top: "━"}, true, false, false, false)
	// textArea.FocusedStyle.Base = textArea.FocusedStyle.Base.BorderStyle(borderStyle}).
	// BorderForeground(lipgloss.Color("63"))
	// textArea.BlurredStyle.Base = textArea.FocusedStyle.Base.BorderForeground(lipgloss.Color("12"))
	textArea.CharLimit = 0 // Set so there is no limit to how many characters you can input

	m := writeModel{
		mode:      title,
		textInput: textInput,
		textArea:  textArea,
	}
	m.selectTitle()
	return m
}

// Styles
var (
	bottomLeftBorder = lipgloss.RoundedBorder()
	titleStyle       = lipgloss.NewStyle().
				Border(lipgloss.Border{Bottom: "━", BottomLeft: "┣", Left: "┃"}, false, false, true, true)

	titleUnselected = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("15")).Width(45)

// titleSelected = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
//
//	BorderForeground(lipgloss.Color("63"))
)
