package qt

import (
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/model"
	"github.com/therecipe/qt/core"
)

type Search struct {
	core.QObject

	lyricManager lyrics.LyricsManager
	files        []*lyrics.File

	_ *LyricListModel `property:"lyricList"`

	_ func() `constructor:"init"`

	_ func(artist, album, title string) `slot:"searchLyrics"`
	_ func(row int)                     `slot:"select"`
}

func (s *Search) init() {
	s.ConnectSearchLyrics(s.searchLyrics)
	s.ConnectSelect(s.selectRow)
	s.SetLyricList(NewLyricListModel(nil))
}

func (s *Search) setLyricManager(manager lyrics.LyricsManager) {
	s.lyricManager = manager
}

func (s *Search) selectRow(row int) {
	s.lyricManager.Select(s.files[row])
}

func (s *Search) searchLyrics(artist, album, title string) {
	song := &model.Song{
		Artist: artist,
		Album:  album,
		Title:  title,
	}
	s.files = s.lyricManager.GetList(song)

	s.LyricList().clear()
	for _, file := range s.files {
		lrc := NewLyric(nil)
		lrc.SetAlbum(file.Song.Album)
		lrc.SetTitle(file.Song.Title)
		lrc.SetArtist(file.Song.Artist)
		lrc.SetDownloads(file.Downloads)
		lrc.SetRating(file.Rating)
		lrc.SetSource(file.Source)
		s.LyricList().AddLyric(lrc)
	}

}
