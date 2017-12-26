package application

import (
	"time"

	"github.com/benaan/flyrics/src/controller"
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/metadata"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/state"
	"github.com/benaan/flyrics/src/view"
)

func Run(view view.LyricView) {
	stopChannel := make(chan bool)
	lyricChannel := make(chan model.Lyrics)
	metadataChannel := make(chan model.Metadata)

	lyricManager := &lyrics.Manager{
		Output:    lyricChannel,
		Providers: CreateLyricProviders(),
	}
	view.SetLyricManager(lyricManager)

	metadataManager := &metadata.Manager{
		Delay:     1000 * time.Millisecond,
		Output:    metadataChannel,
		Stop:      stopChannel,
		Listeners: CreateMetadataListeners(),
	}

	ctrl := &controller.Controller{
		LyricManager:  lyricManager,
		State:         &state.State{},
		LyricInput:    lyricChannel,
		MetadataInput: metadataChannel,
		Stop:          stopChannel,
		View:          view,
	}

	go ctrl.Run()
	go metadataManager.Listen()
	view.Present()
}
