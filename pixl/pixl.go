package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/NoCluePS/pixl/apptype"
	"github.com/NoCluePS/pixl/swatch"
	"github.com/NoCluePS/pixl/ui"
)

func main() {
	pixl := app.New()
	pixlWindow := pixl.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
