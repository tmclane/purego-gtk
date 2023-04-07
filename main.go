// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

//go:build darwin || linux

package main

import (
	"fmt"
	"runtime"

	"purego-test/pkg/gtk"
)

const gtk3 = "libgtk-3.so"
const gtk4 = "libgtk-4.so"

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/lib/libSystem.B.dylib"
	case "linux":
		return "libc.so.6"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	fmt.Println("main.go")

	application := gtk.NewApplication("io.wails")
	//	window := application.NewWindow("Name").
	//		SetSize(100, 100)
	//	window.Show()

	application.Run(0, []string{})
}
