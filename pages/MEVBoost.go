package pages

import (
	"strings"

	"github.com/stader-labs/ethcli-ui/components"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/state"
	"github.com/stader-labs/ethcli-ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MEVBoost struct {
	*PageType
	buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
}

func (p *MEVBoost) Page() tview.Primitive {
	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	left, buttons := components.BodyWithOptions(
		`Stader Node comes with an MEV-Boost, which
enables you to capture additional profits from
block proposals.

Would you prefer Stader to manage your MEV-
Boost, or would you like to handle it yourself?`,
		config.MEVBoost.Options,
		config.MEVBoost.OptionLabels,
		p.onSumit,
	)
	p.buttons = buttons

	p.rightSide = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(p.titleTextView, 2, 1, false).
		AddItem(p.descriptionTextView, 1, 1, false).
		AddItem(nil, 0, 1, false)

	p.updateRightSidebar()

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false).
		AddItem(components.VerticalLine(tcell.ColorDarkSlateGray), 4, 1, false).
		AddItem(p.rightSide, 40, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.MEVBoost), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *MEVBoost) updateRightSidebar() {
	desc := config.MEVBoost.Descriptions[state.MEVBoost.SelectedOption]
	title := config.MEVBoost.OptionLabels[state.MEVBoost.SelectedOption]

	p.titleTextView.SetText(title)
	p.descriptionTextView.SetText(desc)

	if p.rightSide != nil {
		p.rightSide.ResizeItem(p.descriptionTextView, strings.Count(desc, "\n")+1, 1)
	} else {
		log.Error("Update right sidebar: ", "nil")
	}
}

func (p *MEVBoost) onSumit(option string) {
	state.MEVBoost.SelectedOption = option

	if option == config.MEVBoost.Option.LocallyManaged {
		ChangePage(config.PageID.MEVBoostLocal)
	} else {
		ChangePage(config.PageID.MEVBoostExternal)
	}
}

func (p *MEVBoost) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.MEVBoost.Options, state.MEVBoost.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.MEVBoost.SelectedOption, prevItem)
	state.MEVBoost.SelectedOption = prevItem
	p.updateRightSidebar()
	p.App.SetFocus(p.buttons[prevItem])
}

func (p *MEVBoost) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.MEVBoost.Options, state.MEVBoost.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.MEVBoost.SelectedOption, nextItem)
	state.MEVBoost.SelectedOption = nextItem
	p.updateRightSidebar()
	p.App.SetFocus(p.buttons[nextItem])
}

func (p *MEVBoost) GoBack() {
	ChangePage(config.PageID.Monitoring)
}

func (p *MEVBoost) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	if key == tcell.KeyLeft {
		p.selectPrevOption()
	}

	if key == tcell.KeyRight {
		p.selectNextOption()
	}

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (p *MEVBoost) OnResume() {
	log.Info("MEVBoost page changed")
}

func (p *MEVBoost) GetFirstElement() tview.Primitive {
	fb := p.buttons[state.MEVBoost.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", p.ID, fb.GetLabel())
	return fb
}
