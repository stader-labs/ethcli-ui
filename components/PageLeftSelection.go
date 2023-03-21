package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func PageLeftNav(options []string, selectedNav string) *tview.Flex {
	leftNav := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 1, 1, false)

	for i, stage := range options {
		leftNav.AddItem(NavText(fmt.Sprintf("  %d %s", i+1, stage), stage == selectedNav), 3, 1, false)
		if i != len(options)-1 {
			leftNav.AddItem(nil, 1, 1, false)
		}
	}

	leftNav.AddItem(nil, 0, 1, false)

	leftNavWrap := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 1, 1, false).
		AddItem(leftNav, 0, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(VerticalLine(tcell.ColorDarkSlateGray), 2, 1, false)

	return leftNavWrap
}
