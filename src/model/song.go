package model

import "strings"

type Song struct {
	Artist string
	Album  string
	Title  string
}

func (song *Song) GetFileName() string {
	var parts []string
	if song.Artist != "" {
		parts = append(parts, song.Artist)
		if song.Album != "" {
			parts = append(parts, song.Album)
		}
	}
	parts = append(parts, song.Title)
	return strings.Join(parts, " - ") + ".lrc"
}
