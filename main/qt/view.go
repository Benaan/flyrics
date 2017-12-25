package qt

import (
	"os"
	"sort"

	"github.com/benaan/flyrics/src/model"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

type View struct {
	activeLine int
	model      *LyricModel
	status     *Status
}

func (view *View) SetLyrics(lines model.Lines) {
	if view.model != nil {
		view.model.clearLines()
		for _, key := range getSortedKeys(lines) {
			qtLine := NewLine(nil)
			qtLine.SetLinetext(lines[key])
			view.model.AddLine(qtLine)
		}
	}
}

func (view *View) SetActiveLine(row int) {
	view.status.SetActiveLine(row)
}

type Status struct {
	core.QObject

	_ int `property:"activeLine"`
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

func (view *View) Present() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)
	quickcontrols2.QQuickStyle_SetStyle("material")
	app := qml.NewQQmlApplicationEngine(nil)

	view.model = NewLyricModel(nil)

	app.RootContext().SetContextProperty("LineModel", view.model)
	app.SetProperty("activeLine", core.NewQVariant7(0))

	view.status = NewStatus(nil)
	app.RootContext().SetContextProperty("status", view.status)

	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
