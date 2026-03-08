package playsong

import (
	"os"
	"path/filepath"
	"strings"
	"tumusic/models"

	tea "github.com/charmbracelet/bubbletea"
)

func GetSongs() ([]models.DownloadedSongs, error) {
	var songs []models.DownloadedSongs
	songsDir := "./songs"
	files, err := os.ReadDir(songsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.DownloadedSongs{}, nil
		}
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".m4a" {
				fullPath := filepath.Join(songsDir, file.Name())
				title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
				songs = append(songs, models.DownloadedSongs{
					Title:    title,
					PathSong: fullPath,
				})
			}
		}
	}
	return songs, nil
}

func GetSongsCmd() tea.Cmd {
	return func() tea.Msg {
		songs, err := GetSongs()
		if err != nil {
			return models.ErrorMessage{Err: err}
		}
		return songs
	}
}
