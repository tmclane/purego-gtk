// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

//go:build darwin || linux

package main

import (
	"fmt"
	"time"

	"purego-gtk/pkg/gtk"
)

func main() {
	fmt.Println("main.go")

	application := gtk.NewApplication("io.wails")
	var window *gtk.Window
	go func() {
		time.Sleep(1 * time.Second)
		window = application.NewWindow("Name").SetSize(100, 100)
		window.Show()
	}()
	application.Run(0, []string{})
}
