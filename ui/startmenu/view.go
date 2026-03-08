package startmenu

import (
	"fmt"
)

func (m EmbedingStartMenu) View() string {
	header := "TuMusic - Listen your favorite songs without waste resources\n"
	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		header += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	header += "\nPress ctrl+c to exit\n"
	return header
}