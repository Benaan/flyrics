package controller

import (
	"reflect"
	"testing"
	"time"

	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/state"
)

var song = &model.Song{"artist", "album", "title"}
var lyricsObj = &model.Lyrics{Lines: model.Lines{
	0:   "Line 1",
	100: "Line 2",
	200: "Line 3",
}}
var controller *Controller
var lyricInput chan model.Lyrics
var lyricManager *lyricManagerMock
var metadataInput chan model.Metadata
var stop chan bool
var stateObj *state.State
var viewObj *viewSpy

func setup() {
	stop = make(chan bool)
	lyricInput = make(chan model.Lyrics)
	metadataInput = make(chan model.Metadata)
	lyricManager = &lyricManagerMock{output: lyricInput}
	stateObj = &state.State{}
	viewObj = &viewSpy{}
	controller = &Controller{
		LyricManager:  lyricManager,
		State:         stateObj,
		LyricInput:    lyricInput,
		MetadataInput: metadataInput,
		Stop:          stop,
		View:          viewObj,
	}
	go controller.Run()
}

func TestViewIsUpdated(t *testing.T) {
	setup()
	lyricManager.lyrics = lyricsObj

	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   song,
		Time:   100,
	}
	time.Sleep(time.Millisecond)
	stop <- true

	if viewObj.lines[0] != "Line 1" {
		t.Errorf("Expected viewObj to have received line 1, received %d", stateObj.GetTime())
	}

	if viewObj.activeLine != 1 {
		t.Errorf("Expected viewObj to have activated line 1, received %d", viewObj.activeLine)
	}
}

func TestWhenSongIsNotChangedLyricsAreNotFetched(t *testing.T) {
	setup()
	lyricManager.lyrics = lyricsObj

	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   song,
		Time:   100,
	}
	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   song,
		Time:   200,
	}
	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   &model.Song{"artist", "album", "title"},
		Time:   300,
	}
	time.Sleep(time.Millisecond)
	stop <- true

	if lyricManager.count != 1 {
		t.Errorf("Expected GetLyrics to be called 1 time, but called %d times", lyricManager.count)
	}

	if viewObj.activeLine != 2 {
		t.Errorf("Expected viewObj to have activated line 2, received %d", viewObj.activeLine)

	}
}

func TestWhenSongIsChangedThenLyricsAreFetched(t *testing.T) {
	setup()
	lyricManager.lyrics = lyricsObj

	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   song,
		Time:   100,
	}
	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   &model.Song{},
		Time:   200,
	}
	time.Sleep(time.Millisecond)
	stop <- true

	if lyricManager.count != 2 {
		t.Errorf("Expected GetLyrics to be called 2 time, but called %d times", lyricManager.count)
	}
}

func TestLyricsAreChangedOnSongchange(t *testing.T) {
	setup()
	lyricManager.lyrics = lyricsObj
	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   song,
	}
	time.Sleep(time.Millisecond)
	if !reflect.DeepEqual(viewObj.lines, lyricsObj.Lines) {
		t.Errorf("Lyrics should be presented, expected %s, received %s", lyricsObj.Lines, viewObj.lines)
	}

	lyricManager.lyrics = nil
	metadataInput <- model.Metadata{
		Status: model.PLAYING,
		Song:   &model.Song{Title: "Song2"},
	}
	time.Sleep(time.Millisecond)
	if !reflect.DeepEqual(viewObj.lines, lyrics.EmptyLyrics.Lines) {
		t.Errorf("Lyrics should be cleared on song change, expected %s, received %s", lyrics.EmptyLyrics.Lines, viewObj.lines)
	}
}
