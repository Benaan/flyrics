package viewlyrics

import (
	"errors"
	"sort"
	"strings"

	"github.com/benaan/flyrics/src/model"
	"github.com/benaan/flyrics/src/util"
)

func filterFiles(list []*File) []*File {
	var files []*File
	for _, file := range list {
		if strings.HasSuffix(file.Link, ".lrc") {
			files = append(files, file)
		}
	}
	return files
}

func filterSong(song *model.Song, list []*File) []*File {
	artist := util.ToMatchable(song.Artist)
	title := util.ToMatchable(song.Title)
	var files []*File
	for _, file := range list {
		if util.ToMatchable(file.Artist) == artist && util.ToMatchable(file.Title) == title {
			files = append(files, file)
		}
	}

	return filterOnAlbum(song.Album, files)
}

func filterOnAlbum(album string, list []*File) []*File {
	cleanedAlbum := util.ToMatchable(album)
	if cleanedAlbum == "" {
		return list
	}

	var files []*File
	for _, file := range list {
		if util.ToMatchable(file.Album) == cleanedAlbum {
			files = append(files, file)
		}
	}
	if len(files) > 0 {
		return files
	}
	return list
}

func sortFiles(files []*File) {
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Downloads > files[j].Downloads
	})
}

func getBestMatch(song *model.Song, files []*File) (string, error) {
	lrcFiles := filterFiles(files)
	if len(lrcFiles) == 0 {
		return "", errors.New("No lyric files found")
	}
	cleanedFiles := filterSong(song, lrcFiles)
	sortFiles(cleanedFiles)
	if len(cleanedFiles) == 0 {
		return "", errors.New("No valid matches found")
	}
	return cleanedFiles[0].Link, nil
}
