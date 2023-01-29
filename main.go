package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	mw := a.NewWindow("uwallet")

	mw.Resize(fyne.NewSize(400, 600))
	mw.SetIcon(theme.FyneLogo())

	mw.SetContent(nil)

	mw.ShowAndRun()
}
