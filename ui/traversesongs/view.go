package getdownloadedsongs

import "fmt"

func (m EmbedGetDownloadedSongs) View() string {
	header := "TuMusic - Listen your favorite songs without waste resources\n"
	for i, song := range m.SongsFound {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		header += fmt.Sprintf("%s %s\n", cursor, song.Title)
	}
	header += "\nPress ctrl+c to exit\n"
	return header
}
