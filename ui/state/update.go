package state

import (
	"fmt"
	"log"
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
		// Imprime el error para saber qué falló
		fmt.Println("Error:", msg.Err)
		m.state = menuView
		return m, nil
	case models.StopMsg:
		log.Print("StopMsg recibido")
		m.state = menuView
		return m, nil
	case models.AudioStartedMsg:
		newMenu, newCmd := m.playingSongMenu.Update(msg)
		m.playingSongMenu = newMenu.(playingsong.EmbedPlayer)
		return m, newCmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case menuView:
			if msg.String() == "enter" {
				selected := m.startmenu.Choices[m.startmenu.Cursor]
				if selected == "Add song" {
					m.state = getSongTitleView // Cambia el estado
					return m, textinput.Blink
				}
				if selected == "Play songs" {
					m.getDownloadedSongsMenu.SongsFound, _ = playsong.GetSongs()
					m.state = getDownloadedSongsView
					return m, nil
				}
			}
		case getSongTitleView:
			if msg.String() == "enter" {
				m.addsongmenu.Input = m.getsongtitelmenu.TextInput.Value()
				m.state = addSongView
				return m, m.addsongmenu.Init()
			}
			if msg.String() == "esc" {
				m.state = menuView
				return m, nil
			}
		case addSongView:
			if msg.String() == "esc" {
				m.state = menuView
				return m, nil
			}
			if msg.String() == "enter" {
				cursor := m.addsongmenu.Cursor
				if len(m.addsongmenu.Songs) > 0 {
					selectedSong := m.addsongmenu.Songs[cursor]
					m.downlaodingsongmenu.Url = selectedSong.Url
					m.state = downloadingsongView
					return m, m.downlaodingsongmenu.Init()
				}
			}
		case getDownloadedSongsView:
			if msg.String() == "esc" {
				m.state = menuView
				return m, nil
			}
			if msg.String() == "enter" {
				cursor := m.getDownloadedSongsMenu.Cursor
				selectedSong := m.getDownloadedSongsMenu.SongsFound[cursor]
				m.playingSongMenu.Player.Title = selectedSong.Title
				m.playingSongMenu.PathSong = selectedSong.PathSong
				m.state = playingsongView
				return m, tea.Batch(
					m.playingSongMenu.Init(),
					m.playingSongMenu.PlaySongCmd(),
				)
			}
		case playingsongView:
			if msg.String() == "esc" {
				if m.playingSongMenu.Player.PlayerController != nil {
					speaker.Lock()
					speaker.Clear()
					speaker.Unlock()
					_ = m.playingSongMenu.Player.PlayerController.Streamer.Close()
					m.playingSongMenu.Player.PlayerController = nil
				}
				m.state = menuView
				return m, nil
			}
			newMenu, newCmd := m.playingSongMenu.Update(msg)
			m.playingSongMenu = newMenu.(playingsong.EmbedPlayer)
			return m, newCmd
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
	}
	return m, cmd
}