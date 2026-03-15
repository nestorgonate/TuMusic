package downloadingsong

import (
	"tumusic/downloadsongs"

	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedDownloadingSong) Init() tea.Cmd {
	return tea.Batch(downloadsongs.DownloadCmd(m.Url))
}
