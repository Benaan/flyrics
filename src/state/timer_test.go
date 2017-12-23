package state

import (
	"testing"

	"github.com/benaan/flyrics/src/model"
)

func TestNoLyricsReturn0(t *testing.T) {
	state := &State{}
	assertLineEquals(state, 0, t)
}

func TestNoLyricLinesReturn0(t *testing.T) {
	state := &State{lyrics: &model.Lyrics{
		Offset: 500,
	}}
	assertLineEquals(state, 0, t)
}

func TestEmptyLyricLinesReturn0(t *testing.T) {
	state := &State{lyrics: &model.Lyrics{
		Offset: 500,
		Lines:  model.Lines{},
	}}
	assertLineEquals(state, 0, t)
}

func TestFirstLineGetActivatedByDefault(t *testing.T) {
	state := setUpState()
	assertLineEquals(state, 0, t)
}

func TestFirstLineGetActivedBeforeItsTime(t *testing.T) {
	assertLineIsActive(50, 0, t)
}

func TestFirstLineIsShownAfterItsTime(t *testing.T) {
	assertLineIsActive(700, 0, t)
}

func TestSecondLineIsShownOnItsTime(t *testing.T) {
	assertLineIsActive(800, 1, t)
}

func TestThirdLineIsShownAfterItsTime(t *testing.T) {
	assertLineIsActive(1100, 2, t)
}

func TestLastLineIsShownAfterItsTime(t *testing.T) {
	assertLineIsActive(2000, 3, t)
}

func assertLineIsActive(time int, expectedLine int, t *testing.T) {
	t.Helper()
	state := setUpState()
	state.SetTime(time)
	state.status = model.PLAYING
	assertLineEquals(state, expectedLine, t)
}
func assertLineEquals(state *State, expectedLine int, t *testing.T) {
	t.Helper()
	lineNumber := state.GetActiveLine()
	if lineNumber != expectedLine {
		t.Errorf("Expected active line number to be %d received: %d", expectedLine, lineNumber)
	}
}

func setUpState() *State {
	return &State{
		lyrics: &model.Lyrics{
			Offset: 500,
			Lines: model.Lines{
				100: "line 1",
				300: "line 2",
				500: "line 3",
				800: "line 4",
			},
		},
	}
}
