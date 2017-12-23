package viewlyrics

import (
	"io"

	"github.com/benaan/flyrics/src/lyrics/parser"
	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

type ViewLyrics struct {
	Writer util.LyricWriter
}

func (vl *ViewLyrics) GetLyrics(song *model.Song) (*model.Lyrics, error) {
	response, err := sendRequest(createRequest(song))
	if err != nil {
		return nil, err
	}
	defer response.Close()

	files, err := getFileList(response)
	if err != nil {
		return nil, err
	}

	url, err := getBestMatch(song, files)
	if err != nil {
		return nil, err
	}

	file, err := getFile(url)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	lyrics, err := parser.ParseLyrics(file)
	if err != nil {
		return nil, err
	}

	vl.Writer.Write(song, file)

	return lyrics, nil

}

func getFileList(input io.Reader) ([]*File, error) {
	decoded, err := decode(input)
	if err != nil {
		return nil, err
	}

	return parseFileList(decoded)
}
