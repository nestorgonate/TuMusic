package downloadsongs

import (
	"context"
	"fmt"
	"log"
	"os"
	"tumusic/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func SearchSong(input string) ([]models.SongResult, error){
	ctx := context.Background()
	godotenv.Load()
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil{
		return nil, err
	}
	query := service.Search.List([]string{"snippet", "id"}).Q(input).MaxResults(2).Type("video")
	response, err := query.Do()
	if err != nil{
		return nil, err
	}
	var results []models.SongResult
	for _, item := range response.Items{
		video := models.SongResult{
			Title: item.Snippet.Title,
			Channel: item.Snippet.ChannelTitle,
			Url: fmt.Sprintf("https://youtube.com/watch?v=%s", item.Id.VideoId),
		}
		results =append(results, video)
	}
	return results, nil
}

func SearchSongCmd(input string) tea.Cmd{
	log.Print("Searching song")
	return func() tea.Msg {
		songs, err := SearchSong(input)
		if err != nil{
			return err
		}
		return models.SearchResultMsg(songs)
	}
}