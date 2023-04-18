package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Form() *tview.Form {
	form := tview.NewForm()

	btnStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	btnSelectedStyle := tcell.StyleDefault.
		Background(tcell.ColorForestGreen).
		Foreground(tcell.ColorBlack)

	form.
		SetButtonStyle(btnStyle).
		SetButtonActivatedStyle(btnSelectedStyle).
		SetFieldBackgroundColor(tcell.ColorDarkSlateGray).
		SetButtonsAlign(tview.AlignCenter)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown, tcell.KeyRight:
			return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
		case tcell.KeyUp, tcell.KeyLeft:
			return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModNone)
		}
		return event
	})

	return form
}
