package pages

import (
	"fmt"
	"strings"

	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ExecutionClient struct {
	*PageType
	buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
	currentValue        string
}

func (p *ExecutionClient) Page() tview.Primitive {
	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()
	p.currentValue = state.ExecutionClient.SelectedOption

	left, buttons := components.BodyWithOptions(
		`Kindly select the Execution client of your choice.
If you have no strong preference, we
suggest selecting System-recommended for your
convenience.`,
		config.ExecutionClient.Options,
		config.ExecutionClient.OptionLabels,
		p.onSumit,
	)
	p.buttons = buttons

	p.rightSide = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(p.titleTextView, 2, 1, false).
		AddItem(p.descriptionTextView, 1, 1, false).
		AddItem(nil, 0, 1, false)

	p.updateRightSidebar(p.currentValue)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false).
		AddItem(components.VerticalLine(tcell.ColorDarkSlateGray), 1, 1, false).
		AddItem(nil, 2, 1, false).
		AddItem(p.rightSide, 40, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ExecutionClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *ExecutionClient) updateRightSidebar(option string) {
	desc := config.ExecutionClient.Descriptions[option]
	title := config.ExecutionClient.OptionLabels[option]
	p.currentValue = option
	p.titleTextView.SetText(title)
	p.descriptionTextView.SetText(desc)

	if p.rightSide != nil {
		p.rightSide.ResizeItem(p.descriptionTextView, strings.Count(desc, "\n")+1, 1)
	} else {
		log.Error("Update right sidebar: ", "nil")
	}

	p.App.SetFocus(p.buttons[option])
}

func (p *ExecutionClient) onSumit(option string) {
	if option != config.ExecutionClient.Option.SystemRecommended {
		state.ExecutionClient.SelectedOption = option
		ChangePage(config.PageID.ConsensusClientSelection)
		return
	}

	// TODO: fix hardcoded values
	recommendedValue := GetRandomEcClient()
	recommendedLabel := config.ExecutionClient.OptionLabels[recommendedValue]

	alert := components.Alert(
		fmt.Sprintf("%s has been designated as your execution client based on system-recommendation", recommendedLabel),
		[]string{"DONE"},
		map[string]func(){
			"DONE": func() {
				state.ExecutionClient.SelectedOption = recommendedValue
				p.updateRightSidebar(recommendedValue)
				p.App.SetRoot(Pages, true)
				ChangePage(config.PageID.ConsensusClientSelection)
			},
		},
		func() {
			p.App.SetRoot(Pages, true).SetFocus(p.GetFirstElement())
		},
	)

	p.App.SetRoot(alert, true)
}

func (p *ExecutionClient) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.ExecutionClient.Options, p.currentValue)
	log.Infof("Select prev: [%s] to [%s]", p.currentValue, prevItem)
	p.updateRightSidebar(prevItem)
}

func (p *ExecutionClient) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.ExecutionClient.Options, p.currentValue)
	log.Infof("Select next: [%s] to [%s]", p.currentValue, nextItem)
	p.updateRightSidebar(nextItem)
}

func (p *ExecutionClient) GoBack() {
	ChangePage(config.PageID.EthClient)
}

func (p *ExecutionClient) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	shouldHandleEvents := Pages.HasFocus()
	if !shouldHandleEvents {
		return event
	}

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

func (p *ExecutionClient) GetFirstElement() tview.Primitive {
	fb := p.buttons[state.ExecutionClient.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", p.ID, fb.GetLabel())
	return fb
}
