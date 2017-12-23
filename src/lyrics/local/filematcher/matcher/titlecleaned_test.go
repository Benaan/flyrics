package matcher

import (
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func TestMatchesOnCleanedTitle(t *testing.T) {
	sut := TitleCleaned{}
	if !sut.Matches("title", &model.Song{
		Artist: "artist (unimportant)",
		Album:  "album (unimportant)",
		Title:  "title (unimportant)",
	}) {
		t.Error("Expected match")
	}
}

func TestMismatchOnCleanedTitle(t *testing.T) {
	sut := TitleCleaned{}
	if sut.Matches("title two", &model.Song{
		Artist: "artist (unimportant)",
		Album:  "album (unimportant)",
		Title:  "title (unimportant)",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestCertaintyIsTitleCleaned(t *testing.T) {
	sut := TitleCleaned{}
	if sut.GetCertainty() != filematcher.TITLE_CLEANED {
		t.Errorf("Expected certainty to be %d, received %d", sut.GetCertainty(), filematcher.TITLE_CLEANED)
	}
}
