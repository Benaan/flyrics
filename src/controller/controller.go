package controller

import (
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/state"
	"github.com/benaan/flyrics/src/view"
)

type lyricsManager interface {
	GetLyrics(song *model.Song, output chan model.Lyrics)
}

type Controller struct {
	LyricManager  lyricsManager
	View          view.LyricView
	State         *state.State
	LyricInput    chan model.Lyrics
	MetadataInput chan model.Metadata
	Stop          chan bool
}

func (controller *Controller) Run() {
	for {
		select {
		case metadata := <-controller.MetadataInput:
			controller.handleMetadataChange(metadata)
		case lyrics := <-controller.LyricInput:
			controller.handleLyricChange(lyrics)
		case <-controller.Stop:
			return
		}
	}
}

func (controller *Controller) handleLyricChange(lyrics model.Lyrics) {
	controller.State.SetLyrics(&lyrics)
	controller.View.SetLyrics(lyrics.Lines)
	controller.View.SetActiveLine(controller.State.GetActiveLine())
}

func (controller *Controller) handleMetadataChange(metadata model.Metadata) {
	if controller.State.GetSong() == nil || *controller.State.GetSong() != *metadata.Song {
		controller.State.SetSong(metadata.Song)
		go controller.LyricManager.GetLyrics(metadata.Song, controller.LyricInput)
	}
	controller.State.SetStatus(metadata.Status)
	controller.State.SetTime(metadata.Time)
	controller.View.SetActiveLine(controller.State.GetActiveLine())
}
