package getdownloadedsongs

import (
	"tumusic/playsong"

	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedGetDownloadedSongs) Init() tea.Cmd {
	return playsong.GetSongsCmd()
}