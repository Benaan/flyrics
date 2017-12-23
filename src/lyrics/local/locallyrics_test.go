package local

import (
	"reflect"
	"testing"

	"github.com/benaan/flyrics/src/lyrics/local/filematcher"
	"github.com/benaan/flyrics/src/model"
)

func assertErrorIsReturned(song *model.Song, files []string, t *testing.T) {
	t.Helper()
	sut := lyricFinder{}
	if _, err := sut.find(song, files); err == nil {
		t.Error("Expected an error")
	}
}

func TestNoFiles_ReturnsError(t *testing.T) {
	song := &model.Song{
		Artist: "artist",
		Album:  "album",
		Title:  "title",
	}
	var files []string
	assertErrorIsReturned(song, files, t)
}

func TestNoSong_ReturnsError(t *testing.T) {
	assertErrorIsReturned(nil, []string{"file1.lrc"}, t)
}

func TestEmptySong_ReturnsError(t *testing.T) {
	song := &model.Song{
		Artist: "",
		Album:  "",
		Title:  "",
	}
	files := []string{"file1.lrc"}
	assertErrorIsReturned(song, files, t)
}

func TestSongWithNoFile_ReturnsError(t *testing.T) {
	song := &model.Song{
		Artist: "artist",
		Title:  "title",
	}
	files := []string{"file1.lrc"}
	assertErrorIsReturned(song, files, t)
}

type FileMatcherSpy struct {
	callCount    int
	lastFilename string
}

func (matcher *FileMatcherSpy) Matches(fileName string, song *model.Song) bool {
	matcher.callCount++
	matcher.lastFilename = fileName
	return false
}

func (*FileMatcherSpy) GetCertainty() filematcher.Certainty {
	return 0
}

type FakeFileMatcher struct {
}

func (*FakeFileMatcher) Matches(fileName string, song *model.Song) bool {
	return true
}

func (*FakeFileMatcher) GetCertainty() filematcher.Certainty {
	return 1
}

func TestFilesAreIterated(t *testing.T) {
	song := &model.Song{
		Artist: "artist",
		Title:  "title",
	}
	files := []string{"file1.lrc", "file2.lrc", "file3.lrc"}
	spy := &FileMatcherSpy{}
	sut := lyricFinder{Matchers: []filematcher.FileMatcher{spy}}
	sut.find(song, files)
	if spy.callCount != 3 {
		t.Error("Expected 3 files to be iterated, got", spy.callCount)
	}
}

func TestBaseDirectoryIsStrippedFilesAreIterated(t *testing.T) {
	song := &model.Song{
		Artist: "artist",
		Title:  "title",
	}
	files := []string{"/path/to/file/file1.lrc"}
	spy := &FileMatcherSpy{}
	sut := lyricFinder{Matchers: []filematcher.FileMatcher{spy}}
	sut.find(song, files)
	if spy.lastFilename != "file1" {
		t.Error("directory should have been stripped, but last filename was:", spy.lastFilename)
	}
}

func TestFileIsMatched(t *testing.T) {
	song := &model.Song{
		Artist: "artist",
		Title:  "title",
	}
	files := []string{"file.lrc"}
	matcher := &FakeFileMatcher{}
	sut := lyricFinder{Matchers: []filematcher.FileMatcher{matcher}}
	matches, err := sut.find(song, files)
	if err != nil {
		t.Error("Didnt expect an error, received:", err)
	}
	if len(matches) != 1 {
		t.Error("Expected 1 match, got", len(matches))
	}
	expected := []*filematcher.Match{{Path: "file.lrc", Certainty: 1}}
	if !reflect.DeepEqual(matches, expected) {
		t.Errorf("received %s expected %s", matches, expected)
	}

}
