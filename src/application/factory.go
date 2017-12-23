package application

import (
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/lyrics/local"
	"github.com/benaan/flyrics/src/lyrics/viewlyrics"
	"github.com/benaan/flyrics/src/metadata"
	"github.com/benaan/flyrics/src/metadata/listeners/googleplaymusicdesktopplayer"
)

func CreateLyricProviders() []lyrics.LyricProvider {
	lyricDirectory := "/home/gertjan/Music/lyrics/"
	return []lyrics.LyricProvider{
		&local.Provider{
			FileReader: &FileOpener{},
			FileLister: &FileLister{lyricDirectory},
		},
		&viewlyrics.ViewLyrics{Writer: &LyricWriter{lyricDirectory}},
	}
}

func CreateMetadataListeners() []metadata.Listener {
	return []metadata.Listener{
		&googleplaymusicdesktopplayer.Reader{
			Path:   "/home/gertjan/.config/Google Play Music Desktop Player/json_store/playback.json",
			Opener: &FileOpener{},
		},
	}
}
