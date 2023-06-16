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

type FallbackClientsLighthouse struct {
	*PageType
	firstElement tview.Primitive
}

func (p *FallbackClientsLighthouse) Page() tview.Primitive {
	form := components.Form().
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

	bodyText := `Fallback Clients > Lighthouse

Your selected Consensus client's validator has the ability to connect with any
Execution client and any Consensus client. 

To ensure proper connectivity, please enter the URLs of the HTTP APIs for your
fallback clients.e.g.
http://192.168.1.45:8545 for your Execution client and
http://192.168.1.45:5052 for your fallback Lighthouse node.`

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
