package navmodel

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	listView  list.Model
	items     list.ItemDelegate // The view and style for all items in the list
	searchBar textinput.Model
}
