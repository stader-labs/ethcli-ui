package components

import (
	"strings"

	"github.com/stader-labs/ethcli-ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Footer() tview.Primitive {
	keyShortcuts := []string{
		"Esc: Go Back",
		"Arrow keys: Navigate",
		"Space/Enter: Select",
		"Ctrl+C: Quit without Saving",
	}
	keyShortcutsText := strings.Join(keyShortcuts, "    ")

	return utils.CenterText("\n" + keyShortcutsText).
		SetTextStyle(
			tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).
				Foreground(tcell.ColorAntiqueWhite),
		)
}
