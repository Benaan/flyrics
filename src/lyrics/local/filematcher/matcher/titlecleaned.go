package matcher

import (
	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type TitleCleaned struct {
}

func (*TitleCleaned) Matches(fileName string, song *model.Song) bool {
	return util.ToMatchable(fileName) == util.ToMatchable(song.Title)
}

func (*TitleCleaned) GetCertainty() filematcher.Certainty {
	return filematcher.TITLE_CLEANED
}
