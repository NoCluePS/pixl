package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) Scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize++
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize--
		}
	default:
		pxCanvas.PxSize = 10
	}
}

func (pxCanvas *PxCanvas) Pan(prevCoord, currCoord fyne.PointEvent) {
	xDif := currCoord.Position.X - prevCoord.Position.X
	yDif := currCoord.Position.Y - prevCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDif
	pxCanvas.CanvasOffset.Y += yDif
	pxCanvas.Refresh()
}