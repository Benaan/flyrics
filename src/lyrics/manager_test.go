package lyrics

import (
	"reflect"
	"testing"

	"github.com/benaan/flyrics/src/model"
)

type FakeLyricProvider struct {
	lyrics *model.Lyrics
	error  error
}

func (provider *FakeLyricProvider) GetList(song *model.Song) []*File {
	panic("implement me")
}

func (provider *FakeLyricProvider) GetLyricList(song *model.Song) []*File {
	return nil
}

func (provider *FakeLyricProvider) GetLyrics(song *model.Song) (*model.Lyrics, error) {
	return provider.lyrics, provider.error
}

func TestGetLyrics(t *testing.T) {
	lyrics := &model.Lyrics{Lines: map[int]string{1: "line 1"}}
	provider := &FakeLyricProvider{lyrics, nil}

	song := &model.Song{}
	output := make(chan model.Lyrics)

	manager := Manager{
		Providers: []Provider{provider},
		Output:    output,
	}
	go manager.GetLyrics(song)

	if result := <-output; !reflect.DeepEqual(result, *lyrics) {
		t.Errorf("Expected fake lyrics, received %s", result)
	}
}
