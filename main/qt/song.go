package qt

import "github.com/therecipe/qt/core"

type Song struct {
	core.QObject

	_ string `property:"artist"`
	_ string `property:"album"`
	_ string `property:"title"`
}

func init() {
	Lyric_QRegisterMetaType()
}
