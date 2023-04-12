package components

import (
	"strings"

	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BodyWithOptions(
	instructions string,
	options []string,
	optionLabels map[string]string,
	onSelect func(option string),
) (body *tview.Flex, buttons map[string]*tview.Button) {
	buttons = make(map[string]*tview.Button)

	btnStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkSlateGray).
		Foreground(tcell.ColorAntiqueWhite)

	btnSelectedStyle := tcell.StyleDefault.
		Background(tcell.ColorForestGreen).
		Foreground(tcell.ColorBlack)

	chunks := utils.ChunkSlice(options, 2)

	body = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			utils.CenterText(instructions),
			strings.Count(instructions, "\n")+2, 1, false,
		).
		AddItem(nil, 1, 1, false)

	for ci, chunk := range chunks {
		row := tview.NewFlex().
			AddItem(nil, 0, 1, false)

		for i, v := range chunk {
			btn := func(v string) *tview.Button {
				return tview.NewButton(optionLabels[v]).
					SetSelectedFunc(func() {
						onSelect(v)
					}).
					SetStyle(btnStyle).
					SetActivatedStyle(btnSelectedStyle)
			}(v)

			if i > 0 {
				row.AddItem(nil, 2, 1, false)
			}

			row.AddItem(btn, 30, 1, false)
			buttons[v] = btn
		}

		row.AddItem(nil, 0, 1, false)

		if ci > 0 {
			body.AddItem(nil, 1, 1, false)
		}

		body.AddItem(row, 3, 1, false)
	}

	body.AddItem(nil, 0, 1, false)

	return body, buttons
}
