package viewlyrics

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/lyrics/parser"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type ViewLyrics struct {
	Writer util.LyricWriter
}

func (vl *ViewLyrics) GetLyrics(song *model.Song) (*model.Lyrics, error) {
	files, err := getFileList(song)
	if err != nil {
		return nil, err
	}

	url, err := getBestMatch(song, files)
	if err != nil {
		return nil, err
	}

	return vl.getLyricsFromUrl(url, song)

}

func getFileList(song *model.Song) ([]*File, error) {
	response, err := sendRequest(createRequest(song))
	if err != nil {
		return nil, err
	}
	defer response.Close()

	return getFileListFromResponse(response)
}

func (vl *ViewLyrics) GetList(song *model.Song) []*lyrics.File {
	var list []*lyrics.File

	files, err := getFileList(song)
	if err != nil {
		return list
	}

	files = filterFiles(files)

	for _, file := range files {
		list = append(list, &lyrics.File{
			Song: &model.Song{
				Artist: file.Artist,
				Album:  file.Album,
				Title:  file.Title,
			},
			Downloads: file.Downloads,
			Rating:    file.Rating,
			Source:    "ViewLyrics",
			Get: func() (*model.Lyrics, error) {
				return vl.getLyricsFromUrl(file.Link, song)
			},
		})
	}

	return list
}

func (vl *ViewLyrics) getLyricsFromUrl(url string, song *model.Song) (*model.Lyrics, error) {
	file, err := getFile(url)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	lrcs, err := parser.ParseLyrics(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	vl.Writer.Write(song, bytes.NewReader(data))
	return lrcs, nil
}

func getFileListFromResponse(input io.Reader) ([]*File, error) {
	decoded, err := decode(input)
	if err != nil {
		return nil, err
	}

	return parseFileList(decoded)
}
