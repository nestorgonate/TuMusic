package getsongtitle

func (m EmbedGetSongTitleMenu) View() string {
	header := "Type the title of your favorite song:\n"
    header += m.TextInput.View() + "\n"
    header += "Press ctrl+c to exit"
    return header
}