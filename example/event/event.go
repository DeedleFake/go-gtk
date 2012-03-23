package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Events")
	window.Connect("destroy", gtk.MainQuit)

	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		kev := *(**gdk.EventKey)(unsafe.Pointer(&arg))
		println("key-press-event:", kev.Keyval)
	})
	window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		mev := *(**gdk.EventMotion)(unsafe.Pointer(&arg))
		println("motion-notify-event:", int(mev.X), int(mev.Y))
	})

	window.SetEvents(int(gdk.GDK_POINTER_MOTION_MASK | gdk.GDK_POINTER_MOTION_HINT_MASK | gdk.GDK_BUTTON_PRESS_MASK))
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk.Main()
}
