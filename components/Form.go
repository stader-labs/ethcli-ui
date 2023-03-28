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

	return form
}
