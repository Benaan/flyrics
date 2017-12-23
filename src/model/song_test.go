package model

import "testing"

func TestSongWithTitleGetFileName(t *testing.T) {
	song := &Song{Title: "title"}
	assertFileName(song, "title.lrc", t)
}

func TestSongWithArtistTitleGetFileName(t *testing.T) {
	song := &Song{Artist: "artist", Title: "title"}
	assertFileName(song, "artist - title.lrc", t)
}

func TestSongWithArtistAlbumTitleGetFileName(t *testing.T) {
	song := &Song{Artist: "artist", Album: "album", Title: "title"}
	assertFileName(song, "artist - album - title.lrc", t)
}

func TestSongWithAlbumTitleGetFileName(t *testing.T) {
	song := &Song{Album: "album", Title: "title"}
	assertFileName(song, "title.lrc", t)
}
func assertFileName(song *Song, expected string, t *testing.T) {
	t.Helper()
	if name := song.GetFileName(); name != expected {
		t.Errorf("Expected filename to be %s, received %s", expected, name)
	}
}
