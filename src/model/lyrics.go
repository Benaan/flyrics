package model

type Lines map[int]string

type Lyrics struct {
	Offset int
	Lines  Lines
}
