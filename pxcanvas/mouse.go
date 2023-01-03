package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/NoCluePS/pixl/pxcanvas/brush"
)

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.Scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut() {}
func (pxCanvas *PxCanvas) MouseUp() {}