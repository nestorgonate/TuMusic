package state

import (
	"fmt"
	"tumusic/models"
	"tumusic/playsong"
	"tumusic/ui/addsong"
	"tumusic/ui/downloadingsong"
	"tumusic/ui/getsongtitle"
	"tumusic/ui/playingsong"
	"tumusic/ui/startmenu"
	getdownloadedsongs "tumusic/ui/traversesongs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep/speaker"
)

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case models.DownloadFinishedMessage:
		m.addsongmenu = addsong.AddSongMenu()
		m.getsongtitelmenu = getsongtitle.GetSongTitleMenu()
		m.state = menuView
		return m, nil
	case models.ErrorMessage:
		fmt.Println("Error:", msg.Err)
		m.state = menuView
		return m, nil
	case models.StopMsg:
		m.playingSongMenu = playingsong.PlayerMenu()
		m.state = getDownloadedSongsView
		return m, tea.ClearScreen
	case models.AudioStartedMsg:
		m.playingSongMenu.Player.PlayerController = msg.PlayerController
		m.playingSongMenu.Player.State = models.Play
		return m, m.playingSongMenu.WatchingSongStopCmd()
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case menuView:
			if msg.String() == "enter" {
				selected := m.startmenu.Choices[m.startmenu.Cursor]
				if selected == "Add song" {
					m.state = getSongTitleView
					return m, textinput.Blink
				}
				if selected == "Play songs" {
					m.getDownloadedSongsMenu.SongsFound, _ = playsong.GetSongs()
					m.state = getDownloadedSongsView
					return m, tea.ClearScreen
				}
			}
		case getSongTitleView:
			if msg.String() == "enter" {
				m.addsongmenu.Input = m.getsongtitelmenu.TextInput.Value()
				m.getsongtitelmenu = getsongtitle.GetSongTitleMenu()
				m.state = addSongView
				return m, tea.Batch(tea.ClearScreen, m.addsongmenu.Init())
			}
			if msg.String() == "esc" {
				m.state = menuView
				return m, tea.ClearScreen
			}
		case addSongView:
			if msg.String() == "esc" {
				m.addsongmenu = addsong.AddSongMenu()
				m.state = menuView
				return m, tea.ClearScreen
			}
			if msg.String() == "enter" {
				cursor := m.addsongmenu.Cursor
				if len(m.addsongmenu.Songs) > 0 {
					selectedSong := m.addsongmenu.Songs[cursor]
					m.downlaodingsongmenu.Url = selectedSong.Url
					m.addsongmenu = addsong.AddSongMenu()
					m.state = downloadingsongView
					return m, tea.Batch(
						tea.ClearScreen,
						m.downlaodingsongmenu.Init(),
					)
				}
			}
		case getDownloadedSongsView:
			if msg.String() == "esc" {
				m.state = menuView
				return m, tea.ClearScreen
			}
			if msg.String() == "enter" {
				cursor := m.getDownloadedSongsMenu.Cursor
				selectedSong := m.getDownloadedSongsMenu.SongsFound[cursor]
				m.playingSongMenu.Player.Title = selectedSong.Title
				m.playingSongMenu.PathSong = selectedSong.PathSong
				m.state = playingsongView
				return m, tea.Batch(
					tea.ClearScreen,
					m.playingSongMenu.Init(),
					m.playingSongMenu.PlaySongCmd(),
				)
			}
		case playingsongView:
			if msg.String() == "esc" {
				if m.playingSongMenu.Player.PlayerController != nil {
					_ = m.playingSongMenu.Player.PlayerController.Streamer.Close()
					speaker.Clear()
					m.playingSongMenu.Player.PlayerController = nil
				}
				m.playingSongMenu = playingsong.PlayerMenu()
				m.state = getDownloadedSongsView
				return m, tea.ClearScreen
			}
		}
	}

	switch m.state {
	case menuView:
		newMenu, newCmd := m.startmenu.Update(msg)
		m.startmenu = newMenu.(startmenu.EmbedingStartMenu)
		cmd = newCmd
	case getSongTitleView:
		newMenu, newCmd := m.getsongtitelmenu.Update(msg)
		m.getsongtitelmenu = newMenu.(getsongtitle.EmbedGetSongTitleMenu)
		cmd = newCmd
	case addSongView:
		newMenu, newCmd := m.addsongmenu.Update(msg)
		m.addsongmenu = newMenu.(addsong.EmbedingAddSongMenu)
		cmd = newCmd
	case downloadingsongView:
		newMenu, newCmd := m.downlaodingsongmenu.Update(msg)
		m.downlaodingsongmenu = newMenu.(downloadingsong.EmbedDownloadingSong)
		cmd = newCmd
	case getDownloadedSongsView:
		newMenu, newCmd := m.getDownloadedSongsMenu.Update(msg)
		m.getDownloadedSongsMenu = newMenu.(getdownloadedsongs.EmbedGetDownloadedSongs)
		cmd = newCmd
	case playingsongView:
		newMenu, newCmd := m.playingSongMenu.Update(msg)
		m.playingSongMenu = newMenu.(playingsong.EmbedPlayer)
		return m, newCmd
	}
	return m, cmd
}
