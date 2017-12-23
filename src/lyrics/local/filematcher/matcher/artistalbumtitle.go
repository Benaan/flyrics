package matcher

import (
	"strings"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

type ArtistAlbumTitle struct {
}

func (*ArtistAlbumTitle) Matches(fileName string, song *model.Song) bool {
	if song.Album == "" {
		return false
	}
	return strings.ToLower(fileName) == strings.ToLower(song.Artist+" - "+song.Album+" - "+song.Title)
}

func (*ArtistAlbumTitle) GetCertainty() filematcher.Certainty {
	return filematcher.ARTIST_ALBUM_TITLE
}
