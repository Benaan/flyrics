package model

type Status int

const (
	STOPPED Status = iota
	PAUSED
	PLAYING
)

type Metadata struct {
	Status Status
	Song   *Song
	Time   int
}
