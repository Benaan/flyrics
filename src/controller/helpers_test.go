package controller

import (
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/model"
)

type viewSpy struct {
	activeLine int
	lines      model.Lines
}

func (spy *viewSpy) SetSong(song *model.Song) {

}

func (spy *viewSpy) SetLyricManager(manager lyrics.LyricsManager) {
	panic("implement me")
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

func (manager *lyricManagerMock) Select(*lyrics.File) {
	panic("implement me")
}

func (manager *lyricManagerMock) GetList(song *model.Song) (list []*lyrics.File) {
	panic("implement me")
}

func (manager *lyricManagerMock) GetLyrics(song *model.Song) {
	manager.count++
	if manager.lyrics != nil {
		manager.output <- *manager.lyrics
	}
}
