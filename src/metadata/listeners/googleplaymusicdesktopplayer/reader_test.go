package googleplaymusicdesktopplayer

import (
	"errors"
	"io"
	"strings"
	"testing"
)

var fakeError = errors.New("Fake error")

type openerMock struct {
	reader   io.ReadCloser
	error    error
	lastPath string
}

func (opener *openerMock) Open(path string) (io.ReadCloser, error) {
	opener.lastPath = path
	return opener.reader, opener.error
}

type closerSpy struct {
	io.Reader
	closed bool
}

func (closer *closerSpy) Close() error {
	closer.closed = true
	return nil
}

type confMock struct {
}

func (confMock) GetStringConfig(key string) string {
	return ""
}

func TestReadErrorIsHandled(t *testing.T) {
	opener := &openerMock{error: fakeError}
	reader := Reader{Config: &confMock{}, Opener: opener}
	_, err := reader.GetMetadata()
	if err != fakeError {
		t.Error("Expected a fake error, received:", err)
	}
}

func TestFileIsClosed(t *testing.T) {
	spy := &closerSpy{Reader: strings.NewReader("")}
	opener := &openerMock{reader: spy}

	reader := Reader{Config: &confMock{}, Opener: opener}
	reader.GetMetadata()

	if !spy.closed {
		t.Error("The file was not closed")
	}
}

type configMock struct {
}

func (*configMock) GetStringConfig(key string) string {
	return "/path/to/file"
}

func TestReadsFileInPath(t *testing.T) {
	spy := &closerSpy{Reader: strings.NewReader("")}
	opener := &openerMock{reader: spy}

	reader := Reader{
		Config: &configMock{},
		Opener: opener,
	}
	reader.GetMetadata()

	if opener.lastPath != "/path/to/file" {
		t.Errorf("Expect last file to be \"/path/to/file\", received: %s", opener.lastPath)
	}
}
