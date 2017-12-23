package viewlyrics

import (
	"testing"

	"github.com/benaan/flyrics/src/model"
)

var files = []*File{
	{
		Link:      "file1.lrc",
		Artist:    "Artist",
		Album:     "Album 1",
		Title:     "Title 1",
		Downloads: 1,
	},
	{
		Link:      "file2.lrc",
		Artist:    "Artist",
		Album:     "Album 1",
		Title:     "Title 1",
		Downloads: 2,
	},

	{
		Link:      "file3.txt",
		Artist:    "artist 2",
		Album:     "album 2",
		Title:     "title 2",
		Downloads: 2,
	},
	{
		Link:      "file4.lrc",
		Artist:    "artist",
		Title:     "title second (extra)",
		Downloads: 3,
	},
}

func TestFilterLrc(t *testing.T) {
	files := filterFiles(files)

	if count := len(files); count != 3 {
		t.Errorf("Expected length to be 3, received %d", count)
	}
}

func TestFilterAlbum(t *testing.T) {
	song := &model.Song{Album: "album"}
	files := filterSong(song, files)

	if count := len(files); count != 3 {
		t.Errorf("Expected length to be 3, received %d", count)
	}
}

func TestSortFilterAlbum(t *testing.T) {
	sortFiles(files)
	if downloads := files[0].Downloads; downloads != 3 {
		t.Errorf("Expected the top file to be downloaded 3 times, but was downloaded %d times", downloads)
	}
}

func TestEmptyFileListBestMatch(t *testing.T) {
	song := &model.Song{Album: "album"}
	if _, err := getBestMatch(song, []*File{}); err == nil {
		t.Errorf("Expected an error, recieved nil")
	}
}

func TestGetBestMatch(t *testing.T) {
	song := &model.Song{Album: "album"}
	if match, _ := getBestMatch(song, files); match != "file2.lrc" {
		t.Errorf("Expected best match to be file2.lrc, received: %s", match)
	}
}
