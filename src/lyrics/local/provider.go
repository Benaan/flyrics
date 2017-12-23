package local

import (
	"errors"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/lyrics/local/filematcher/factory"
	"github.com/benaan/flyrics/src/lyrics/parser"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type FileLister interface {
	GetAllFiles() ([]string, error)
}

type Provider struct {
	FileReader util.FileOpener
	FileLister FileLister
}

func (provider *Provider) GetLyrics(song *model.Song) (*model.Lyrics, error) {
	files, err := provider.FileLister.GetAllFiles()
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("No files provided")
	}

	finder := lyricFinder{Matchers: factory.CreateMatchers()}
	matches, err := finder.find(song, files)
	if err != nil {
		return nil, err
	}
	bestMatch, err := filematcher.GetBestMatch(matches)

	file, err := provider.FileReader.Open(bestMatch)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parser.ParseLyrics(file)
}
