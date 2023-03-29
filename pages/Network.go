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

type Network struct {
	*PageType
	buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
}

func (n *Network) Page() tview.Primitive {
	cDescriptions := config.Network.Descriptions

	n.titleTextView = tview.NewTextView()
	n.descriptionTextView = tview.NewTextView()

	left, buttons := components.BodyWithOptions(
		"Please select the network where you want to set\nup your node.",
		config.Network.Options,
		config.Network.OptionLabels,
		n.onSumit,
	)
	n.buttons = buttons

	n.rightSide = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(n.titleTextView, 2, 1, false).
		AddItem(
			n.descriptionTextView,
			strings.Count(cDescriptions[state.Network.SelectedOption], "\n")+1, 1, false,
		).
		AddItem(nil, 0, 1, false)

	n.updateRightSidebar()

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false).
		AddItem(components.VerticalLine(tcell.ColorDarkSlateGray), 4, 1, false).
		AddItem(n.rightSide, 40, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.Network), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (n *Network) updateRightSidebar() {
	desc := config.Network.Descriptions[state.Network.SelectedOption]
	title := config.Network.OptionLabels[state.Network.SelectedOption]

	n.titleTextView.SetText(title)
	n.descriptionTextView.SetText(desc)

	if n.rightSide != nil {
		n.rightSide.ResizeItem(n.descriptionTextView, strings.Count(desc, "\n")+1, 1)
	} else {
		log.Error("Update right sidebar: ", "nil")
	}
}

func (n *Network) onSumit(option string) {
	state.Network.SelectedOption = option
	ChangePage(config.PageID.EthClient)
}

func (n *Network) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.Network.Options, state.Network.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.Network.SelectedOption, prevItem)
	state.Network.SelectedOption = prevItem
	n.updateRightSidebar()
	n.App.SetFocus(n.buttons[prevItem])
}

func (n *Network) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.Network.Options, state.Network.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.Network.SelectedOption, nextItem)
	state.Network.SelectedOption = nextItem
	n.updateRightSidebar()
	n.App.SetFocus(n.buttons[nextItem])
}

func (n *Network) GoBack() {
	// main.go will handle exit app confirm dialog
}

func (n *Network) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	if key == tcell.KeyLeft {
		n.selectPrevOption()
	}

	if key == tcell.KeyRight {
		n.selectNextOption()
	}

	if key == tcell.KeyEscape {
		n.GoBack()
	}

	return event
}

func (n *Network) OnResume() {
	log.Info("Network page changed")
}

func (n *Network) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.Network.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
