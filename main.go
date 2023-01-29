package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	makeTray(a)
	mw := a.NewWindow("uwallet")

	mw.Resize(fyne.NewSize(400, 600))
	mw.SetIcon(theme.FyneLogo())

	tabs := makeAppTabsTab()
	mw.SetContent(tabs)

	mw.ShowAndRun()
}

func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}
