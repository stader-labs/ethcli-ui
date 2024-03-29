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

type ConsensusClientExternalSelectedLodestar struct {
	*PageType
	firstElement tview.Primitive
	leftSidebar  *tview.Flex
	body         *tview.Flex
}

func (p *ConsensusClientExternalSelectedLodestar) Page() tview.Primitive {
	form := components.Form().
		AddInputField("HTTP URL", state.ConsensusClientExternalSelectedLodestar.HTTPUrl, 0, nil, trimWrap(func(text string) {
			state.ConsensusClientExternalSelectedLodestar.HTTPUrl = text
		})).
		AddButton("Next", func() {
			p.onSumit()
		})

	formHeight := 5 + 2

	p.firstElement = form

	bodyText := `Enter the HTTP API URL and JSON RPC API URL of your
Lodestar client.
e.g. http://192.168.1.40:5052 for your HTTP API URL and
http://192.168.40.5053 for you JSON RPC API URL

Note: When running this client on the same machine as the
Stader Node, use your machine's LAN IP address instead of
localhost or 127.0.0.1.`

	bodyTextHeight := strings.Count(bodyText, "\n") + 1

	formWrap := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(utils.CenterText(bodyText), bodyTextHeight, 1, false).
				AddItem(nil, 1, 1, false).
				AddItem(form, formHeight, 1, false).
				AddItem(nil, 0, 1, false),
			60, 1, false,
		).
		AddItem(nil, 0, 1, false)

	formWrapHeight := bodyTextHeight + formHeight + 1

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(formWrap, formWrapHeight, 1, false).
		AddItem(nil, 0, 1, false)

	p.leftSidebar = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(components.PageLeftNav(
			append([]string{"Lodestar"}, config.ConsensusClient.Stages[1:]...),
			"Lodestar",
		), 0, 1, false)

	p.body = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(
			p.leftSidebar,
			40, 1, false,
		).
		AddItem(content, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(p.body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *ConsensusClientExternalSelectedLodestar) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ConsensusClientExternalSelectedLodestar)
	ChangePage(config.PageID.ConsensusClientGraffiti)
}

func (p *ConsensusClientExternalSelectedLodestar) GoBack() {
	ChangePage(config.PageID.ExecutionClientExternalSelection)
}

func (p *ConsensusClientExternalSelectedLodestar) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	if key == tcell.KeyEsc {
		p.GoBack()
	}

	return event
}

func (p *ConsensusClientExternalSelectedLodestar) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
