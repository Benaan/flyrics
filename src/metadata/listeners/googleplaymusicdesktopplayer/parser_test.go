package googleplaymusicdesktopplayer

import (
	"errors"
	"strings"
	"testing"

	"github.com/benaan/flyrics/src/model"
)

type nullReader struct {
}

func (reader *nullReader) Read(p []byte) (int, error) {
	return 0, errors.New("Couldn't read file")
}

func TestEmptyFile_ReturnsError(t *testing.T) {
	parser := Parser{}
	json := strings.NewReader("")

	if _, err := parser.Parse(json); err == nil {
		t.Error("Expected an error")
	}
}

func TestReaderError_ReturnsError(t *testing.T) {
	parser := Parser{}
	json := &nullReader{}
	if _, err := parser.Parse(json); err == nil {
		t.Error("Expected an error")
	}
}

func TestValidJson_ReturnsMetadata(t *testing.T) {
	parser := Parser{}
	json := strings.NewReader(`
	{
		"playing": true,
		"song": {
			"title": "title",
			"artist": "artist",
			"album": "album"
		},
		"time": {
			"current": 100,
			"total": 1000
		}
	}`)

	metadata, err := parser.Parse(json)
	if err != nil {
		t.Error("Didn't expect an error received:", err)
	}

	assertFieldHasValue(t, metadata.Status, model.PLAYING)
	assertFieldHasValue(t, metadata.Song.Title, "title")
	assertFieldHasValue(t, metadata.Song.Artist, "artist")
	assertFieldHasValue(t, metadata.Song.Album, "album")
	assertFieldHasValue(t, metadata.Time, 100)
}

func assertFieldHasValue(t *testing.T, fieldValue interface{}, expected interface{}) {
	t.Helper()
	if fieldValue != expected {
		t.Error(expected, "is not the same as:", fieldValue)
	}
}

func TestDetermineStatusPlaying(t *testing.T) {
	if status := determineStatus(true, 0); status != model.PLAYING {
		t.Errorf("Expected status to be playing, received: %d", status)
	}
}

func TestDetermineStatusStopped(t *testing.T) {
	if status := determineStatus(false, 0); status != model.STOPPED {
		t.Errorf("Expected status to be stopped, received: %d", status)
	}
}

func TestDetermineStatusPaused(t *testing.T) {
	if status := determineStatus(false, 1); status != model.PAUSED {
		t.Errorf("Expected status to be paused, received: %d", status)
	}
}
