package downloadsongs

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"tumusic/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kkdai/youtube/v2"
)

func DownloadSong(url string) error{
    client := youtube.Client{}

    video, err := client.GetVideo(url)
    if err != nil {
        return err
    }

    formats := video.Formats.WithAudioChannels()
    stream, _, err := client.GetStream(video, &formats[0])
    if err != nil {
        return err
    }
    defer stream.Close()

    return SaveSong(video.Title, stream)
}

func SaveSong(title string, stream io.ReadCloser) error {
    outputDir := "./songs"
    if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
        return err
    }

    reg := regexp.MustCompile(`[^a-zA-Z0-9\s\.\-]`)
    safeTitle := reg.ReplaceAllString(title, "")
    safeTitle = strings.ReplaceAll(strings.TrimSpace(safeTitle), " ", "_")
    safeTitle = strings.ToLower(safeTitle)

    // Guardar stream como mp4 temporal
    tmpPath := filepath.Join(outputDir, safeTitle+".tmp.mp4")
    tmpFile, err := os.Create(tmpPath)
    if err != nil {
        return err
    }

    written, err := io.Copy(tmpFile, stream)
    tmpFile.Close()
    log.Printf("Bytes escritos: %d (%.2f MB)", written, float64(written)/1024/1024)
    if err != nil || written == 0 {
        os.Remove(tmpPath)
        return fmt.Errorf("descarga fallida, bytes escritos: %d, err: %w", written, err)
    }
    defer os.Remove(tmpPath)

    // Convertir mp4 -> m4a extrayendo solo el audio
    finalPath := filepath.Join(outputDir, safeTitle+".m4a")
    cmd := exec.Command("ffmpeg",
        "-i", tmpPath,
        "-vn",          // ignorar video
        "-acodec", "copy", // copiar audio sin re-encodear (ya es aac)
        "-y",
        finalPath,
    )

    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    if err = cmd.Run(); err != nil {
        return fmt.Errorf("ffmpeg error: %w\n%s", err, stderr.String())
    }

    info, _ := os.Stat(finalPath)
    log.Printf("Guardado: %s (%.2f MB)", finalPath, float64(info.Size())/1024/1024)

    return nil
}

func DownloadCmd(url string) tea.Cmd {
    return func() tea.Msg {
        err := DownloadSong(url)
        if err != nil {
            return models.ErrorMessage{Err: err}
        }
        return models.DownloadFinishedMessage{}
    }
}