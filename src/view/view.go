package view

import (
	"github.com/benaan/flyrics/src/model"
)

type LyricView interface {
	SetLyrics(lines model.Lines)
	SetActiveLine(row int)
	Present()
}
