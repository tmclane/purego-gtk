package gtk

import (
	"fmt"
	"log"

	"github.com/ebitengine/purego"
)

const (
	gtk3 = "libgtk-3.so"
	gtk4 = "libgtk-4.so"
)

var (
	gtk     uintptr
	version int
)

func init() {
	var err error
	gtk, err = purego.Dlopen(gtk4, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err == nil {
		version = 4
		return
	}

	log.Println("Failed to open GTK4: Falling back to GTK3")
	gtk, err = purego.Dlopen(gtk3, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}

func NewApplication(identifier string) *Application {
	// FIXME: Flags?
	var gtkNew func(string, uint) uintptr
	purego.RegisterLibFunc(&gtkNew, gtk, "gtk_application_new")
	return &Application{
		application: gtkNew(identifier, 0),
	}
}

type Application struct {
	application uintptr
}

func (a *Application) activate() {
	fmt.Println("Application.activate()")
	var hold func(uintptr)
	purego.RegisterLibFunc(&hold, gtk, "g_application_hold")
	hold(a.application)
}

func (a *Application) Run(argc int, argv []string) int {
	var run func(uintptr, int, []string) int
	purego.RegisterLibFunc(&run, gtk, "g_application_run")

	/*
	  GObject* instance,
	  const gchar* detailed_signal,
	  GCallback c_handler,
	  gpointer data,
	  GClosureNotify destroy_data,
	  GConnectFlags connect_flags
	*/
	var signal func(uintptr, string, uintptr, uintptr, bool, int) int
	purego.RegisterLibFunc(&signal, gtk, "g_signal_connect_data")
	signal(a.application, "activate", purego.NewCallback(a.activate), a.application, false, 0)

	status := run(a.application, argc, argv)
	//g_object_unref(app)
	return status
}

type Window struct {
	name   string
	window uintptr
	height uint
	width  uint
}

func (w *Window) Show() *Window {
	var show func(uintptr)
	purego.RegisterLibFunc(&show, gtk, "gtk_widget_show_all")

	show(w.window)
	return w
}

func (w *Window) SetTitle(title string) *Window {
	var set func(uintptr, string)
	purego.RegisterLibFunc(&set, gtk, "gtk_window_set_title")
	set(w.window, title)
	return w
}

func (w *Window) SetSize(height, width int) *Window {
	var set func(uintptr, int, int)
	purego.RegisterLibFunc(&set, gtk, "gtk_window_set_default_size")
	set(w.window, height, width)
	return w
}

func (g *Application) NewWindow(name string) *Window {
	var gtkNew func(uintptr) uintptr
	purego.RegisterLibFunc(&gtkNew, gtk, "gtk_application_window_new")

	window := &Window{
		window: gtkNew(g.application),
	}
	return window
}
