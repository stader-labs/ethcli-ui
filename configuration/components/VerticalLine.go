package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func VerticalLine(color tcell.Color) *tview.Box {
	box := tview.NewBox().
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for i := 0; i < height; i++ {
				screen.SetContent(
					x,
					y+i,
					tview.BoxDrawingsLightVertical,
					nil,
					tcell.StyleDefault.Foreground(color),
				)
			}
			return x, y, width, height
		})

	return box
}
