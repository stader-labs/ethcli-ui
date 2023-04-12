package components

import (
	"github.com/rivo/tview"
)

func CenterText(text string) *tview.TextView {
	view := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text)
	return view
}
