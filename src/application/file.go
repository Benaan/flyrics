package application

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/benaan/flyrics/src/config"
	"github.com/benaan/flyrics/src/model"
)

type FileLister struct {
	Config config.Reader
}

func (f *FileLister) GetAllFiles() ([]string, error) {
	return filepath.Glob(getBasePath(f.Config) + "*.lrc")
}
func getBasePath(c config.Reader) string {
	basePath := c.GetStringConfig(config.LyricDirectory)
	if !strings.HasSuffix(basePath, string(os.PathSeparator)) {
		basePath += string(os.PathSeparator)
	}
	return os.ExpandEnv(basePath)
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
	path := getBasePath(w.Config) + song.GetFileName()
	if _, err := os.Stat(path); err == nil {
		return errors.New("The file already exists")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("Empty lyric file")
	}
	return ioutil.WriteFile(path, data, 0644)
}
