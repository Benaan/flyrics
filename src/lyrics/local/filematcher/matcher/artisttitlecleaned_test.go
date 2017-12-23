package matcher

import (
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func TestMatchesOnCleanedArtistTitle(t *testing.T) {
	sut := ArtistTitleCleaned{}
	if !sut.Matches("artist - title", &model.Song{
		Artist: "artist (unimportant)",
		Album:  "album (unimportant)",
		Title:  "title (unimportant)",
	}) {
		t.Error("Expected match")
	}
}

func TestMismatchOnCleanedArtistTitle(t *testing.T) {
	sut := ArtistTitleCleaned{}
	if sut.Matches("artist - title two", &model.Song{
		Artist: "artist (unimportant)",
		Album:  "album (unimportant)",
		Title:  "title (unimportant)",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestCertaintyIsArtistTitleCleaned(t *testing.T) {
	sut := ArtistTitleCleaned{}
	if sut.GetCertainty() != filematcher.ARTIST_TITLE_CLEANED {
		t.Errorf("Expected certainty to be %d, received %d", sut.GetCertainty(), filematcher.ARTIST_TITLE_CLEANED)
	}
}
