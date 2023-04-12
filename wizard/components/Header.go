package components

import (
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Header() tview.Primitive {
	return utils.CenterText("\nStader Node 0.0.2-alpha Configuration").
		SetTextStyle(
			tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
				Foreground(tcell.ColorAntiqueWhite),
		)
}
