package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Header(text string) *tview.TextView {
	return CenterText(fmt.Sprintf("\n%s", text)).
		SetTextStyle(
			tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
				Foreground(tcell.ColorAntiqueWhite),
		)
}
