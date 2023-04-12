package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"
)

func emptyText() tview.Primitive {
	textStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	return utils.CenterText("").SetTextStyle(textStyle)
}

func Footer(App *tview.Application) tview.Primitive {
	textStyle := tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	btnStyle := tcell.StyleDefault.Background(tcell.ColorGray).
		Foreground(tcell.ColorAntiqueWhite)

	saveNExitBtn := tview.NewButton("Save and Exit").
		SetStyle(btnStyle).
		SetSelectedFunc(func() {
			state.Confirmed = true
			state.OpenConfigurationUI = false
			App.Stop()
		})

	openConfig := tview.NewButton("Open the Configuration Manager").
		SetStyle(btnStyle).
		SetSelectedFunc(func() {
			state.Confirmed = false
			state.OpenConfigurationUI = true
			App.Stop()
		})

	footerTopRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(emptyText(), 0, 1, false).
		AddItem(saveNExitBtn, len(" Save and Exit "), 1, false).
		AddItem(emptyText(), 3, 1, false).
		AddItem(openConfig, len(" Open the Configuration Manager "), 1, false).
		AddItem(emptyText(), 0, 1, false)

	footerActions := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(emptyText(), 0, 1, false)

	{
		btn := tview.NewTextView().
			SetText("Esc: Go Back").
			SetTextStyle(textStyle)
		footerActions.AddItem(btn, 12, 1, false)
	}

	{
		btn := tview.NewTextView().
			SetText("Arrow keys: Navigate").
			SetTextStyle(textStyle)
		footerActions.AddItem(emptyText(), 3, 1, false)
		footerActions.AddItem(btn, 20, 1, false)
	}

	{
		btn := tview.NewTextView().
			SetText("Enter: Select").
			SetTextStyle(textStyle)
		footerActions.AddItem(emptyText(), 3, 1, false)
		footerActions.AddItem(btn, 13, 1, false)
	}

	{
		btn := tview.NewTextView().
			SetText("Ctrl+C: Quit without Saving").
			SetTextStyle(textStyle)
		footerActions.AddItem(emptyText(), 3, 1, false)
		footerActions.AddItem(btn, 30, 1, false)
	}

	footerActions.AddItem(emptyText(), 0, 1, false)

	footer := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(emptyText(), 1, 1, false).
		AddItem(footerTopRow, 0, 1, false).
		AddItem(emptyText(), 1, 1, false).
		AddItem(footerActions, 0, 1, false).
		AddItem(emptyText(), 1, 1, false)

	return footer
}
