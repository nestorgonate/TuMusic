# 🎵 TuMusic

A lightweight terminal music player for YouTube songs, built with Go.

## Features

- Search and download songs
- Play songs directly from the terminal
- Pause and resume playback

## Dependencies

- [Go](https://golang.org/) 1.21+
- [ffmpeg](https://ffmpeg.org/) - Audio processing and conversion
- [Bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Beep](https://github.com/gopxl/beep) - Audio playback
- [kkdai/youtube](https://github.com/kkdai/youtube) - Video downloader

## Configuration

### YouTube Data API v3

TuMusic uses the YouTube Data API to search for songs. You need to create an API key:

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project
3. Enable **YouTube Data API v3**
4. Create credentials → **API Key**
5. Create a `.env` file in the root of the project:
```env
YOUTUBE_API_KEY = your_api_key_here
```

## Installation
### 1. Install ffmpeg

**Windows:**
```powershell
winget install ffmpeg
```

**Linux:**
```bash
sudo apt install ffmpeg
```

**Mac:**
```bash
brew install ffmpeg
```

### 2. Clone and run
```bash
git clone https://github.com/nestorgonate/tumusic.git
cd tumusic
go run .
```

Or build the binary:
```bash
go build -o tumusic.exe .
./tumusic.exe
```

## Usage

1. Select **Add song** to search and download a song
2. Select **Play songs** to browse your downloaded songs
3. Use arrow keys to navigate and `Enter` to select
4. Use **Pause/Resume** to control playback
5. Press `Esc` to go back, `Ctrl+C` to exit

## License

MIT
