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

func Footer(App *tview.Application) *tview.Flex {
	textStyle := tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	btnStyle := tcell.StyleDefault.Background(tcell.ColorGray).
		Foreground(tcell.ColorAntiqueWhite)

	saveNExitTxt := "Save and Exit (Ctrl+S)"
	saveNExitBtn := tview.NewButton(saveNExitTxt).
		SetStyle(btnStyle).
		SetSelectedFunc(func() {
			state.Confirmed = true
			state.OpenConfigurationUI = false
			App.Stop()
		})

	openConfigTxt := "Open the Configuration Settings (Ctrl+U)"
	openConfigBtn := tview.NewButton(openConfigTxt).
		SetStyle(btnStyle).
		SetSelectedFunc(func() {
			state.Confirmed = false
			state.OpenConfigurationUI = true
			App.Stop()
		})

	footerTopRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(emptyText(), 0, 1, false).
		AddItem(saveNExitBtn, len(saveNExitTxt)+2, 1, false).
		AddItem(emptyText(), 3, 1, false).
		AddItem(openConfigBtn, len(openConfigTxt)+2, 1, false).
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
		txt := "Tab/Arrow keys: Navigate"
		btn := tview.NewTextView().
			SetText(txt).
			SetTextStyle(textStyle)
		footerActions.AddItem(emptyText(), 3, 1, false)
		footerActions.AddItem(btn, len(txt), 1, false)
	}

	{
		btn := tview.NewTextView().
			SetText("Enter: Select").
			SetTextStyle(textStyle)
		footerActions.AddItem(emptyText(), 3, 1, false)
		footerActions.AddItem(btn, 13, 1, false)
	}

	{
		txt := "Ctrl+C: Quit without Saving"
		btn := tview.NewButton(txt).SetStyle(textStyle)
		btn.SetStyle(textStyle.Underline(true)).SetSelectedFunc(func() {
			state.OpenConfigurationUI = false
			state.Confirmed = false
			App.Stop()
		})
		footerActions.AddItem(emptyText(), 3, 1, false)
		footerActions.AddItem(btn, len(txt), 1, false)
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
