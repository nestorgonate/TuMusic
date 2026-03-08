package getsongtitle

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type EmbedGetSongTitleMenu struct {
	TextInput textinput.Model
	Exiting   bool
}

func GetSongTitleMenu() EmbedGetSongTitleMenu {
	ti := textinput.New()
	ti.Placeholder = "Song title..."
	ti.SetValue("")
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	return EmbedGetSongTitleMenu{
		TextInput: ti,
	}
}
