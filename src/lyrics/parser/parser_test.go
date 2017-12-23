package parser

import (
	"strings"
	"testing"

	"github.com/benaan/flyrics/src/model"
)

func TestParseEmptyLyrics(t *testing.T) {
	r := strings.NewReader("")
	_, err := ParseLyrics(r)
	if err == nil {
		t.Error("The parser shouldn't be able to parse empty strings")
	}

	r = strings.NewReader("lorem ipsum")
	_, err = ParseLyrics(r)
	if err == nil {
		t.Error("The parser shouldn't be able to parse invalid strings")
	}
}

func TestParseLyricLines(t *testing.T) {
	r := strings.NewReader(`
	[ar:artist]
	[al:album]
	[ti:title]
	[by:creator]

	[00:01.00] line 1
	[01:03.00] line 2
	[02:22.10][00:11.15]line 3
	[00:10.00]
	[00:10.00][00:10.00]
	`)
	lyrics, err := ParseLyrics(r)
	if err != nil {
		t.Fatal("The parser failed to parse the lyrics:", err)
	}

	checkLine(t, lyrics, "line 1", 1000)
	checkLine(t, lyrics, "line 2", 63000)
	checkLine(t, lyrics, "line 3", 142100)
	checkLine(t, lyrics, "line 3", 11150)
	checkLine(t, lyrics, "", 10000)
	checkLine(t, lyrics, "", 10000)
	checkLine(t, lyrics, "", 10000)
}

func TestParseOffset(t *testing.T) {
	checkOffset(t, 0, `
	[00:01.00] line 1`)

	checkOffset(t, 0, `
	[offset: 0]
	[00:01.00] line 1`)

	checkOffset(t, -100, `
	[offset: -100]
	[00:01.00] line 1`)

	checkOffset(t, 12345, `
	[offset: +12345]
	[00:01.00] line 1`)

	checkOffset(t, 6789, `
	[offset: +6789]
	[00:01.00] line 1`)

}

func checkOffset(t *testing.T, expected int, input string) {
	t.Helper()
	lyrics, err := ParseLyrics(strings.NewReader(input))
	if err != nil {
		t.Fatal("The parser failed to parse the lyrics:", err)
	}

	if lyrics.Offset != expected {
		t.Error("Offset should be ", expected, ", received:", lyrics.Offset)
	}
}

func checkLine(t *testing.T, lyrics *model.Lyrics, expectedText string, time int) {
	t.Helper()
	line, ok := lyrics.Lines[time]
	if !ok {
		t.Error("Couldn't get lyrics on time:", time, "Expected:", expectedText)
	}
	if line != expectedText {
		t.Error("Wrong lyrics, expected:", expectedText, ", received:", line)
	}
}
