package pages

import (
	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Monitoring struct {
	*PageType
	buttons map[string]*tview.Button
}

func (n *Monitoring) Page() tview.Primitive {
	body, buttons := components.BodyWithOptions(
		`Monitoring system will provide valuable insights into your hardware stats, such
as CPU usage, RAM usage, and free disk space. Additionally, it will also monitor
your validator stats and ETH rewards, among other things. 

Please note that none of this information will not be transmitted to any distant
server for examination - purely for your consumption on your node.
		
Are you interested in enabling Monitoring system?`,
		config.Monitoring.Options,
		config.Monitoring.OptionLabels,
		n.onSumit,
	)
	n.buttons = buttons

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.Monitoring), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(n.App), 5, 1, false)
}

func (n *Monitoring) onSumit(option string) {
	state.Monitoring.SelectedOption = option
	ChangePage(config.PageID.MEVBoost)
}

func (n *Monitoring) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.Monitoring.Options, state.Monitoring.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.Monitoring.SelectedOption, prevItem)
	state.Monitoring.SelectedOption = prevItem
	n.App.SetFocus(n.buttons[prevItem])
}

func (n *Monitoring) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.Monitoring.Options, state.Monitoring.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.Monitoring.SelectedOption, nextItem)
	state.Monitoring.SelectedOption = nextItem
	n.App.SetFocus(n.buttons[nextItem])
}

func (n *Monitoring) GoBack() {
	nextPage := ""
	pid := config.PageID
	if state.FallbackClients.SelectedOption == config.FallbackClients.Option.Yes {
		if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Prysm {
			nextPage = pid.FallbackClientsPrysm
		} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.LightHouse {
			nextPage = pid.FallbackClientsLighthouse
		} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Teku {
			nextPage = pid.FallbackClientsTeku
		} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Nimbus {
			nextPage = pid.ConsensusClientDopelgangerProtection
		} else {
			nextPage = pid.FallbackClients
		}
	} else {
		nextPage = pid.FallbackClients
	}
	ChangePage(nextPage)
}

func (n *Monitoring) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
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

func (n *Monitoring) OnResume() {
	log.Info("Monitoring page changed")
}

func (n *Monitoring) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.Monitoring.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
