package pages

import (
	"strings"

	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FallbackClientsPrysm struct {
	*PageType
	firstElement tview.Primitive
}

func (p *FallbackClientsPrysm) Page() tview.Primitive {
	form := components.Form().
		AddInputField("Execution client URL", state.FallbackClientsPrysm.ExecutionClientUrl, 0, nil, trimWrap(func(text string) {
			state.FallbackClientsPrysm.ExecutionClientUrl = text
		})).
		AddInputField("Beacon Node HTTP URL", state.FallbackClientsPrysm.BeaconNodeHttpUrl, 0, nil, trimWrap(func(text string) {
			state.FallbackClientsPrysm.BeaconNodeHttpUrl = text
		})).
		AddInputField("Beacon Node JSON-RPC URL", state.FallbackClientsPrysm.BeaconNodeJsonRpcpUrl, 0, nil, trimWrap(func(text string) {
			state.FallbackClientsPrysm.BeaconNodeJsonRpcpUrl = text
		})).
		AddButton("NEXT", func() {
			p.onSumit()
		})

	p.firstElement = form

	formWrap := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(form, 60, 1, false).
		AddItem(nil, 0, 1, false)

	bodyText := `Fallback Clients > Prysm

Please note that you have selected Prysm as your primary Consensus client. To
ensure proper connectivity, it's crucial that your fallback Consensus client is
also running Prysm.

Please enter the URLs of the HTTP APIs for your fallback clients, including the
JSON-RPC URL for your fallback Prysm node.
			
e.g. http://192.168.1.45:8545 for your Execution client and
http://192.168.1.45:5052 for your fallback Prysm node.`

	bodyTextHeight := strings.Count(bodyText, "\n") + 1

	left := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			utils.CenterText(bodyText),
			bodyTextHeight, 1, false,
		).
		AddItem(nil, 2, 1, false).
		AddItem(formWrap, 0, 1, false).
		AddItem(nil, 0, 1, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.FallbackClients), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *FallbackClientsPrysm) onSumit() {
	ChangePage(config.PageID.Monitoring)
}

func (p *FallbackClientsPrysm) GoBack() {
	ChangePage(config.PageID.FallbackClients)
}

func (p *FallbackClientsPrysm) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	if key == tcell.KeyEscape {
		p.GoBack()
	}
	return event
}

func (p *FallbackClientsPrysm) OnResume() {
	log.Info("MEVBoost page changed")
}

func (p *FallbackClientsPrysm) GetFirstElement() tview.Primitive {
	return p.firstElement
}
