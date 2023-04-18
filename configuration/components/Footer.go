package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/state"
)

func emptyText() tview.Primitive {
	textStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	return CenterText("").
		SetTextStyle(textStyle)
}

func Footer(pageName string, app *tview.Application, onGoBack func()) tview.Primitive {

	textStyle := tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	btnStyle := tcell.StyleDefault.Background(tcell.ColorGray).
		Foreground(tcell.ColorAntiqueWhite)

	footer := tview.NewFlex().
		SetDirection(tview.FlexRow)

	saveNExitTxt := "Save and Exit (Ctrl+S)"
	saveNExitBtn := tview.NewButton(saveNExitTxt).
		SetStyle(btnStyle).
		SetSelectedFunc(func() {
			state.OpenWizard = false
			state.Saved = true
			app.Stop()
		})

	openConfigTxt := "Open the Configuration Manager (Ctrl+U)"
	openConfig := tview.NewButton(openConfigTxt).
		SetStyle(btnStyle).
		SetSelectedFunc(func() {
			state.OpenWizard = true
			state.Saved = false
			app.Stop()
		})

	footerTopRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(emptyText(), 0, 1, false).
		AddItem(saveNExitBtn, len(saveNExitTxt)+2, 1, false).
		AddItem(emptyText(), 3, 1, false).
		AddItem(openConfig, len(openConfigTxt)+2, 1, false).
		AddItem(emptyText(), 0, 1, false)

	footerActions := tview.NewFlex().
		SetDirection(tview.FlexColumn)
	footerActions.AddItem(emptyText(), 0, 1, false)

	{
		txt := "Tab/Arrow keys: Navigate"
		footerActions.AddItem(
			tview.NewTextView().
				SetText(txt).
				SetTextStyle(textStyle),
			len(txt), 1, false,
		).AddItem(emptyText(), 3, 1, false)
	}

	{
		txt := "Enter: Select"
		footerActions.AddItem(
			tview.NewTextView().
				SetText(txt).
				SetTextStyle(textStyle),
			len(txt), 1, false,
		).AddItem(emptyText(), 3, 1, false)
	}

	{
		txt := "Ctrl+C: Quit without Saving"
		btn := tview.NewButton(txt).SetStyle(textStyle).SetDisabledStyle(textStyle)
		btn.SetStyle(textStyle.Underline(true)).SetSelectedFunc(func() {
			state.OpenWizard = false
			state.Saved = false
			app.Stop()
		})
		footerActions.AddItem(btn, len(txt), 1, false)
	}

	if pageName != config.PageID.Categories {
		footerActions.AddItem(emptyText(), 3, 1, false)
		backTxt := "Esc: Go back"
		footerActions.AddItem(
			tview.NewButton(backTxt).
				SetStyle(textStyle.Underline(true)).
				SetSelectedFunc(onGoBack),
			len(backTxt), 1, false,
		)

		footerActions.AddItem(emptyText(), 3, 1, false)
	}

	footerActions.AddItem(emptyText(), 0, 1, false)

	footer.AddItem(emptyText(), 1, 1, false).
		AddItem(footerTopRow, 0, 1, false).
		AddItem(emptyText(), 1, 1, false).
		AddItem(footerActions, 0, 1, false).
		AddItem(emptyText(), 1, 1, false)

	return footer
}
