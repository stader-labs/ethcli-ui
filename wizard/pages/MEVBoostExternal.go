package pages

import (
	"strings"

	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MEVBoostExternal struct {
	*PageType
	// buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	// rightSide           *tview.Flex
	firstElement tview.Primitive
}

func (p *MEVBoostExternal) Page() tview.Primitive {
	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	form := components.Form().
		AddInputField("MEV URL", state.MEVBoostExternal.MevUrl, 0, nil, trimWrap(func(text string) {
			state.MEVBoostExternal.MevUrl = text
		})).
		AddButton("NEXT", func() {
			p.onSumit()
		})

	formHeight := 5

	p.firstElement = form

	formWrap := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(form, 60, 1, false).
		AddItem(nil, 0, 1, false)

	bodyText := `Enter the externally managed MEV Boost client URL
	e.g. http://192.168.1.46:18550`

	bodyTextHeight := strings.Count(bodyText, "\n") + 1

	body := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			utils.CenterText(bodyText),
			bodyTextHeight, 1, false,
		).
		AddItem(nil, 2, 1, false).
		AddItem(formWrap, formHeight, 1, false).
		AddItem(nil, 0, 1, false)

	// 	left, buttons := components.BodyWithOptions(
	// 		`MEV Boost > Externally Managed

	// Enter the externally managed MEV Boost client URL
	// e.g. http://192.168.1.46:18550`,
	// 		cOptions,
	// 		p.onSumit,
	// 	)
	// 	p.buttons = buttons

	// p.rightSide = tview.NewFlex().
	// 	SetDirection(tview.FlexRow).
	// 	AddItem(nil, 0, 1, false).
	// 	AddItem(p.titleTextView, 2, 1, false).
	// 	AddItem(p.descriptionTextView, 1, 1, false).
	// 	AddItem(nil, 0, 1, false)

	// p.updateRightSidebar()

	// body := tview.NewFlex().
	// 	SetDirection(tview.FlexColumn).
	// 	AddItem(left, 0, 1, false).
	// 	AddItem(components.VerticalLine(tcell.ColorDarkSlateGray), 4, 1, false).
	// 	AddItem(p.rightSide, 40, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.MEVBoost), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

// func (p *MEVBoostExternal) updateRightSidebar() {
// 	// desc := config.MEVBoostExternal.Descriptions[state.MEVBoostExternal.SelectedOption]
// 	// p.titleTextView.SetText(state.MEVBoostExternal.SelectedOption)
// 	// p.descriptionTextView.SetText(desc)

// 	// if p.rightSide != nil {
// 	// 	p.rightSide.ResizeItem(p.descriptionTextView, strings.Count(desc, "\n")+2, 1)
// 	// } else {
// 	// 	log.Error("Update right sidebar: ", "nil")
// 	// }
// }

func (p *MEVBoostExternal) onSumit() {
	// state.MEVBoostExternal.SelectedOption = option
	// ChangePage(config.PageID.MEVBoostType)
	ChangePage(config.PageID.Confirmation)
}

// func (p *MEVBoostExternal) selectPrevOption() {
// 	// prevItem, _ := utils.GetPrevItem(config.MEVBoostExternal.Options, state.MEVBoostExternal.SelectedOption)
// 	// log.Infof("Select prev: [%s] to [%s]", state.MEVBoostExternal.SelectedOption, prevItem)
// 	// state.MEVBoostExternal.SelectedOption = prevItem
// 	// p.updateRightSidebar()
// 	// p.App.SetFocus(p.buttons[prevItem])
// }

// func (p *MEVBoostExternal) selectNextOption() {
// 	// nextItem, _ := utils.GetNextItem(config.MEVBoostExternal.Options, state.MEVBoostExternal.SelectedOption)
// 	// log.Infof("Select next: [%s] to [%s]", state.MEVBoostExternal.SelectedOption, nextItem)
// 	// state.MEVBoostExternal.SelectedOption = nextItem
// 	// p.updateRightSidebar()
// 	// p.App.SetFocus(p.buttons[nextItem])
// }

func (p *MEVBoostExternal) GoBack() {
	ChangePage(config.PageID.MEVBoost)
}

func (p *MEVBoostExternal) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	// if key == tcell.KeyLeft {
	// 	p.selectPrevOption()
	// }

	// if key == tcell.KeyRight {
	// 	p.selectNextOption()
	// }

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (p *MEVBoostExternal) OnResume() {
	log.Info("MEVBoost page changed")
}

func (p *MEVBoostExternal) GetFirstElement() tview.Primitive {
	// fb := p.buttons[state.MEVBoostExternal.SelectedOption]
	// log.Infof("%s GetFirstElement: [%s]", p.ID, fb.GetLabel())
	return p.firstElement
}
