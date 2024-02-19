package pages

import (
	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ConsensusClientDopelgangerProtection struct {
	*PageType
	buttons map[string]*tview.Button
}

func (p *ConsensusClientDopelgangerProtection) Page() tview.Primitive {
	left, buttons := components.BodyWithOptions(
		`Your client comes packed with an awesome feature -
Doppelganger Protection. It safeguards your validators
from being slashed and removed from the Beacon Chain,
just in case you accidentally run your validator keys on
multiple machines simultaneously.

Here's how it works: whenever your validator client
restarts, it will deliberately skip 2-3 attestations for
each validator. If all of them are skipped successfully,
you can be rest assured that you're good to start
attesting again
		
Do you want to enable this awesome feature?`,
		config.ConsensusClient.Stage.DopelgangerProtection.Options,
		config.ConsensusClient.Stage.DopelgangerProtection.OptionLabels,
		p.onSumit,
	)
	p.buttons = buttons

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(components.PageLeftNav(config.ConsensusClient.Stages, config.ConsensusClient.Stage.DopelgangerProtection.Name), 40, 1, false).
		AddItem(left, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *ConsensusClientDopelgangerProtection) onSumit(option string) {
	log.Infof("Selected option: [%s] to [%s]", state.ConsensusClient.DopelgangerProtectionSelectedOption, option)
	state.ConsensusClient.DopelgangerProtectionSelectedOption = option

	ChangePage(config.PageID.FallbackClients)
}

func (p *ConsensusClientDopelgangerProtection) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.ConsensusClient.Stage.DopelgangerProtection.Options, state.ConsensusClient.DopelgangerProtectionSelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.ConsensusClient.DopelgangerProtectionSelectedOption, prevItem)
	state.ConsensusClient.DopelgangerProtectionSelectedOption = prevItem
	p.App.SetFocus(p.buttons[prevItem])
}

func (p *ConsensusClientDopelgangerProtection) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.ConsensusClient.Stage.DopelgangerProtection.Options, state.ConsensusClient.DopelgangerProtectionSelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.ConsensusClient.DopelgangerProtectionSelectedOption, nextItem)
	state.ConsensusClient.DopelgangerProtectionSelectedOption = nextItem
	p.App.SetFocus(p.buttons[nextItem])
}

func (p *ConsensusClientDopelgangerProtection) GoBack() {
	if state.Network.SelectedOption == "mainnet" {
		ChangePage(config.PageID.ConsensusClientCheckpointSyncMainnet)
	} else {
		ChangePage(config.PageID.ConsensusClientCheckpointSyncPrater)
	}
}

func (p *ConsensusClientDopelgangerProtection) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
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

func (n *ConsensusClientDopelgangerProtection) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.ConsensusClient.DopelgangerProtectionSelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
