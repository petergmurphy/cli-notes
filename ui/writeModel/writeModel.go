package writemodel

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

type model struct {
	mode      mode
	textInput textinput.Model
	textArea  textarea.Model
}

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	return nil
}

// View implements tea.Model.
func (m model) View() string {
	titleTextbox := titleSelected.Render(m.textInput.View())
	if m.mode != title {
		titleTextbox = titleUnselected.Render(m.textInput.View())
	}

	bodyTextbox := m.textArea.View()
	return lipgloss.JoinVertical(0, titleTextbox, bodyTextbox)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textArea.SetWidth(msg.Width)
		m.textArea.SetHeight(msg.Height - 2)
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

func (m *model) selectTitle() {
	m.mode = title
	m.textInput.Focus()
	m.textArea.Blur()
}

func (m *model) selectBody() {
	m.mode = body
	m.textArea.Focus()
	m.textInput.Blur()
}

func CreateModel() tea.Model {
	textInput := textinput.New()
	textInput.CursorStart()
	textInput.Placeholder = "Enter title..."
	textInput.Prompt = "✍️  "
	textInput.CharLimit = 40

	textArea := textarea.New()
	textArea.FocusedStyle.Base = textArea.FocusedStyle.Base.BorderForeground(lipgloss.Color("63"))
	textArea.BlurredStyle.Base = textArea.FocusedStyle.Base.BorderForeground(lipgloss.Color("12"))

	m := model{
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
	titleSelected    = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder(), false, false, true, false).
				BorderForeground(lipgloss.Color("63")).Width(45)

	titleUnselected = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("15")).Width(45)

// titleSelected = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
//
//	BorderForeground(lipgloss.Color("63"))
)
