package googleplaymusicdesktopplayer

import (
	"encoding/json"
	"io"

	"github.com/benaan/flyrics/src/model"
)

type Parser struct {
}

type GPMDPJson struct {
	Playing bool
	Song    *model.Song
	Time    *Time
}

type Time struct {
	Current int
	Total   int
}

func (parser Parser) Parse(input io.Reader) (*model.Metadata, error) {
	parsedJson := &GPMDPJson{}
	decoder := json.NewDecoder(input)
	err := decoder.Decode(parsedJson)
	if err != nil {
		return nil, err
	}
	return convertJsonToMetadata(parsedJson), err
}

func convertJsonToMetadata(json *GPMDPJson) *model.Metadata {
	return &model.Metadata{
		Status: determineStatus(json.Playing, json.Time.Current),
		Song:   json.Song,
		Time:   json.Time.Current,
	}
}

func determineStatus(playing bool, time int) model.Status {
	if playing {
		return model.PLAYING
	}
	if time == 0 {
		return model.STOPPED
	}
	return model.PAUSED
}
