package playingsong

import "tumusic/models"

type EmbedPlayer struct {
	models.Player
}

func PlayerMenu() EmbedPlayer {
	player := models.Player{
		Title:    "",
		PathSong: "",
		State:    models.Play,
		Buttons: []string{"Pause", "Stop"},
	}
	embedding := EmbedPlayer{
		Player: player,
	}
	return embedding
}
