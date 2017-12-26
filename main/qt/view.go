package qt

import (
	"os"
	"sort"

	"github.com/benaan/flyrics/src/config"
	"github.com/benaan/flyrics/src/lyrics"
	"github.com/benaan/flyrics/src/model"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

type View struct {
	activeLine    int
	Config        config.Manager
	LyricProvider lyrics.LyricsManager
	model         *LyricModel
	status        *Status
}

func (v *View) SetSong(song *model.Song) {
	if v.status != nil {
		songViewModel := v.status.CurrentSong()
		songViewModel.SetArtist(song.Artist)
		songViewModel.SetAlbum(song.Album)
		songViewModel.SetTitle(song.Title)
	}
}

func (v *View) SetLyricManager(manager lyrics.LyricsManager) {
	v.LyricProvider = manager
}

func (v *View) SetLyrics(lines model.Lines) {
	if v.model != nil {
		v.model.clearLines()
		for _, key := range getSortedKeys(lines) {
			qtLine := NewLine(nil)
			qtLine.SetLinetext(lines[key])
			v.model.AddLine(qtLine)
		}
	}
}

func (v *View) SetActiveLine(row int) {
	v.status.SetActiveLine(row)
}

func (v *View) Present() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)
	quickcontrols2.QQuickStyle_SetStyle("material")
	app := qml.NewQQmlApplicationEngine(nil)

	v.model = NewLyricModel(nil)

	app.RootContext().SetContextProperty("LineModel", v.model)
	app.SetProperty("activeLine", core.NewQVariant7(0))

	v.status = NewStatus(nil)
	v.status.setConfig(v.Config)
	v.status.listenToSettingsChanges()

	search := NewSearch(nil)
	search.setLyricManager(v.LyricProvider)
	app.RootContext().SetContextProperty("search", search)
	app.RootContext().SetContextProperty("status", v.status)

	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}

func getSortedKeys(lines model.Lines) []int {
	keys := make([]int, len(lines))
	i := 0
	for k := range lines {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}
