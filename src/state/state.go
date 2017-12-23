package state

import (
	"github.com/benaan/flyrics/src/model"
)

type State struct {
	song   *model.Song
	lyrics *model.Lyrics
	status model.Status
	time   int
}

type LyricsQuery struct {
	stop     chan bool
	metadata model.Metadata
	lyrics   model.Lyrics
}

func (state *State) GetSong() *model.Song {
	return state.song
}

func (state *State) SetSong(song *model.Song) {
	state.song = song
}

func (state *State) SetLyrics(lyrics *model.Lyrics) {
	state.lyrics = lyrics
}

func (state *State) GetLyrics() *model.Lyrics {
	return state.lyrics
}

func (state *State) SetStatus(status model.Status) {
	state.status = status
}

func (state *State) GetStatus() model.Status {
	return state.status
}

func (state *State) SetTime(time int) {
	state.time = time
}
func (state *State) GetTime() int {
	return state.time
}
