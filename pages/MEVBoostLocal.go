package pages

import (
	"github.com/stader-labs/ethcli-ui/components"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/state"
	"github.com/stader-labs/ethcli-ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MEVBoostLocal struct {
	*PageType
	// buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
	firstElement        tview.Primitive
}

func (p *MEVBoostLocal) Page() tview.Primitive {
	p.PageType.ID = config.PageID.MEVBoostLocal
	// cOptions := config.MEVBoostLocal.Options

	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	form := tview.NewForm().
		AddCheckbox("Regulated", state.MEVBoostLocal.Regulated, func(checked bool) {
			state.MEVBoostLocal.Regulated = checked
		}).
		AddCheckbox("Unregulated", state.MEVBoostLocal.Unregulated, func(checked bool) {
			state.MEVBoostLocal.Unregulated = checked
		}).
		AddButton("NEXT", func() {
			p.onSumit()
		})

	p.firstElement = form

	formWrap := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(form, 60, 1, false).
		AddItem(nil, 0, 1, false)

	left := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			utils.CenterText(`Read the MEV profile description and select the
one you wish to activate. If you're not
interested in using MEV-Boost at this time,
leave both options unchecked.`),
			4, 1, false,
		).
		AddItem(nil, 2, 1, false).
		AddItem(formWrap, 0, 1, false).
		AddItem(nil, 0, 1, false)

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

func (p *MEVBoostLocal) updateRightSidebar() {
	// desc := config.MEVBoostLocal.Descriptions[state.MEVBoostLocal.SelectedOption]
	// p.titleTextView.SetText(state.MEVBoostLocal.SelectedOption)
	// p.descriptionTextView.SetText(desc)

	// if p.rightSide != nil {
	// 	p.rightSide.ResizeItem(p.descriptionTextView, strings.Count(desc, "\n")+2, 1)
	// } else {
	// 	log.Error("Update right sidebar: ", "nil")
	// }
}

func (p *MEVBoostLocal) onSumit() {
	// state.MEVBoostLocal.SelectedOption = option
	// ChangePage(config.PageID.MEVBoostType)
	ChangePage(config.PageID.Confirmation)
}

func (p *MEVBoostLocal) selectPrevOption() {
	// prevItem, _ := utils.GetPrevItem(config.MEVBoostLocal.Options, state.MEVBoostLocal.SelectedOption)
	// log.Infof("Select prev: [%s] to [%s]", state.MEVBoostLocal.SelectedOption, prevItem)
	// state.MEVBoostLocal.SelectedOption = prevItem
	// p.updateRightSidebar()
	// p.App.SetFocus(p.buttons[prevItem])
}

func (p *MEVBoostLocal) selectNextOption() {
	// nextItem, _ := utils.GetNextItem(config.MEVBoostLocal.Options, state.MEVBoostLocal.SelectedOption)
	// log.Infof("Select next: [%s] to [%s]", state.MEVBoostLocal.SelectedOption, nextItem)
	// state.MEVBoostLocal.SelectedOption = nextItem
	// p.updateRightSidebar()
	// p.App.SetFocus(p.buttons[nextItem])
}

func (p *MEVBoostLocal) GoBack() {
	ChangePage(config.PageID.MEVBoost)
}

func (p *MEVBoostLocal) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
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

func (p *MEVBoostLocal) OnResume() {
	log.Info("MEVBoost page changed")
}

func (p *MEVBoostLocal) GetFirstElement() tview.Primitive {
	// fb := p.buttons[state.MEVBoostLocal.SelectedOption]
	// log.Infof("%s GetFirstElement: [%s]", p.ID, fb.GetLabel())
	return p.firstElement
}
