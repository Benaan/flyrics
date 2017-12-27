package filematcher

import (
	"errors"

	"github.com/benaan/flyrics/src/model"
)

type Certainty int

const (
	ARTIST_ALBUM_TITLE Certainty = iota
	ARTIST_ALBUM_TITLE_CLEANED
	ARTIST_TITLE
	ARTIST_TITLE_CLEANED
	TITLE
	TITLE_CLEANED
)

type Match struct {
	Path      string
	Certainty Certainty
}

type FileMatcher interface {
	Matches(fileName string, song *model.Song) bool
	GetCertainty() Certainty
}

func GetBestMatch(matches []*Match) (string, error) {
	if len(matches) == 0 {
		return "", errors.New("No input given")
	}
	var bestMatch *Match

	for _, match := range matches {
		if (bestMatch == nil || match.Certainty < bestMatch.Certainty) && match.Certainty > TITLE {
			bestMatch = match
		}
	}
	if bestMatch == nil {
		return "", errors.New("No files found above threshold")
	}
	return bestMatch.Path, nil
}
