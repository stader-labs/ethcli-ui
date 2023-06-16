package components

import (
	"fmt"

	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Header(version string) tview.Primitive {
	return utils.CenterText(fmt.Sprintf("\nStader Node %s", version)).
		SetTextStyle(
			tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
				Foreground(tcell.ColorAntiqueWhite),
		)
}
