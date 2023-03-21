package utils

import (
	"github.com/rivo/tview"
)

func CenterText(text string) *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}
