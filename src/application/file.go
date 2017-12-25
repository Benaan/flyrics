package application

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/benaan/flyrics/src/config"
	"github.com/benaan/flyrics/src/model"
)

type FileLister struct {
	Config config.Reader
}

func (lister *FileLister) GetAllFiles() ([]string, error) {
	return filepath.Glob(lister.Config.GetStringConfig(config.LyricDirectory) + "*.lrc")
}

type FileOpener struct {
}

func (*FileOpener) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

type LyricWriter struct {
	Config config.Reader
}

func (w *LyricWriter) Write(song *model.Song, file io.Reader) error {
	path := w.Config.GetStringConfig(config.LyricDirectory) + song.GetFileName()
	if _, err := os.Stat(path); err == nil {
		return errors.New("The file already exists")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}
