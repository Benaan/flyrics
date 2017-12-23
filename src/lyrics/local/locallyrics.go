package local

import (
	"errors"
	"strings"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

type lyricFinder struct {
	Matchers []filematcher.FileMatcher
	matches  []*filematcher.Match
}

func (finder *lyricFinder) find(song *model.Song, files []string) ([]*filematcher.Match, error) {
	finder.matches = []*filematcher.Match{}

	if !hasFiles(files) {
		return nil, errors.New("No files received")
	}

	if !isValidSong(song) {
		return nil, errors.New("The song song is invalid")
	}

	for _, file := range files {
		finder.parseFile(file, song)
	}

	if !finder.hasMatches() {
		return nil, errors.New("No lyrics found")
	}

	return finder.matches, nil
}

func (finder *lyricFinder) parseFile(file string, song *model.Song) {
	fileName := stripPath(stripSuffix(file))
	for _, matcher := range finder.Matchers {
		if matcher.Matches(fileName, song) {
			finder.addMatch(file, matcher.GetCertainty())
		}
	}
}
func stripPath(path string) string {
	lastIndex := strings.LastIndex(path, "/")
	if lastIndex > -1 {
		return path[lastIndex+1:]
	}
	return path
}

func (finder *lyricFinder) addMatch(file string, certainty filematcher.Certainty) {
	finder.matches = append(finder.matches, &filematcher.Match{
		Path:      file,
		Certainty: certainty,
	})
}

func (finder *lyricFinder) hasMatches() bool {
	return len(finder.matches) > 0
}

func isValidSong(song *model.Song) bool {
	return song != nil && song.Title != "" && song.Artist != ""
}

func hasFiles(files []string) bool {
	return len(files) > 0
}

func stripSuffix(name string) string {
	return name[:len(name)-4]
}
