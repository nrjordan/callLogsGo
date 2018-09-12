package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func main() {
	mw := new(MyMainWindow)

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Report Generator",
		MinSize:  Size{300, 100},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "Call Log",
				OnClicked: func() {
					//file select
					mw.openActionTriggered()
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

func (mw *MyMainWindow) openActionTriggered() {
	if err := mw.openFile(); err != nil {
		log.Print(err)
	}
}

func (mw *MyMainWindow) openFile() error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.prevFilePath
	dlg.Filter = "xlxs Files (*.xlsx)|*.xlsx"
	dlg.Title = "Select the Call Log file"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	mw.prevFilePath = dlg.FilePath

	create(dlg.FilePath)

	walk.MsgBox(mw, "Finished", "File created.", walk.MsgBoxIconInformation)

	return nil
}

type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
}
