package addsong

import (
	"tumusic/downloadsongs"

	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedingAddSongMenu) Init() tea.Cmd {
	return downloadsongs.SearchSongCmd(m.Input)
}