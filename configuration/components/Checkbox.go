package components

import "github.com/rivo/tview"

func Checkbox() *tview.Checkbox {
	checkbox := tview.NewCheckbox()
	checkbox.SetCheckedString(" ✓ ")
	return checkbox
}
