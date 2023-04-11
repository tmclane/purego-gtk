package application

import "sync"

type Application struct {
	application uintptr
	windows     map[string]*Window
	windowLock  sync.Mutex
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
