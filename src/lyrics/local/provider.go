package local

import (
	"errors"

	"github.com/benaan/flyrics/src/lyrics"
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

func (p *Provider) GetLyrics(song *model.Song) (*model.Lyrics, error) {
	matches, err := p.getMatchingFileList(song)
	if err != nil {
		return nil, err
	}
	bestMatch, err := filematcher.GetBestMatch(matches)
	if err != nil {
		return nil, err
	}
	return p.getLyricsFromFile(bestMatch)
}

func (p *Provider) getLyricsFromFile(path string) (*model.Lyrics, error) {
	file, err := p.FileReader.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parser.ParseLyrics(file)
}

func (p *Provider) getMatchingFileList(song *model.Song) ([]*filematcher.Match, error) {
	files, err := p.FileLister.GetAllFiles()
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("No files provided")
	}

	finder := lyricFinder{Matchers: factory.CreateMatchers()}
	return finder.find(song, files)
}

func (p *Provider) GetList(song *model.Song) (list []*lyrics.File) {
	matches, err := p.getMatchingFileList(song)
	if err != nil {
		return
	}
	for _, match := range matches {
		list = append(list, &lyrics.File{
			Song: &model.Song{
				Title: match.Path,
			},
			Source: "Local",
			Get: func() (*model.Lyrics, error) {
				return p.getLyricsFromFile(match.Path)
			},
		})
	}

	return
}
