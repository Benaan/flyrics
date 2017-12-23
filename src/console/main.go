package main

import (
	"sort"
	"time"

	"github.com/benaan/flyrics/src/application"
	"github.com/benaan/flyrics/src/model"
	"github.com/buger/goterm"
)

func main() {
	application.Run(&consoleView{})
}

type consoleView struct {
	lines      model.Lines
	activeLine int
}

func (view *consoleView) Present() {
	for {
		view.draw()
		time.Sleep(time.Second)

	}
}

func (view *consoleView) SetLyrics(lines model.Lines) {
	view.lines = lines
}

func (view *consoleView) SetActiveLine(row int) {
	view.activeLine = row
	view.draw()
}

func (view *consoleView) getSortedKeys() []int {
	keys := make([]int, len(view.lines))
	i := 0
	for k := range view.lines {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func (view *consoleView) getRenderLines() []int {
	lines := []int{}
	keys := view.getSortedKeys()
	for i, key := range keys {
		if i >= view.activeLine-1 && i <= view.activeLine+1 {
			lines = append(lines, key)
		}
	}
	return lines
}

func (view *consoleView) getActiveLine() int {
	keys := view.getSortedKeys()
	for i, key := range keys {
		if i == view.activeLine {
			return key
		}
	}
	return 0
}

func (view *consoleView) draw() {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	activeLine := view.getActiveLine()
	for _, line := range view.getRenderLines() {
		text, ok := view.lines[line]
		if ok {
			if line == activeLine {
				goterm.Println(goterm.Bold(text))
			} else {
				goterm.Println(text)
			}

		}
	}

	goterm.Flush()
}
