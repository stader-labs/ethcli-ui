package pages

import (
	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Confirmation struct {
	*PageType
	buttons map[string]*tview.Button
}

func (n *Confirmation) Page() tview.Primitive {
	body, buttons := components.BodyWithOptions(
		`Great job! You're all set to complete the
configuration. Take a moment to review and adjust
your Stader Node settings, or simply save and exit
to get started.`,
		config.Confirmation.Options,
		config.Confirmation.OptionLabels,
		n.onSumit,
	)
	n.buttons = buttons

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.Confirmation), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(n.App), 5, 1, false)
}

func (n *Confirmation) onSumit(option string) {
	if option == config.Confirmation.Option.ReviewSettings {
		state.OpenConfigurationUI = true
		state.Confirmed = false
		n.App.Stop()
		return
	}

	if option == config.Confirmation.Option.SaveAndExit {
		log.Info("Save and exit")
		state.OpenConfigurationUI = false
		state.Confirmed = true
		n.App.Stop()
		return
	}
}

func (n *Confirmation) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.Confirmation.Options, state.Confirmation.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.Confirmation.SelectedOption, prevItem)
	state.Confirmation.SelectedOption = prevItem
	n.App.SetFocus(n.buttons[prevItem])
}

func (n *Confirmation) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.Confirmation.Options, state.Confirmation.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.Confirmation.SelectedOption, nextItem)
	state.Confirmation.SelectedOption = nextItem
	n.App.SetFocus(n.buttons[nextItem])
}

func (n *Confirmation) GoBack() {
	nextPage := ""
	pid := config.PageID
	if state.MEVBoost.SelectedOption == config.MEVBoost.Option.ExternallyManaged {
		nextPage = pid.MEVBoostExternal
	} else if state.MEVBoost.SelectedOption == config.MEVBoost.Option.LocallyManaged {
		nextPage = pid.MEVBoostLocal
	} else {
		// This should never happen
		nextPage = pid.MEVBoost
		log.Warnf("This else shouldn't be reached. Current page: %s, next page: %s", n.PageType.ID, nextPage)
	}
	ChangePage(nextPage)
}

func (n *Confirmation) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	if key == tcell.KeyLeft || key == tcell.KeyBacktab {
		n.selectPrevOption()
	}

	if key == tcell.KeyRight || key == tcell.KeyTAB {
		n.selectNextOption()
	}

	if key == tcell.KeyEscape {
		n.GoBack()
	}

	return event
}

func (n *Confirmation) OnResume() {
	log.Info("Confirmation page changed")
}

func (n *Confirmation) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.Confirmation.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
