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

type ExecutionClientExternalSelectedPrysm struct {
	*PageType
	firstElement tview.Primitive
	leftSidebar  *tview.Flex
	body         *tview.Flex
}

func (p *ExecutionClientExternalSelectedPrysm) Page() tview.Primitive {
	p.PageType.ID = config.PageID.ExecutionClientExternalSelectedPrysm

	form := tview.NewForm().
		AddInputField("HTTP URL", "", 0, nil, func(text string) {
			state.ExecutionClientExternalSelectedPrysm.HTTPUrl = text
		}).
		AddInputField("JSON-RPC URL", "", 0, nil, func(text string) {
			state.ExecutionClientExternalSelectedPrysm.JSONRpcUrl = text
		}).
		AddButton("Next", func() {
			p.onSumit()
		})

	p.firstElement = form

	bodyText := `Enter the HTTP API URL and JSON RPC API URL of your
Prysm client.
e.g. http://192.168.1.40:5052 for your HTTP API URL and
http://192.168.40.5053 for you JSON RPC API URL

Note: When running this client on the same machine as the
Stader Node, use your machine's LAN IP address instead of
localhost or 127.0.0.1.`

	formWrap := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(utils.CenterText(bodyText), strings.Count(bodyText, "\n")+1, 1, false).
				AddItem(nil, 1, 1, false).
				AddItem(form, 7, 1, false).
				AddItem(nil, 0, 1, false),
			60, 1, false,
		).
		AddItem(nil, 0, 1, false)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(formWrap, 10, 1, false).
		AddItem(nil, 0, 1, false)

	p.leftSidebar = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(components.PageLeftNav(
			append([]string{"Prysm"}, config.ConsensusClient.Stages[1:]...),
			"Prysm",
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
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(p.body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *ExecutionClientExternalSelectedPrysm) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ExecutionClientExternalSelectedPrysm)
	ChangePage(config.PageID.ConsensusClientGraffiti)
}

func (p *ExecutionClientExternalSelectedPrysm) GoBack() {
	ChangePage(config.PageID.ExecutionClientExternalSelection)
}

func (p *ExecutionClientExternalSelectedPrysm) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	if key == tcell.KeyEsc {
		p.GoBack()
	}

	return event
}

func (p *ExecutionClientExternalSelectedPrysm) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
