package parser

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/benaan/flyrics/src/model"
)

func ParseLyrics(reader io.Reader) (*model.Lyrics, error) {
	lyrics := model.Lyrics{}
	lines := make(map[int]string)
	lyrics.Lines = lines

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		parseLine(&lyrics, scanner.Text())
	}

	if len(lyrics.Lines) == 0 {
		return &lyrics, errors.New("No lyrics in file")
	}

	return &lyrics, nil

}
func parseLine(lyrics *model.Lyrics, line string) {
	line = strings.TrimSpace(line)
	if !isValidLine(line) {
		return
	}
	extractOffsetFromLine(line, lyrics)
	extractLyricFromLine(line, lyrics)

}
func extractLyricFromLine(line string, lyrics *model.Lyrics) {
	times := getTime(line)
	if len(times) == 0 {
		return
	}
	text := getText(line)
	for _, time := range times {
		lyrics.Lines[time] = text
	}
}
func isValidLine(line string) bool {
	return strings.HasPrefix(line, "[")
}
func extractOffsetFromLine(line string, lyrics *model.Lyrics) {
	offset, err := getOffset(line)
	if err == nil {
		lyrics.Offset = offset
	}
}
func getOffset(line string) (int, error) {
	re := regexp.MustCompile(`\[offset:(.*)]`)
	result := re.FindAllStringSubmatch(line, 1)
	if len(result) > 0 {
		return strconv.Atoi(strings.TrimSpace(result[0][1]))
	}
	return 0, errors.New("Offset not found")

}
func getText(line string) string {
	re := regexp.MustCompile(`(\[\d+:\d+\.\d+])+(.*)`)
	return strings.TrimSpace(re.FindAllStringSubmatch(line, 1)[0][2])
}

func getTime(line string) []int {
	var times []int

	re := regexp.MustCompile(`\[(\d+):(\d+)\.(\d+)]`)

	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		total := 0
		for i := 1; i <= 3; i++ {
			time, _ := strconv.Atoi(match[i])

			switch i {
			case 1:
				total += time * 60000
			case 2:
				total += time * 1000
			case 3:
				total += time * 10
				times = append(times, total)
			}
		}

	}

	return times
}
