package util

import (
	"io"

	"github.com/benaan/flyrics/src/model"
)

type FileOpener interface {
	Open(path string) (io.ReadCloser, error)
}

type LyricWriter interface {
	Write(song *model.Song, file io.Reader) error
}
