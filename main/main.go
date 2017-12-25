package main

import (
	"github.com/benaan/flyrics/main/qt"
	"github.com/benaan/flyrics/src/application"
)

func main() {
	application.Run(&qt.View{Config: application.Config})
}
