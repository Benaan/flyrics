package lyrics

import (
	"github.com/benaan/flyrics/src/model"
)

var EmptyLyrics = model.Lyrics{Lines: map[int]string{}}

type LyricProvider interface {
	GetLyrics(song *model.Song) (*model.Lyrics, error)
}

type Manager struct {
	Providers []LyricProvider
}

func (manager *Manager) GetLyrics(song *model.Song, output chan model.Lyrics) {
	for _, provider := range manager.Providers {
		lyrics, err := provider.GetLyrics(song)
		if err == nil {
			output <- *lyrics
			return
		}
	}
}
