package matcher

import (
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func TestMatchesOnCleanedArtistAlbumTitle(t *testing.T) {
	sut := ArtistAlbumTitleCleaned{}
	if !sut.Matches("artist - album - title", &model.Song{
		Artist: "artist (unimportant)",
		Album:  "album (unimportant)",
		Title:  "title (unimportant)",
	}) {
		t.Error("Expected match")
	}
}

func TestMismatchOnCleanedArtistAlbumTitle(t *testing.T) {
	sut := ArtistAlbumTitleCleaned{}
	if sut.Matches("artist - album - title two", &model.Song{
		Artist: "artist  (unimportant)",
		Album:  "album  (unimportant)",
		Title:  "title  (unimportant)",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestDoesntMatchEmptyAlbumCleaned(t *testing.T) {
	sut := ArtistAlbumTitleCleaned{}
	if sut.Matches("artist -  - title", &model.Song{
		Artist: "artist  (unimportant)",
		Album:  "",
		Title:  "title  (unimportant)",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestCertaintyIsArtistAlbumTitleCleaned(t *testing.T) {
	sut := ArtistAlbumTitleCleaned{}
	if sut.GetCertainty() != filematcher.ARTIST_ALBUM_TITLE_CLEANED {
		t.Errorf("Expected certainty to be %d, received %d", sut.GetCertainty(), filematcher.ARTIST_ALBUM_TITLE_CLEANED)
	}
}
