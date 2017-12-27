import QtQuick 2.7
import QtQuick.Window 2.0
import QtQuick.Controls 1.4

ApplicationWindow {
    id: mainWindow
    property bool locked: true
    flags:   locked ? Qt.WindowStaysOnTopHint | Qt.FramelessWindowHint : Qt.Window
    visible: true
    title: "Flyrics"
    width: 300
    height: 400
    color: locked ? "transparent" : "white"

    menuBar: MenuBar {
        __contentItem.scale: locked ? 0 : 1
        Menu {
            visible: !locked
            title: "File"
            MenuItem {
                text: "Search"
                onTriggered: {
                    var component = Qt.createComponent("search.qml");
                    component.createObject(mainWindow).show();
                }
            }
            MenuItem {
                text: "Options"
                onTriggered: {
                    var component = Qt.createComponent("options.qml");
                    component.createObject(mainWindow).show();
                }
            }
            MenuItem {
                text: "Exit"
                onTriggered:  Qt.quit();
            }
        }
    }

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
        preferredHighlightBegin: mainWindow.height / 2 - 15
        preferredHighlightEnd: mainWindow.height / 2 + 15
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

