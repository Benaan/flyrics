package matcher

import (
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func TestMatchesOnArtistAlbumTitle(t *testing.T) {
	sut := ArtistAlbumTitle{}
	if !sut.Matches("artist - album - title", &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}) {
		t.Error("Expected match")
	}
}

func TestMismatchOnArtistAlbumTitle(t *testing.T) {
	sut := ArtistAlbumTitle{}
	if sut.Matches("artist - album - title2", &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestDoesntMatchOnEmptyAlbum(t *testing.T) {
	sut := ArtistAlbumTitle{}
	if sut.Matches("artist -  - title", &model.Song{
		Artist: "artist",
		Album:  "",
		Title:  "title",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestDoesntMatchEmptyAlbum(t *testing.T) {
	sut := ArtistAlbumTitleCleaned{}
	if sut.Matches("artist -  - title", &model.Song{
		Artist: "artist",
		Album:  "",
		Title:  "title",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestCertaintyIsArtistAlbumTitle(t *testing.T) {
	sut := ArtistAlbumTitle{}
	if sut.GetCertainty() != filematcher.ARTIST_ALBUM_TITLE {
		t.Errorf("Expected certainty to be %d, received %d", sut.GetCertainty(), filematcher.ARTIST_ALBUM_TITLE)
	}
}
