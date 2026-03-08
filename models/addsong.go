package models

type AddSong struct {
	Input   string
	Cursor  int
	Songs   []SongResult
	Selected map[int]struct{}
}

type ErrorMessage struct{
	Err error
}
type DownloadFinishedMessage struct{}
