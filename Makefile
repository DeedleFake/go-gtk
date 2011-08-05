include $(GOROOT)/src/Make.inc

GC=${O}g -I../glib/_obj -I../gdk/_obj -I../gdkpixbuf/_obj -I$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

all:
	cd pango && gomake
	cd glib && gomake
	cd gdk && gomake
	cd gdkpixbuf && gomake
	cd gtk && gomake

install:
	cd pango && gomake install
	cd glib && gomake install
	cd gdk && gomake install
	cd gdkpixbuf && gomake install
	cd gtk && gomake install

clean:
	cd pango && gomake clean
	cd glib && gomake clean
	cd gdk && gomake clean
	cd gdkpixbuf && gomake clean
	cd gtk && gomake clean
	cd example && gomake clean

fmt_all:
	gofmt -w ./gdk/gdk.go
	gofmt -w ./gtk/gtk.go
	gofmt -w ./gdkpixbuf/gdkpixbuf.go
	gofmt -w ./glib/glib.go
	gofmt -w ./pango/pango.go

example: install
	cd example && gomake
