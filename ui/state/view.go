package state

func (m mainModel) View() string {
	switch m.state {
	case menuView:
		return m.startmenu.View()
	case addSongView:
		return m.addsongmenu.View()
	case getSongTitleView:
		return m.getsongtitelmenu.View()
	case downloadingsongView:
		return m.downlaodingsongmenu.View()
	case getDownloadedSongsView:
		return m.getDownloadedSongsMenu.View()
	case playingsongView:
		return m.playingSongMenu.View()
	default:
		return "Loading..."
	}
}
