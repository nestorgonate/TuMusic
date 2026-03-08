package models

type Songs struct{
	SongsFound []DownloadedSongs
	Cursor int
	Selected map[int]struct{}
}