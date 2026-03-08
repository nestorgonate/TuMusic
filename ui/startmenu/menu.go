package startmenu

import "tumusic/models"

type EmbedingStartMenu struct {
	models.StartMenuState
}

func StartMenu() EmbedingStartMenu{
	startmenu := models.StartMenuState{
		Choices: []string{"Add song", "Play songs"},
		Cursor: 0,
		Selected: make(map[int]struct{}),
		Exiting: false,
	}
	embedding := EmbedingStartMenu{
		StartMenuState: startmenu,
	}
	return embedding
}