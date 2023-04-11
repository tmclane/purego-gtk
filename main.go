// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

//go:build darwin || linux

package main

import (
	"time"

	"purego-gtk/pkg/application"
)

func main() {
	app := application.New("PureGo Test")
	window := app.NewWindow("Name").SetSize(500, 500)
	go func() {
		time.Sleep(1 * time.Second)
		window.Show()
	}()
	app.Run(0, []string{})
}
