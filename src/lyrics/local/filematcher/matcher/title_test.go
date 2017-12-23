package matcher

import (
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func TestMatchesOnTitle(t *testing.T) {
	sut := Title{}
	if !sut.Matches("title", &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}) {
		t.Error("Expected match")
	}
}

func TestMismatchOnTitle(t *testing.T) {
	sut := Title{}
	if sut.Matches("title2", &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}) {
		t.Error("Didn't expected match")
	}
}

func TestCertaintyIsTitle(t *testing.T) {
	sut := Title{}
	if sut.GetCertainty() != filematcher.TITLE {
		t.Errorf("Expected certainty to be %d, received %d", sut.GetCertainty(), filematcher.TITLE)
	}
}
