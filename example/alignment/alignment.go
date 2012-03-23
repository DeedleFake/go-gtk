package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Alignment")
	window.Connect("destroy", gtk.MainQuit)

	notebook := gtk.Notebook()
	window.Add(notebook)

	align := gtk.Alignment(0.5, 0.5, 0.5, 0.5)
	notebook.AppendPage(align, gtk.Label("Alignment"))

	button := gtk.ButtonWithLabel("Hello World!")
	align.Add(button)

	fixed := gtk.Fixed()
	notebook.AppendPage(fixed, gtk.Label("Fixed"))

	button2 := gtk.ButtonWithLabel("Pulse")
	fixed.Put(button2, 30, 30)

	progress := gtk.ProgressBar()
	fixed.Put(progress, 30, 70)

	button.Connect("clicked", func() {
		progress.SetFraction(0.1 + 0.9*progress.GetFraction()) //easter egg
	})
	button2.Connect("clicked", func() {
		progress.Pulse()
	})

	window.ShowAll()
	window.SetSizeRequest(200, 200)

	gtk.Main()
}
