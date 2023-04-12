package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/config"
)

func emptyText() tview.Primitive {
	textStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	return CenterText("").
		SetTextStyle(textStyle)
}

func Footer(pageName string, onGoBack, onQuit, onSaveExit, onOpenConfig func()) tview.Primitive {

	textStyle := tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	btnStyle := tcell.StyleDefault.Background(tcell.ColorGray).
		Foreground(tcell.ColorAntiqueWhite)

	footer := tview.NewFlex().
		SetDirection(tview.FlexRow)

	footerActions := tview.NewFlex().
		SetDirection(tview.FlexColumn)
	footerActions.AddItem(emptyText(), 0, 1, false)

	{
		txt := "Arrow keys: Navigate"
		if pageName == config.PageID.ConfigurationForm {
			txt = "Tab: Navigate fields"
		}

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
		if onQuit != nil {
			btn.SetStyle(textStyle.Underline(true)).SetSelectedFunc(onQuit)
		} else {
			btn.SetDisabled(true)
		}
		footerActions.AddItem(btn, len(txt), 1, false)
	}

	if pageName == config.PageID.ConfigurationForm {
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

	saveNExitBtn := tview.NewButton("Save and Exit").
		SetStyle(btnStyle).
		SetSelectedFunc(onSaveExit)
	openConfig := tview.NewButton("Open the Configuration Manager").
		SetStyle(btnStyle).
		SetSelectedFunc(onOpenConfig)

	footerTopRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(emptyText(), 0, 1, false).
		AddItem(saveNExitBtn, len(" Save and Exit "), 1, false).
		AddItem(emptyText(), 3, 1, false).
		AddItem(openConfig, len(" Open the Configuration Manager "), 1, false).
		AddItem(emptyText(), 0, 1, false)

	footer.AddItem(emptyText(), 1, 1, false).
		AddItem(footerTopRow, 0, 1, false).
		AddItem(emptyText(), 1, 1, false).
		AddItem(footerActions, 0, 1, false).
		AddItem(emptyText(), 1, 1, false)

	return footer
}
