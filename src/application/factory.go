package application

import (
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/lyrics/local"
	"github.com/benaan/flyrics/src/lyrics/viewlyrics"
	"github.com/benaan/flyrics/src/metadata"
	"github.com/benaan/flyrics/src/metadata/listeners/googleplaymusicdesktopplayer"
)

func CreateLyricProviders() []lyrics.LyricProvider {
	return []lyrics.LyricProvider{
		&local.Provider{
			FileReader: &FileOpener{},
			FileLister: &FileLister{Config},
		},
		&viewlyrics.ViewLyrics{Writer: &LyricWriter{Config}},
	}
}

func CreateMetadataListeners() []metadata.Listener {
	return []metadata.Listener{
		&googleplaymusicdesktopplayer.Reader{
			Config: Config,
			Opener: &FileOpener{},
		},
	}
}
