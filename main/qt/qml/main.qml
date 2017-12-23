import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Window 2.0

ApplicationWindow {
    id: window
    property bool locked: false
    flags:   Qt.WindowStaysOnTopHint | (locked ? Qt.FramelessWindowHint : Qt.Window)
    visible: true
    title: "Flyrics"
    width: 300
    height: 400
    color: "#00000000"

    ListView {
        anchors.fill: parent
        interactive: false
        delegate: Text {
            width: parent.width
            text: linetext
            font.pixelSize: 30
            font.family: "Helvetica"
            font.bold: true
            wrapMode: Text.WordWrap
            horizontalAlignment: Text.AlignHCenter
            style: Text.Outline
            styleColor: "black"
            color: ListView.isCurrentItem ? "white": "lightgrey"
        }
        currentIndex: status.activeLine
        preferredHighlightBegin: window.height / 2 - 15
        preferredHighlightEnd: window.height / 2 + 15
        highlightRangeMode: ListView.ApplyRange
        //highlightMoveVelocity: 10
        model: LineModel



        MouseArea {
            id: mouseArea
            anchors.fill: parent
            acceptedButtons: Qt.MiddleButton
            onPressed: {
                locked = !locked
            }

        }
    }

}