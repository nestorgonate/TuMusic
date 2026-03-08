package addsong

import "tumusic/models"

type EmbedingAddSongMenu struct {
	models.AddSong
}

func AddSongMenu() EmbedingAddSongMenu{
	addsongmenu := models.AddSong{
		Input: "",
		Cursor: 0,
		Songs: nil,
		Selected: make(map[int]struct{}),
	}
	embedding := EmbedingAddSongMenu{
		AddSong: addsongmenu,
	}
	return embedding
}