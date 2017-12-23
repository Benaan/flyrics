import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Window 2.0

ApplicationWindow {
    id: window
    flags: Qt.FramelessWindowHint | Qt.WindowStaysOnTopHint
    visible: true
    title: "Flyrics"
    minimumWidth: 400
    minimumHeight: 800

    color: mouseArea.containsMouse ? "#ffcccccc" : "#00000000"


    MouseArea {
        id: mouseArea
        hoverEnabled: true
        anchors.fill: parent

        ListView {
            anchors.fill: parent
            delegate: Text {
                width: parent.width
                text: linetext
                font.pixelSize: 30
                font.family: "Helvetica"
                font.bold: ListView.isCurrentItem
                wrapMode: Text.WordWrap
                horizontalAlignment: Text.AlignHCenter
                style: Text.Outline
                styleColor: "black"
                color: ListView.isCurrentItem ? "lightgrey" : "white"
            }
            currentIndex: status.activeLine
            preferredHighlightBegin: window.height / 2 - 15
            preferredHighlightEnd: window.height / 2 + 15
            highlightRangeMode: ListView.ApplyRange
            model: LineModel
        }
    }



}
