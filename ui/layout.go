package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	pickerContainer := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, pickerContainer, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}