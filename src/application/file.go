package application

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/benaan/flyrics/src/model"
)

type FileLister struct {
	Directory string
}

func (lister *FileLister) GetAllFiles() ([]string, error) {
	return filepath.Glob(lister.Directory + "*.lrc")
}

type FileOpener struct {
}

func (*FileOpener) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

type LyricWriter struct {
	Directory string
}

func (w *LyricWriter) Write(song *model.Song, file io.Reader) error {
	path := w.Directory + song.GetFileName()
	if _, err := os.Stat(path); err == nil {
		return errors.New("The file already exists")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}
