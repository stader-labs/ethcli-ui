package components

import (
	"strings"

	"github.com/rivo/tview"
)

func DescriptionSidbar(
	description string,
) (func(string), tview.Primitive) {
	textView := tview.NewTextView().
		SetText(description)

	textViewHeight := strings.Count(description, "\n") + 1

	body := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(textView, 0, textViewHeight, true)

	return func(desc string) {
		textView.SetText(desc)
		textViewHeight := strings.Count(desc, "\n") + 1
		body.ResizeItem(textView, textViewHeight, 1)
	}, body
}
