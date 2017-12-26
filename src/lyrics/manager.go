package lyrics

import (
	"github.com/benaan/flyrics/src/model"
)

var EmptyLyrics = &model.Lyrics{Lines: map[int]string{}}

type LyricsManager interface {
	GetLyrics(song *model.Song)
	ListProvider
	Select(*File)
}

type BestMatchProvider interface {
	GetLyrics(song *model.Song) (*model.Lyrics, error)
}

type ListProvider interface {
	GetList(song *model.Song) []*File
}

type Provider interface {
	BestMatchProvider
	ListProvider
}

type File struct {
	Song      *model.Song
	Downloads int
	Rating    string
	Source    string
	Get       func() (*model.Lyrics, error)
}

type Manager struct {
	Providers []Provider
	Output    chan model.Lyrics
}

func (manager *Manager) Select(file *File) {
	lyrics, err := file.Get()
	if err == nil {
		manager.Output <- *lyrics
	}
}

func (manager *Manager) GetLyrics(song *model.Song) {
	for _, provider := range manager.Providers {
		lyrics, err := provider.GetLyrics(song)
		if err == nil {
			manager.Output <- *lyrics
			return
		}
	}
}

func (manager *Manager) GetList(song *model.Song) (list []*File) {
	for _, provider := range manager.Providers {
		list = append(list, provider.GetList(song)...)
	}
	return
}
