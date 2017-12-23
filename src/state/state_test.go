package state

import (
	"testing"

	"github.com/benaan/flyrics/src/model"
)

func TestSetSong(t *testing.T) {
	song := &model.Song{Title: "title"}
	state := State{}
	state.SetSong(song)
	if state.song != song {
		t.Error("Expected song to be %s, but received %s", song, state.song)
	}
}

func TestSetLyrics(t *testing.T) {
	lyrics := &model.Lyrics{
		Lines: map[int]string{0: "line 1"},
	}
	state := State{}
	state.SetLyrics(lyrics)
	if state.lyrics != lyrics {
		t.Errorf("Expected lyrics to be %s, received %s", lyrics, state.lyrics)
	}
}

func TestStateIsStoppedByDefault(t *testing.T) {
	state := State{}
	if state.status != model.STOPPED {
		t.Errorf("Expected status to be stopped, received: %s", state.status)
	}
}

func TestSetStatus(t *testing.T) {
	state := State{}
	state.SetStatus(model.PLAYING)
	if state.status != model.PLAYING {
		t.Errorf("Expected status to be playing, received: %s", state.status)
	}
}

func TestTimeIsZeroByDefault(t *testing.T) {
	state := State{}
	if state.time != 0 {
		t.Errorf("Expected time to be 0, received %d", state.time)
	}
}

func TestSetTime(t *testing.T) {
	state := State{}
	state.SetTime(2)
	if state.time != 2 {
		t.Errorf("Expected time to be 2, received %d", state.time)
	}
}
