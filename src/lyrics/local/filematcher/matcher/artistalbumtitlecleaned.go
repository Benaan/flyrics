package matcher

import (
	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type ArtistAlbumTitleCleaned struct {
}

func (*ArtistAlbumTitleCleaned) Matches(fileName string, song *model.Song) bool {
	if song.Album == "" {
		return false
	}
	return util.ToMatchable(fileName) == util.ToMatchable(song.Artist+song.Album+song.Title)
}

func (*ArtistAlbumTitleCleaned) GetCertainty() filematcher.Certainty {
	return filematcher.ARTIST_ALBUM_TITLE_CLEANED
}
