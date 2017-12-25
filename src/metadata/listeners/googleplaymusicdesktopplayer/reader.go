package googleplaymusicdesktopplayer

import (
	"github.com/benaan/flyrics/src/config"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type Reader struct {
	Config config.Reader
	Opener util.FileOpener
}

func (l *Reader) GetMetadata() (*model.Metadata, error) {
	file, err := l.Opener.Open(l.Config.GetStringConfig(config.GpmdpPath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	parser := Parser{}
	return parser.Parse(file)
}
