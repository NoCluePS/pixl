package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/NoCluePS/pixl/apptype"
	"github.com/NoCluePS/pixl/pxcanvas"
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

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea: fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(20, 20),
		PxRows: 200,
		PxCols: 200,
		PxSize: 20,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		PixlCanvas: pixlCanvas,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
