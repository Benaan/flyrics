package matcher

import (
	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

type ArtistTitle struct {
}

func (*ArtistTitle) Matches(fileName string, song *model.Song) bool {
	return fileName == (song.Artist + " - " + song.Title)
}

func (*ArtistTitle) GetCertainty() filematcher.Certainty {
	return filematcher.ARTIST_TITLE
}
