package downloadingsong

import (
	"tumusic/downloadsongs"

	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedDownloadingSong) Init() tea.Cmd {
	return downloadsongs.DownloadCmd(m.Url)
}
