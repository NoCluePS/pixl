package ui

import (
	"fyne.io/fyne/v2"
	"github.com/NoCluePS/pixl/apptype"
	"github.com/NoCluePS/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}