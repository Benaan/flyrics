package local

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/benaan/flyrics/src/model"
)

type fileListerMock struct {
	files []string
	error error
}

func (lister *fileListerMock) GetAllFiles() ([]string, error) {
	return lister.files, lister.error
}

type readCloserMock struct {
	io.Reader
	closed bool
}

func (closer *readCloserMock) Close() error {
	closer.closed = true
	return nil
}

type fileReaderMock struct {
	contents string
	error    error
}

func (reader *fileReaderMock) Open(path string) (io.ReadCloser, error) {
	return &readCloserMock{strings.NewReader(reader.contents), false}, reader.error
}

func TestReturnsErrorWhenReadErrorOccurs(t *testing.T) {
	song := &model.Song{}

	provider := Provider{FileLister: &fileListerMock{nil, errors.New("fakeerror")}}
	_, err := provider.GetLyrics(song)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestReturnsErrorWhenNoFilesProvided(t *testing.T) {
	song := &model.Song{}

	provider := Provider{FileLister: &fileListerMock{[]string{}, nil}}
	_, err := provider.GetLyrics(song)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestMatcherErrorReturnsError(t *testing.T) {
	song := &model.Song{}
	lister := &fileListerMock{[]string{"artist - title.lrc"}, nil}

	provider := Provider{
		FileLister: lister,
	}
	_, err := provider.GetLyrics(song)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestOpenerErrorReturnsError(t *testing.T) {
	song := &model.Song{Artist: "artist", Title: "title"}
	lister := &fileListerMock{[]string{"artist - title.lrc"}, nil}
	reader := &fileReaderMock{error: errors.New("fake error")}

	provider := Provider{
		FileLister: lister,
		FileReader: reader,
	}
	_, err := provider.GetLyrics(song)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestValidInputReturnsLyrics(t *testing.T) {
	song := &model.Song{Artist: "artist", Title: "title"}
	lister := &fileListerMock{[]string{"artist - title.lrc"}, nil}
	reader := &fileReaderMock{contents: "[00:01.00] line 1"}

	provider := Provider{
		FileLister: lister,
		FileReader: reader,
	}

	_, err := provider.GetLyrics(song)
	if err != nil {
		t.Errorf("Didnt expect an error, received: %s", err)
	}
}
