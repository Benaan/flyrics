package controller

import (
	"github.com/benaan/flyrics/src/model"
)

type viewSpy struct {
	activeLine int
	lines      model.Lines
}

func (spy *viewSpy) Present() {
}

func (spy *viewSpy) SetLyrics(lines model.Lines) {
	spy.lines = lines
}

func (spy *viewSpy) SetActiveLine(row int) {
	spy.activeLine = row
}

type lyricManagerMock struct {
	output chan model.Lyrics
	count  int
	lyrics *model.Lyrics
}

func (manager *lyricManagerMock) GetLyrics(song *model.Song, output chan model.Lyrics) {
	manager.count++
	if manager.lyrics != nil {
		output <- *manager.lyrics
	}
}
