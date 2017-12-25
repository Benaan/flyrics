import QtQuick 2.7
import QtQuick.Window 2.0
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.3

Window{
    id: optionsWindow
    title: "Options"
    minimumWidth: 500
    minimumHeight: 300
    flags: Qt.SubWindow


    GridLayout {
        columns: 2
        width: parent.width - 20
        height: parent.height - 20
        anchors.centerIn: parent

        Label {
            text: "Lyric directory"
            Layout.alignment: Qt.AlignTop
            verticalAlignment: Text.AlignVCenter
        }
        TextField {
            text: status.lyricDirectory
            Layout.alignment:  Qt.AlignTop
            Layout.fillWidth: true
            onEditingFinished: {
                status.lyricDirectory = text.toString()
            }
        }

        Label {
            text: "GPMDP JSON api file"
            Layout.alignment: Qt.AlignTop
            verticalAlignment: Text.AlignVCenter
        }
        TextField {
            text: status.gpmdpPath
            Layout.alignment:  Qt.AlignTop
            Layout.fillWidth: true
            onEditingFinished: {
                status.gpmdpPath = text.toString()
            }
        }

        Rectangle {

            Layout.fillHeight: true

        }
    }
}
