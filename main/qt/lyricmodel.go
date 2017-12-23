package qt

import "github.com/therecipe/qt/core"

const (
	Linetext = int(core.Qt__UserRole) + 1<<iota
)

type LyricModel struct {
	core.QAbstractListModel
	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Line                  `property:"lines"`

	_ func(*Line) `slot:"addLine"`
	_ func()      `slot:"clearLines"`
}

type Line struct {
	core.QObject

	_ string `property:"linetext"`
}

func init() {
	Line_QRegisterMetaType()
}

func (m *LyricModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Linetext: core.NewQByteArray2("linetext", len("linetext")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddLine(m.addLine)
	m.ConnectClearLines(m.clearLines)

}

func (m *LyricModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Lines()) {
		return core.NewQVariant()
	}

	var p = m.Lines()[index.Row()]

	switch role {
	case Linetext:
		{
			return core.NewQVariant14(p.Linetext())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *LyricModel) addLine(p *Line) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Lines()), len(m.Lines()))
	m.SetLines(append(m.Lines(), p))
	m.EndInsertRows()
}

func (m *LyricModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Lines())
}

func (m *LyricModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *LyricModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *LyricModel) clearLines() {
	m.BeginResetModel()
	m.SetLines([]*Line{})
	m.EndResetModel()
}
