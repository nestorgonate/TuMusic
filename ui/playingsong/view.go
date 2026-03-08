package playingsong

import (
	"fmt"
	"tumusic/models"
)

func (m EmbedPlayer) View() string {
	header := "TuMusic - Listen your favorite songs without waste resources\n"
	if m.Player.Title != "" {
		header += fmt.Sprintf("Listening: %s\n", m.Player.Title)
	}
	statusIcon := "⏹"
	switch m.Player.State {
	case models.Play:
		statusIcon = "▶"
	case models.Pause:
		statusIcon = "⏸"
	}
	header += fmt.Sprintf("%s\n\n", statusIcon)

	buttons := []string{"Pause", "Stop"}
	if m.Player.State == models.Pause {
		buttons[0] = "Resume"
	}

	for i, btn := range buttons {
		cursor := " "
		if m.Player.Cursor == i {
			cursor = ">"
			header += fmt.Sprintf("%s [%s]\n", cursor, btn)
		} else {
			header += fmt.Sprintf("%s  [%s]\n", cursor, btn)
		}
	}
	header += "\nPress ctrl+c to exit\n"
	return header
}
