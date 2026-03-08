package getdownloadedsongs

import "tumusic/models"

type EmbedGetDownloadedSongs struct {
	models.Songs
}

func GetDownloadedSongsMenu() EmbedGetDownloadedSongs {
	songs := models.Songs{
		SongsFound: make([]models.DownloadedSongs, 0),
	}
	embedding := EmbedGetDownloadedSongs{
		Songs: songs,
	}
	return embedding
}
