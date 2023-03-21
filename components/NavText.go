package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NavText(text string, selected bool) *tview.TextView {
	defaultNavStyle := tcell.
		StyleDefault.
		Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	selectedNavStyle := tcell.StyleDefault.
		Background(tcell.ColorForestGreen).
		Foreground(tcell.ColorBlack)

	navItem := tview.NewTextView().
		SetText(fmt.Sprintf("\n%s ", text)).
		SetSize(3, 0)

	if selected {
		navItem.SetTextStyle(selectedNavStyle)
	} else {
		navItem.SetTextStyle(defaultNavStyle)
	}

	return navItem
}
