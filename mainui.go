package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeAppTabsTab() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), makeHome()),
		container.NewTabItemWithIcon("Account", theme.AccountIcon(), widget.NewLabel("Content of account")),
		container.NewTabItemWithIcon("Setting", theme.SettingsIcon(), widget.NewLabel("Content of setting")),
		container.NewTabItemWithIcon("Transaction", theme.DocumentIcon(), widget.NewLabel("Content of transactions")),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	return container.NewBorder(nil, nil, nil, nil, tabs)
}
