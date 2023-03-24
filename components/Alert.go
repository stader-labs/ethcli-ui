package components

import "github.com/rivo/tview"

func Alert(
	text string,
	buttonLabels []string,
	buttonActions map[string]func(),
) tview.Primitive {
	return tview.NewModal().
		SetText(text).
		AddButtons(buttonLabels).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			action, ok := buttonActions[buttonLabel]
			if ok {
				action()
			}
		})
}

func ConfirmAlert(
	text string,
	onConfirm func(),
	onDismiss func(),
) tview.Primitive {
	return Alert(
		text,
		[]string{"Yes", "No"},
		map[string]func(){
			"Yes": onConfirm,
			"No":  onDismiss,
		},
	)
}
