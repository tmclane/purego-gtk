//go:build linux

package application

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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
	_ = os.Setenv("GDK_BACKEND", "x11")

	var err error
	gtk, err = purego.Dlopen(gtk4, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err == nil {
		version = 4
		return
	}

	log.Println("Failed to open GTK4: Falling back to GTK3")
	gtk, err = purego.Dlopen(gtk3, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	version = 3
}

func New(name string) *Application {
	identifier := fmt.Sprintf("io.puregotest.%s", strings.Replace(name, " ", "-", -1))
	// FIXME: Flags?
	var gtkNew func(string, uint) uintptr
	purego.RegisterLibFunc(&gtkNew, gtk, "gtk_application_new")

	a := &Application{
		application: gtkNew(identifier, 0),
		windows:     map[string]*Window{},
	}

	var g_signal_connect func(uintptr, string, uintptr, uintptr, bool, int) int
	purego.RegisterLibFunc(&g_signal_connect, gtk, "g_signal_connect_data")
	g_signal_connect(a.application, "activate", purego.NewCallback(a.activate), a.application, false, 0)

	return a
}

type Application struct {
	application uintptr
	windows     map[string]*Window
	windowLock  sync.Mutex
}

func (a *Application) activate() {
	fmt.Println("Application.activate - callback")
}

func (a *Application) Run(argc int, argv []string) int {
	var run func(uintptr, int, []string) int
	purego.RegisterLibFunc(&run, gtk, "g_application_run")
	var hold func(uintptr)
	purego.RegisterLibFunc(&hold, gtk, "g_application_hold")
	hold(a.application)

	status := run(a.application, argc, argv)
	//g_object_unref(app)
	return status
}

func (g *Application) NewWindow(name string) *Window {
	g.windowLock.Lock()
	defer g.windowLock.Unlock()
	window := &Window{
		application: g,
		name:        name,
	}
	g.windows[name] = window
	return window
}

type Window struct {
	application *Application
	name        string
	window      uintptr
	height      uint
	width       uint
}

func (w *Window) Show() *Window {
	if w.window == 0 {
		var windowNew func(uintptr) uintptr
		purego.RegisterLibFunc(&windowNew, gtk, "gtk_application_window_new")
		w.window = windowNew(w.application.application)
		w.SetSize(w.height, w.width)
		w.SetTitle(w.name)
	}
	var show func(uintptr)
	if version == 3 {
		purego.RegisterLibFunc(&show, gtk, "gtk_widget_show")
	} else {
		purego.RegisterLibFunc(&show, gtk, "gtk_window_present")
	}
	show(w.window)
	return w
}

func (w *Window) SetTitle(title string) *Window {
	var set func(uintptr, string)
	purego.RegisterLibFunc(&set, gtk, "gtk_window_set_title")
	set(w.window, title)
	return w
}

func (w *Window) SetSize(height, width uint) *Window {
	w.height = height
	w.width = width
	if w.window != 0 {
		var set func(uintptr, uint, uint)
		purego.RegisterLibFunc(&set, gtk, "gtk_window_set_default_size")
		set(w.window, w.height, w.width)
	}
	return w
}
