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

type ConsensusClientSelection struct {
	*PageType
	buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
	currentValue        string
}

func (p *ConsensusClientSelection) Page() tview.Primitive {
	cOptions := config.ConsensusClient.Stage.Selection.Options
	cDescriptions := config.ConsensusClient.Stage.Selection.Descriptions
	p.PageType.ID = config.PageID.ConsensusClientSelection

	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	left, buttons := components.BodyWithOptions(
		`Select the consensus client you
wish to use. If you're
uncertain, we suggest selecting
System-recommended for your
convenience.`,
		cOptions,
		p.onSumit,
	)
	p.buttons = buttons

	p.rightSide = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(p.titleTextView, 2, 1, false).
		AddItem(
			p.descriptionTextView,
			strings.Count(cDescriptions[state.ConsensusClient.SelectionSelectedOption], "\n"), 1, false,
		).
		AddItem(nil, 0, 1, false)

	p.updateRightSidebar(state.ConsensusClient.SelectionSelectedOption)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(components.PageLeftNav(config.ConsensusClient.Stages, config.ConsensusClient.Stage.Selection.Name), 40, 1, false).
		AddItem(left, 0, 1, false).
		AddItem(components.VerticalLine(tcell.ColorDarkSlateGray), 1, 1, false).
		AddItem(nil, 2, 1, false).
		AddItem(p.rightSide, 41, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *ConsensusClientSelection) updateRightSidebar(option string) {
	desc := config.ConsensusClient.Stage.Selection.Descriptions[option]
	p.currentValue = option
	p.titleTextView.SetText(option)
	p.descriptionTextView.SetText(desc)

	if p.rightSide != nil {
		p.rightSide.ResizeItem(p.descriptionTextView, strings.Count(desc, "\n")+1, 1)
	} else {
		log.Error("Update right sidebar: ", "nil")
	}

	state.ConsensusClient.SelectionSelectedOption = option
	p.App.SetFocus(p.buttons[option])
}

func (p *ConsensusClientSelection) onSumit(option string) {
	log.Infof("Selected option: [%s] to [%s]", state.ConsensusClient.SelectionSelectedOption, option)

	if option != config.ConsensusClient.Stage.Selection.Option.SystemRecommended {
		state.ConsensusClient.SelectionSelectedOption = option
	} else {
		state.ConsensusClient.SelectionSelectedOption = utils.GetRandomItem(
			config.ConsensusClient.Stage.Selection.Options,
			option,
		)
	}
	ChangePage(config.PageID.ConsensusClientGraffiti)
}

func (p *ConsensusClientSelection) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.ConsensusClient.Stage.Selection.Options, p.currentValue)
	log.Infof("Select prev: [%s] to [%s]", p.currentValue, prevItem)
	p.updateRightSidebar(prevItem)
}

func (p *ConsensusClientSelection) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.ConsensusClient.Stage.Selection.Options, p.currentValue)
	log.Infof("Select next: [%s] to [%s]", p.currentValue, nextItem)
	p.updateRightSidebar(nextItem)
}

func (p *ConsensusClientSelection) GoBack() {
	ChangePage(config.PageID.ExecutionClient)
}

func (p *ConsensusClientSelection) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
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

func (n *ConsensusClientSelection) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.ConsensusClient.SelectionSelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
