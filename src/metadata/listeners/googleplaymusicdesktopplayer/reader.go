package googleplaymusicdesktopplayer

import (
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type Reader struct {
	Path   string
	Opener util.FileOpener
}

func (listener *Reader) GetMetadata() (*model.Metadata, error) {
	file, err := listener.Opener.Open(listener.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	parser := Parser{}
	return parser.Parse(file)
}
