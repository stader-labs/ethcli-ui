package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Dropdown() *tview.DropDown {
	return tview.NewDropDown().
		SetListStyles(
			tcell.StyleDefault.
				Background(tcell.ColorDarkSlateGray).
				Foreground(tcell.ColorAntiqueWhite),
			tcell.StyleDefault.
				Background(tcell.ColorYellow).
				Foreground(tcell.ColorBlack),
		)
}
