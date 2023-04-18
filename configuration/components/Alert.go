package components

import (
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/logger"
)

func Alert(
	text string,
	buttonLabels []string,
	buttonActions map[string]func(),
	onDismiss func(),
) tview.Primitive {
	modal := tview.NewModal().
		SetText(text).
		AddButtons(buttonLabels).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			logger.Log.Infof("Alert: %s", buttonLabel)
			if buttonIndex == -1 {
				if onDismiss != nil {
					onDismiss()
				}
				return
			}

			action, ok := buttonActions[buttonLabel]
			if ok {
				action()
			}
		})

	return modal
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
		onDismiss,
	)
}
