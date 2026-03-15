package getsongtitle

func (m EmbedGetSongTitleMenu) View() string {
	header := "Type the title of your favorite song:\n"
    header += m.TextInput.View() + "\n"
    header += "Press ctrl+c to exit or esc to return"
    return header
}