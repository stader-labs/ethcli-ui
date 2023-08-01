package pages

import (
	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FallbackClients struct {
	*PageType
	buttons map[string]*tview.Button
}

func (p *FallbackClients) Page() tview.Primitive {
	body, buttons := components.BodyWithOptions(
		`Your Stader Node is equipped with fallback clients feature!

In case your primary Execution and Consensus clients go offline for any reason,
you can use an externally-managed client pair that you trust as a backup option.
This ensures that your Validator Client and Stader Node stay connected and 
ontinue functioning properly.
		
So, do you want to set up a fallback client pair?`,
		config.FallbackClients.Options,
		config.FallbackClients.OptionLabels,
		p.onSumit,
	)
	p.buttons = buttons

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.FallbackClients), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *FallbackClients) onSumit(option string) {
	log.Infof("Selected option: [%s] to [%s]", state.FallbackClients.SelectedOption, option)
	state.FallbackClients.SelectedOption = option

	if state.FallbackClients.SelectedOption == config.FallbackClients.Option.Yes {

		switch state.ETHClient.SelectedOption {
		case config.ETHClient.Option.LocallyManaged:
			if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Prysm {
				ChangePage(config.PageID.FallbackClientsPrysm)
			} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.LightHouse {
				ChangePage(config.PageID.FallbackClientsLighthouse)
			} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Teku {
				ChangePage(config.PageID.FallbackClientsTeku)
			} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Nimbus {
				ChangePage(config.PageID.FallbackClientsNimbus)
			} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Lodestar {
				ChangePage(config.PageID.FallbackClientsLodestar)
			} else {
				// Skip for all others
				ChangePage(config.PageID.Monitoring)
			}
		case config.ETHClient.Option.ExternallyManaged:
			if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Teku {
				ChangePage(config.PageID.FallbackClientsTeku)
			} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Lighthouse {
				ChangePage(config.PageID.FallbackClientsLighthouse)
			} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Prysm {
				ChangePage(config.PageID.FallbackClientsPrysm)
			} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClient.Stage.Selection.Option.Nimbus {
				ChangePage(config.PageID.FallbackClientsNimbus)
			} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClient.Stage.Selection.Option.Lodestar {
				ChangePage(config.PageID.FallbackClientsLodestar)
			} else {
				// Skip for all others
				ChangePage(config.PageID.Monitoring)
			}
		}

	} else {
		ChangePage(config.PageID.Monitoring)
	}
}

func (p *FallbackClients) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.FallbackClients.Options, state.FallbackClients.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.FallbackClients.SelectedOption, prevItem)
	state.FallbackClients.SelectedOption = prevItem
	p.App.SetFocus(p.buttons[prevItem])
}

func (p *FallbackClients) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.FallbackClients.Options, state.FallbackClients.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.FallbackClients.SelectedOption, nextItem)
	state.FallbackClients.SelectedOption = nextItem
	p.App.SetFocus(p.buttons[nextItem])
}

func (p *FallbackClients) GoBack() {
	ChangePage(config.PageID.ConsensusClientDopelgangerProtection)
}

func (p *FallbackClients) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
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

func (n *FallbackClients) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.FallbackClients.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
