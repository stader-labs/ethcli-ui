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

type ETHClient struct {
	*PageType
	buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
}

func (p *ETHClient) Page() tview.Primitive {
	cOptions := config.ETHClient.Options
	cDescriptions := config.ETHClient.Descriptions
	p.PageType.ID = config.PageID.EthClient

	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	left, buttons := components.BodyWithOptions(
		`Select your preferred method for managing your
Execution and Consensus clients.`,
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
			strings.Count(cDescriptions[state.ETHClient.SelectedOption], "\n")+1, 1, false,
		).
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
		AddItem(components.Nav(config.TopNav.ETHClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *ETHClient) updateRightSidebar() {
	desc := config.ETHClient.Descriptions[state.ETHClient.SelectedOption]
	p.titleTextView.SetText(state.ETHClient.SelectedOption)
	p.descriptionTextView.SetText(desc)

	if p.rightSide != nil {
		p.rightSide.ResizeItem(p.descriptionTextView, strings.Count(desc, "\n")+1, 1)
	} else {
		log.Error("Update right sidebar: ", "nil")
	}
}

func (p *ETHClient) onSumit(option string) {
	state.ETHClient.SelectedOption = option

	if option == config.ETHClient.Option.LocallyManaged {
		ChangePage(config.PageID.ExecutionClient)
	} else if option == config.ETHClient.Option.ExternallyManaged {
		ChangePage(config.PageID.ExecutionClientExternal)
	}

}

func (p *ETHClient) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.ETHClient.Options, state.ETHClient.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.ETHClient.SelectedOption, prevItem)
	state.ETHClient.SelectedOption = prevItem
	p.updateRightSidebar()
	p.App.SetFocus(p.buttons[prevItem])
}

func (p *ETHClient) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.ETHClient.Options, state.ETHClient.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.ETHClient.SelectedOption, nextItem)
	state.ETHClient.SelectedOption = nextItem
	p.updateRightSidebar()
	p.App.SetFocus(p.buttons[nextItem])
}

func (p *ETHClient) GoBack() {
	ChangePage(config.PageID.Network)
}

func (p *ETHClient) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
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

func (n *ETHClient) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.ETHClient.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
