package main

import (
	"reflect"
	"testing"

	"github.com/benaan/flyrics/src/model"
)

func TestDisplaysFirst3Lines(t *testing.T) {
	view := consoleView{
		lines: model.Lines{
			100: "Line 1",
			200: "Line 2",
			300: "Line 3",
			400: "Line 4",
			500: "Line 5",
		},
		activeLine: 1,
	}

	lines := view.getRenderLines()
	if !reflect.DeepEqual(lines, []int{100, 200, 300}) {
		t.Errorf("Expected lines 100, 200 and 300 to be displayed, received %s", lines)
	}
}

func TestDisplaysFirst2Lines(t *testing.T) {
	view := consoleView{
		lines: model.Lines{
			100: "Line 1",
			200: "Line 2",
			300: "Line 3",
			400: "Line 4",
			500: "Line 5",
		},
		activeLine: 0,
	}

	lines := view.getRenderLines()
	if !reflect.DeepEqual(lines, []int{100, 200}) {
		t.Errorf("Expected lines 100 and 200 to be displayed, received %s", lines)
	}
}

func TestDisplaysLast2Lines(t *testing.T) {
	view := consoleView{
		lines: model.Lines{
			100: "Line 1",
			200: "Line 2",
			300: "Line 3",
			400: "Line 4",
			500: "Line 5",
		},
		activeLine: 4,
	}

	lines := view.getRenderLines()
	if !reflect.DeepEqual(lines, []int{400, 500}) {
		t.Errorf("Expected lines 400 and 500 to be displayed, received %s", lines)
	}
}

func TestActiveLine(t *testing.T) {
	view := consoleView{
		lines: model.Lines{
			100: "Line 1",
			200: "Line 2",
			300: "Line 3",
			400: "Line 4",
			500: "Line 5",
		},
		activeLine: 1,
	}

	activeLine := view.getActiveLine()
	if activeLine != 200 {
		t.Errorf("Expected active line to be 200, received %s", activeLine)
	}
}
