package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type navModel struct {
	listView      list.Model
	items         list.ItemDelegate // The view and style for all items in the list
	searchBar     textinput.Model
	isHidden      bool
	width, height int
}

func (m navModel) Init() tea.Cmd {
	return nil
}

func (m navModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.listView.SetHeight(msg.Height)
		m.listView.SetWidth(m.width)
		// case tea.KeyMsg:
		// 	key := msg.String()
		// 	if m.mode == title && (key == "down" || key == "enter") {
		// 		m.selectBody()
		// 		return m, nil
		// 	}
		// 	if m.mode == body && key == "up" && m.textArea.Line() == 0 {
		// 		m.selectTitle()
		// 		return m, nil
		// 	}
	}

	return m, cmd
}

func (m navModel) View() string {
	return m.listView.View()
}

func CreateNavModel() navModel {
	tempItems := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "Bitter melon", desc: "It cools you down"},
		item{title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{title: "Eight hours of sleep", desc: "I had this once"},
		item{title: "Cats", desc: "Usually"},
		item{title: "Plantasia, the album", desc: "My plants love it too"},
		item{title: "Pour over coffee", desc: "It takes forever to make though"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
		item{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		item{title: "Linux", desc: "Pretty much the best OS"},
		item{title: "Business school", desc: "Just kidding"},
		item{title: "Pottery", desc: "Wet clay is a great feeling"},
		item{title: "Shampoo", desc: "Nothing like clean hair"},
		item{title: "Table tennis", desc: "It’s surprisingly exhausting"},
		item{title: "Milk crates", desc: "Great for packing in your extra stuff"},
		item{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		item{title: "Stickers", desc: "The thicker the vinyl the better"},
		item{title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		item{title: "Warm light", desc: "Like around 2700 Kelvin"},
		item{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		item{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		item{title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	itemDelegate := list.NewDefaultDelegate()
	newList := list.New(tempItems, itemDelegate, sidebarWidth, 20)
	newList.Title = "Notes"

	model := navModel{
		listView: newList,
	}
	return model
}
