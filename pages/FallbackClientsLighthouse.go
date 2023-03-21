package pages

import (
	"github.com/stader-labs/ethcli-ui/components"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/state"
	"github.com/stader-labs/ethcli-ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FallbackClientsLighthouse struct {
	*PageType
	firstElement tview.Primitive
}

func (p *FallbackClientsLighthouse) Page() tview.Primitive {
	p.PageType.ID = config.PageID.FallbackClientsLighthouse

	form := tview.NewForm().
		AddInputField("Execution client URL", state.FallbackClientsLighthouse.ExecutionClientUrl, 0, nil, func(text string) {
			state.FallbackClientsLighthouse.ExecutionClientUrl = text
		}).
		AddInputField("Beacon Node HTTP URL", state.FallbackClientsLighthouse.BeaconNodeHttpUrl, 0, nil, func(text string) {
			state.FallbackClientsLighthouse.BeaconNodeHttpUrl = text
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
			utils.CenterText(`Fallback Clients > Lighthouse

Your selected Consensus client's validator has the ability to connect with any
Execution client and any Consensus client. 

To ensure proper connectivity, please enter the URLs of the HTTP APIs for your
fallback clients.e.g.
http://192.168.1.45:8545 for your Execution client and
http://192.168.1.45:5052 for your fallback Lighthouse node.`),
			7, 1, false,
		).
		AddItem(nil, 2, 1, false).
		AddItem(formWrap, 0, 1, false).
		AddItem(nil, 0, 1, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.FallbackClients), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *FallbackClientsLighthouse) onSumit() {
	ChangePage(config.PageID.Monitoring)
}

// GoBack
func (p *FallbackClientsLighthouse) GoBack() {
	ChangePage(config.PageID.FallbackClients)
}

func (p *FallbackClientsLighthouse) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (p *FallbackClientsLighthouse) OnResume() {
	log.Info("MEVBoost page changed")
}

func (p *FallbackClientsLighthouse) GetFirstElement() tview.Primitive {
	return p.firstElement
}
