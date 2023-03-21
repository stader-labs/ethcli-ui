package pages

import (
	"github.com/stader-labs/ethcli-ui/components"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/state"
	"github.com/stader-labs/ethcli-ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Confirmation struct {
	*PageType
	buttons map[string]*tview.Button
}

func (n *Confirmation) Page() tview.Primitive {
	n.PageType.ID = config.PageID.Confirmation
	cOptions := config.Confirmation.Options

	body, buttons := components.BodyWithOptions(
		`Confirmation

Great job! You're all set to complete the
configuration. Take a moment to review and adjust
your Stader Node settings, or simply save and exit
to get started.`,
		cOptions,
		n.onSumit,
	)
	n.buttons = buttons

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.MEVBoost), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (n *Confirmation) onSumit(option string) {
	if option == config.Confirmation.Option.GoBack {
		n.GoBack()
		return
	} else if option == config.Confirmation.Option.SaveAndExit {
		log.Info("Save and exit")
		// saveSettings()
		n.App.Stop()

		if state.OnDone != nil {
			state.OnDone(GetSettings(), nil)
		} else {
			log.Warn("OnDone is nil")
		}
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

func (n *Confirmation) OnResume() {
	log.Info("Confirmation page changed")
}

func (n *Confirmation) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.Confirmation.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
