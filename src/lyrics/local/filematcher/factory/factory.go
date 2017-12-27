package factory

import (
	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/lyrics/local/filematcher/matcher"
)

func CreateMatchers() []filematcher.FileMatcher {
	return []filematcher.FileMatcher{
		&matcher.ArtistAlbumTitle{},
		&matcher.ArtistTitle{},
		&matcher.Title{},
		&matcher.ArtistAlbumTitleCleaned{},
		&matcher.ArtistTitleCleaned{},
		&matcher.TitleCleaned{},
	}
}
