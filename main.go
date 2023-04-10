// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

//go:build darwin || linux

package main

import (
	"fmt"

	"purego-gtk/pkg/gtk"
)

func main() {
	fmt.Println("main.go")

	application := gtk.NewApplication("io.wails")
	//	window := application.NewWindow("Name").
	//		SetSize(100, 100)
	//	window.Show()

	application.Run(0, []string{})
}
