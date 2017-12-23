package matcher

import (
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func TestMatchesOnArtistTitle(t *testing.T) {
	sut := ArtistTitle{}
	if !sut.Matches("artist - title", &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}) {
		t.Error("Expected match")
	}
}

func TestMismatchOnArtistTitle(t *testing.T) {
	sut := ArtistTitle{}
	if sut.Matches("artist - title2", &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestCertaintyIsArtistTitle(t *testing.T) {
	sut := ArtistTitle{}
	if sut.GetCertainty() != filematcher.ARTIST_TITLE {
		t.Errorf("Expected certainty to be %d, received %d", sut.GetCertainty(), filematcher.ARTIST_TITLE)
	}
}
