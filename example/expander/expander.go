package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("We love Expander")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.VBox(true, 0)
	vbox.SetBorderWidth(5)
	expander := gtk.Expander("dan the ...")
	expander.Add(gtk.Label("404 contents not found"))
	vbox.PackStart(expander, false, false, 0)

	window.Add(vbox)
	window.ShowAll()

	gtk.Main()
}
