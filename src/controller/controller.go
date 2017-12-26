package controller

import (
	"time"

	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/state"
	"github.com/benaan/flyrics/src/view"
)

type Controller struct {
	LyricManager  lyrics.LyricsManager
	View          view.LyricView
	State         *state.State
	LyricInput    chan model.Lyrics
	MetadataInput chan model.Metadata
	Stop          chan bool
}

func (controller *Controller) Run() {
	time.Sleep(100 * time.Millisecond) // todo add real fix - the first song is updated before the view is ready
	for {
		select {
		case metadata := <-controller.MetadataInput:
			controller.handleMetadataChange(&metadata)
		case lrc := <-controller.LyricInput:
			controller.handleLyricChange(&lrc)
		case <-controller.Stop:
			return
		}
	}
}

func (controller *Controller) handleLyricChange(lyrics *model.Lyrics) {
	controller.State.SetLyrics(lyrics)
	controller.View.SetLyrics(lyrics.Lines)
	controller.View.SetActiveLine(controller.State.GetActiveLine())
}

func (controller *Controller) handleMetadataChange(metadata *model.Metadata) {
	if controller.State.GetSong() == nil || *controller.State.GetSong() != *metadata.Song {
		controller.View.SetSong(metadata.Song)
		controller.State.SetSong(metadata.Song)
		controller.handleLyricChange(lyrics.EmptyLyrics)
		go controller.LyricManager.GetLyrics(metadata.Song)
	}
	controller.State.SetStatus(metadata.Status)
	controller.State.SetTime(metadata.Time)
	controller.View.SetActiveLine(controller.State.GetActiveLine())
}
