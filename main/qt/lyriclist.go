package qt

import "github.com/therecipe/qt/core"

const (
	Artist = int(core.Qt__UserRole) + 1<<iota
	Album
	Title
	Downloads
	Rating
	Source
)

type LyricListModel struct {
	core.QAbstractTableModel
	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Lyric                 `property:"lyrics"`

	_ func(*Lyric) `slot:"addLyric"`
}

type Lyric struct {
	core.QObject

	_ string `property:"artist"`
	_ string `property:"album"`
	_ string `property:"title"`
	_ int    `property:"downloads"`
	_ string `property:"rating"`
	_ string `property:"source"`
}

func init() {
	Lyric_QRegisterMetaType()
}

func (m *LyricListModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Artist:    core.NewQByteArray2("artist", len("artist")),
		Album:     core.NewQByteArray2("album", len("album")),
		Title:     core.NewQByteArray2("title", len("title")),
		Downloads: core.NewQByteArray2("downloads", len("downloads")),
		Rating:    core.NewQByteArray2("rating", len("rating")),
		Source:    core.NewQByteArray2("source", len("source")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddLyric(m.addLyric)

}

func (m *LyricListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Lyrics()) {
		return core.NewQVariant()
	}

	var p = m.Lyrics()[index.Row()]

	switch role {
	case Artist:
		{
			return core.NewQVariant14(p.Artist())
		}
	case Album:
		{
			return core.NewQVariant14(p.Album())
		}
	case Title:
		{
			return core.NewQVariant14(p.Title())
		}
	case Downloads:
		{
			return core.NewQVariant7(p.Downloads())
		}
	case Rating:
		{
			return core.NewQVariant14(p.Rating())
		}
	case Source:
		{
			return core.NewQVariant14(p.Source())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *LyricListModel) addLyric(p *Lyric) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Lyrics()), len(m.Lyrics()))
	m.SetLyrics(append(m.Lyrics(), p))
	m.EndInsertRows()
}

func (m *LyricListModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Lyrics())
}

func (m *LyricListModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *LyricListModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}
