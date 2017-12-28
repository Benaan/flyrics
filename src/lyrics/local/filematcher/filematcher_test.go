package filematcher

import (
	"testing"
)

func TestReturnsErrorWhenNoMatches(t *testing.T) {
	_, err := GetBestMatch([]*Match{})
	if err == nil {
		t.Error("Expected error")
	}
}

func TestReturnsSinglePath(t *testing.T) {

	path, err := GetBestMatch([]*Match{
		{"path", ARTIST_ALBUM_TITLE},
	})

	if err != nil {
		t.Errorf("Didnt expect an error, received: %s", err)
	}
	if path != "path" {
		t.Errorf("Expected path to be \"path\", received: %s", path)
	}
}

func TestReturnsBestMatch(t *testing.T) {
	path, err := GetBestMatch([]*Match{
		{"path", TITLE},
		{"path2", ARTIST_ALBUM_TITLE},
		{"path3", ARTIST_TITLE_CLEANED},
	})
	if err != nil {
		t.Errorf("Didnt expect an error, received: %s", err)
	}

	if path != "path2" {
		t.Errorf("Expected path2 te be returned, received %s", path)
	}
}
