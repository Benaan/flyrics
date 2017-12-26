package view

import (
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/model"
)

type LyricView interface {
	SetLyrics(lines model.Lines)
	SetSong(song *model.Song)
	SetActiveLine(row int)
	Present()
	SetLyricManager(manager lyrics.LyricsManager)
}
