package state

import (
	"tumusic/ui/addsong"
	"tumusic/ui/downloadingsong"
	"tumusic/ui/getsongtitle"
	"tumusic/ui/playingsong"
	"tumusic/ui/startmenu"
	getdownloadedsongs "tumusic/ui/traversesongs"
)

type sessionState int

const (
	menuView sessionState = iota
	getSongTitleView
	addSongView
	downloadingsongView
	getDownloadedSongsView
	playingsongView
)

type mainModel struct {
	state               sessionState
	startmenu           startmenu.EmbedingStartMenu
	getsongtitelmenu    getsongtitle.EmbedGetSongTitleMenu
	addsongmenu         addsong.EmbedingAddSongMenu
	downlaodingsongmenu downloadingsong.EmbedDownloadingSong
	getDownloadedSongsMenu getdownloadedsongs.EmbedGetDownloadedSongs
	playingSongMenu playingsong.EmbedPlayer
	err error
}

func InitialModel() mainModel {
	return mainModel{
		state:               menuView,
		startmenu:           startmenu.StartMenu(),
		getsongtitelmenu:    getsongtitle.GetSongTitleMenu(),
		addsongmenu:         addsong.AddSongMenu(),
		downlaodingsongmenu: downloadingsong.DownloadingSongMenu(),
		getDownloadedSongsMenu: getdownloadedsongs.GetDownloadedSongsMenu(),
		playingSongMenu: playingsong.PlayerMenu(),
	}
}
