package matcher

import (
	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

type Title struct {
}

func (*Title) Matches(fileName string, song *model.Song) bool {
	return fileName == song.Title
}

func (*Title) GetCertainty() filematcher.Certainty {
	return filematcher.TITLE
}
