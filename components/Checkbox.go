package components

import "github.com/rivo/tview"

func Checkbox(label string, checked bool, changed func(bool)) tview.FormItem {
	checkbox := tview.NewCheckbox()
	checkbox.SetLabel(label)
	checkbox.SetChecked(checked)
	checkbox.SetCheckedString(" ✓ ")
	checkbox.SetChangedFunc(changed)
	return checkbox
}
