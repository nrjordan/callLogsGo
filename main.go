package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//var menuitem *gtk.MenuItem
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Call Logs")
	window.SetIconName("Call Logs")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("exit called", ctx.Data().(string))
		gtk.MainQuit()
	}, "all normal")

	//vbox?
	vbox := gtk.NewVBox(false, 1)

	//menubar
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//vpaned
	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	//gtkFrame
	frame1 := gtk.NewFrame("Report Generator")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	vpaned.Pack1(frame1, false, false)

	//Buttons
	buttons := gtk.NewHBox(false, 1)

	button := gtk.NewButtonWithLabel("Call Logs")
	button.Clicked(func() {
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose File...",
			button.GetTopLevelAsWindow(),
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)
		filter := gtk.NewFileFilter()
		filter.AddPattern("*.xlsx")
		filechooserdialog.AddFilter(filter)
		filechooserdialog.Response(func() {
			//run functions
			create(filechooserdialog.GetFilename())
			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
	})
	buttons.Add(button)

	framebox1.PackStart(buttons, false, false, 0)

	window.Add(vbox)
	window.SetSizeRequest(300, 100)
	window.ShowAll()
	gtk.Main()
}
