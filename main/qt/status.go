package qt

import (
	"github.com/benaan/flyrics/src/config"
	"github.com/therecipe/qt/core"
)

type Status struct {
	core.QObject

	config config.Manager

	_ func() `constructor:"init"`

	_ int    `property:"activeLine"`
	_ string `property:"lyricDirectory"`
	_ string `property:"gpmdpPath"`
	_ *Song  `property:"currentSong"`
}

func (s *Status) init() {
	s.SetCurrentSong(NewSong(nil))
}

func (s *Status) setConfig(config config.Manager) {
	s.config = config
}

func (s *Status) listenToSettingsChanges() {
	s.SetLyricDirectory(s.config.GetStringConfig(config.LyricDirectory))
	s.ConnectLyricDirectoryChanged(func(lyricDirectory string) {
		s.config.SetConfig(config.LyricDirectory, lyricDirectory)
	})

	s.SetGpmdpPath(s.config.GetStringConfig(config.GpmdpPath))
	s.ConnectGpmdpPathChanged(func(path string) {
		s.config.SetConfig(config.GpmdpPath, path)
	})
}
