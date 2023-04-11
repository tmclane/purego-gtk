//go:build darwin

package application

import (
	"fmt"
	"time"

	"github.com/ebitengine/purego/objc"
)

// enum NSApplicationActivationPolicy
const (
	NSApplicationActivationPolicyRegular   int = 0
	NSApplicationActivationPolicyAccessory int = 1
	NSApplicationActivationPolicyERROR     int = 2
)

// enum NSWindowStyleMask
const (
	NSWindowStyleMaskBorderless     int = 0
	NSWindowStyleMaskTitled         int = 1 << 0
	NSWindowStyleMaskClosable       int = 1 << 1
	NSWindowStyleMaskMiniaturizable int = 1 << 2
	NSWindowStyleMaskResizable      int = 1 << 3
)

// enum NSBackingStoreType
const (
	NSBackingStoreBuffered int = 2
)

func init() {

}

func New(name string) *Application {
	//	identifier := fmt.Sprintf("io.puregotest.%s", strings.Replace(name, " ", "-", -1))
	// id app = [NSApplication sharedApplication];
	app := objc.ID(objc.GetClass("NSApplication")).Send(objc.RegisterName("sharedApplication"))

	// [app setActivationPolicy:NSApplicationActivationPolicyRegular];
	app.Send(objc.RegisterName("setActivationPolicy:"), NSApplicationActivationPolicyRegular)

	a := &Application{
		application: uintptr(app),
		windows:     map[string]*Window{},
	}

	return a
}

func (a *Application) activate() {
	fmt.Println("Application.activate - callback")
}

func (a *Application) Run(argc int, argv []string) int {
	app := objc.ID(a.application)
	app.Send(objc.RegisterName("activateIgnoringOtherApps:"))

	// [app activateIgnoringOtherApps:YES];
	app.Send(objc.RegisterName(""), true)

	app.Send(objc.RegisterName("run"))
	for {
		time.Sleep(1 * time.Second)
	}
	return 0
}

func (w *Window) Show() *Window {
	fmt.Println("window.Show()", w.window)
	if w.window == 0 {
		//struct CGRect frameRect = {0, 0, 600, 500};
		frameRect := objc.ID(objc.GetClass("NSRect")).Send(objc.RegisterName("alloc")).
			Send(objc.RegisterName("initWithOriginX"), 0, 0, 600, 500)

		// id window = [[NSWindow alloc] initWithContentRect:frameRect styleMask:NSWindowStyleMaskTitled|NSWindowStyleMaskClosable|NSWindowStyleMaskResizable backing:NSBackingStoreBuffered defer:NO];

		window := objc.ID(objc.GetClass("NSWindow")).Send(objc.RegisterName("alloc"))
		window.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
			frameRect,
			NSWindowStyleMaskTitled|NSWindowStyleMaskClosable|NSWindowStyleMaskResizable,
			NSBackingStoreBuffered,
			false,
		)

		//msg_id(window, sel("setTitle:"), msg_cls_chr(cls("NSString"), sel("stringWithUTF8String:"), "Pure C App"));
		window.Send(objc.RegisterName("setTitle:"), objc.GetClass("NSString"), objc.RegisterName("stringWithUTF8String:"), "PureGo App")

		// [window makeKeyAndOrderFront:nil];
		window.Send(objc.RegisterName("makeKeyAndOrderFront"), "")
		w.window = uintptr(window)
	}

	return w
}

func (w *Window) SetTitle(title string) *Window {
	objc.ID(w.window).Send(objc.RegisterName("setTitle:"),
		objc.GetClass("NSString"), objc.RegisterName("stringWithUTF8String:"), title)
	return w
}

func (w *Window) SetSize(height, width uint) *Window {
	w.width = width
	w.height = height
	return w
}
