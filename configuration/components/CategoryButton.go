package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CategoryButton(label string) *tview.Button {
	btnStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorAntiqueWhite)

	btnSelectedStyle := tcell.StyleDefault.
		Background(tcell.ColorForestGreen).
		Foreground(tcell.ColorBlack)

	btn := tview.NewButton(label).
		SetStyle(btnStyle).
		SetActivatedStyle(btnSelectedStyle)

	// All this to align the label to the left of the button
	btn.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		labelColor := tcell.ColorWhite
		if btn.HasFocus() && !btn.IsDisabled() {
			labelColor = tcell.ColorBlack
		}

		// Actually aligning the label to the left
		if width > 0 && height > 0 {
			tview.Print(screen, btn.GetLabel(), x+2, y+height/2, width, tview.AlignLeft, labelColor)
		}

		return 0, 0, 0, 0
	})

	return btn
}
