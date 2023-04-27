package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/wizard/logger"
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

	modal.SetBackgroundColor(tcell.ColorAntiqueWhite)
	modal.SetTextColor(tcell.ColorBlack)
	modal.SetButtonBackgroundColor(tcell.ColorBlack)
	modal.SetButtonTextColor(tcell.ColorForestGreen)
	modal.Frame.SetBorderColor(tcell.ColorBlack)
	// modal.Frame.SetTitle(" Alert ")
	// modal.Frame.SetTitleColor(tcell.ColorBlack)

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
