package downloadingsong

import "fmt"

func (m EmbedDownloadingSong) View() string {
	return fmt.Sprintf(
		"\n  🚀 %s\n\n  %s\n\n",
        "TuMusic - Downloader",
        "Downloading your song...",
	)
}
