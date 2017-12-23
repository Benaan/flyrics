package matcher

import (
	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type ArtistTitleCleaned struct {
}

func (*ArtistTitleCleaned) Matches(fileName string, song *model.Song) bool {
	return util.ToMatchable(fileName) == util.ToMatchable(song.Artist+song.Title)
}

func (*ArtistTitleCleaned) GetCertainty() filematcher.Certainty {
	return filematcher.ARTIST_TITLE_CLEANED
}
