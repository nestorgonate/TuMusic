package models

type SongResult struct{
	Title string
	Channel string
	Url string
}

type SearchResultMsg []SongResult