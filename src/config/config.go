package config

const LyricDirectory = "LyricDirectory"
const GpmdpPath = "GPMDP path"

type Reader interface {
	GetStringConfig(key string) string
}

type Writer interface {
	SetConfig(name string, value interface{})
}

type Manager interface {
	Reader
	Writer
}
