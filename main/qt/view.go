package qt

import (
	"fmt"
	"os"
	"sort"

	"github.com/benaan/flyrics/src/config"
	"github.com/benaan/flyrics/src/model"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

type View struct {
	activeLine int
	Config     config.Manager
	model      *LyricModel
	status     *Status
}

type Status struct {
	core.QObject

	_ int    `property:"activeLine"`
	_ string `property:"lyricDirectory"`
	_ string `property:"gpmdpPath"`
}

type Search struct {
	core.QObject

	_ *LyricListModel `property:"lyricList"`

	_ func() `constructor:"init"`

	_ func(artist, album, title string) `slot:"searchLyrics"`
}

func (s *Search) init() {
	list := NewLyricListModel(nil)
	for i := 0; i < 2; i++ {

		lrc := NewLyric(nil)
		lrc.SetAlbum(fmt.Sprintf("album %d", i))
		lrc.SetTitle(fmt.Sprintf("title %d", i))
		lrc.SetArtist(fmt.Sprintf("artist %d", i))
		lrc.SetDownloads(i)
		lrc.SetRating(fmt.Sprintf("%d", i))
		list.AddLyric(lrc)
	}
	s.SetLyricList(list)
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
	v.listenToConfigChanges()

	search := NewSearch(nil)
	search.ConnectSearchLyrics(func(artist, album, title string) {
		list := NewLyricListModel(nil)
		for i := 0; i < 10; i++ {

			lrc := NewLyric(nil)
			lrc.SetAlbum(fmt.Sprintf("album %d %s", i, album))
			lrc.SetTitle(fmt.Sprintf("title %d %s", i, title))
			lrc.SetArtist(fmt.Sprintf("artist %d %s", i, artist))
			lrc.SetDownloads(i)
			lrc.SetRating(fmt.Sprintf("%d", i))
			list.AddLyric(lrc)
		}
		search.SetLyricList(list)
	})
	app.RootContext().SetContextProperty("search", search)
	app.RootContext().SetContextProperty("status", v.status)

	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
func (v *View) listenToConfigChanges() {
	v.status.SetLyricDirectory(v.Config.GetStringConfig(config.LyricDirectory))
	v.status.ConnectLyricDirectoryChanged(func(lyricDirectory string) {
		v.Config.SetConfig(config.LyricDirectory, lyricDirectory)
	})

	v.status.SetGpmdpPath(v.Config.GetStringConfig(config.GpmdpPath))
	v.status.ConnectGpmdpPathChanged(func(path string) {
		v.Config.SetConfig(config.GpmdpPath, path)
	})
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
