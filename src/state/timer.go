package state

import (
	"sort"

	"github.com/benaan/flyrics/src/model"
)

func (state *State) GetActiveLine() int {
	if state.status == model.STOPPED {
		return 0
	}
	if state.lyrics != nil {
		expectedTime := state.time - state.lyrics.Offset
		keys := getSortedKeys(state)
		return getBestActiveLyricLine(expectedTime, keys)
	}
	return 0
}

func getBestActiveLyricLine(currentTime int, keys []int) int {
	for i, time := range keys {
		if time > currentTime {
			if i == 0 {
				return 0
			}
			return i - 1
		}
	}
	return len(keys) - 1
}

func getSortedKeys(state *State) []int {
	keys := make([]int, len(state.lyrics.Lines))
	i := 0
	for k := range state.lyrics.Lines {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}
