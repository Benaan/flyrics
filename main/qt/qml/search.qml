import QtQuick 2.7
import QtQuick.Window 2.0
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.3

ApplicationWindow{
    id: searchWindow

    title: "Options"
    minimumWidth: 500
    minimumHeight: 300
    flags: Qt.SubWindow
    toolBar: ToolBar {
        RowLayout {
            id: rows
            spacing: 1

            anchors.fill: parent
            TextField {
                id: artistInput
                height: searchButton.height
                placeholderText: "Artist"
            }

             TextField {
                id: albumInput
                height: searchButton.height
                placeholderText: "Album"
            }

            TextField {
                id: titleInput
                height: searchButton.height
                placeholderText: "Title"
            }

            Button  {
                id: searchButton
                text: "Search"
                onClicked: {
                   search.searchLyrics(artistInput.text, albumInput.text, titleInput.text)
                }
            }
        }
    }

    TableView {
        id: resultView
        width: parent.width
        height: parent.height
        frameVisible: false
        sortIndicatorVisible: true
        model: search.lyricList

        TableViewColumn {
            role: "artist"
            title: "Artist"
        }
        TableViewColumn {
            role: "album"
            title: "Album"
        }
        TableViewColumn {
            role: "title"
            title: "Title"
        }
        TableViewColumn {
            role: "downloads"
            title: "Downloads"
        }
        TableViewColumn {
            role: "rating"
            title: "Rating"
        }
        TableViewColumn {
            role: "source"
            title: "Source"
        }
    }
}
