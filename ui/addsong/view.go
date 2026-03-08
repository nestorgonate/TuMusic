package addsong

import (
	"fmt"
)

func (m EmbedingAddSongMenu) View() string {
	header := "TuMusic - Listen your favorite songs without waste resources\n"
	for i, song := range m.Songs {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		header += fmt.Sprintf("%s Title: %s - Uploaded by: %s\n", cursor, song.Title, song.Channel)
	}
	header += "\nPress ctrl+c to exit\n"
	return header
}