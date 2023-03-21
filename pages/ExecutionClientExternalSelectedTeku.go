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

type ExecutionClientExternalSelectedTeku struct {
	*PageType
	firstElement tview.Primitive
	leftSidebar  *tview.Flex
	body         *tview.Flex
}

func (p *ExecutionClientExternalSelectedTeku) Page() tview.Primitive {
	p.PageType.ID = config.PageID.ExecutionClientExternalSelectedTeku

	form := tview.NewForm().
		AddInputField("HTTP URL", "", 0, nil, func(text string) {
			state.ExecutionClientExternalSelectedTeku.HTTPUrl = text
		}).
		AddButton("Next", func() {
			p.onSumit()
		})

	p.firstElement = form

	bodyText := `Enter the HTTP API URL of your Teku client.
e.g. http://192.168.1.40:5052

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
			append([]string{"Teku"}, config.ConsensusClient.Stages[1:]...),
			"Teku",
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

// func (p *ExecutionClientExternalSelectedTeku) OnResume() {
// 	leftSidebar := components.PageLeftNav(
// 		append([]string{state.ExecutionClientExternalSelection.SelectedOption}, config.ConsensusClient.Stages[1:]...),
// 		state.ExecutionClientExternalSelection.SelectedOption,
// 	)
// 	p.leftSidebar.Clear().AddItem(leftSidebar, 0, 1, false)
// }

func (p *ExecutionClientExternalSelectedTeku) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ExecutionClientExternalSelectedTeku)
	ChangePage(config.PageID.ConsensusClientGraffiti)
}

func (p *ExecutionClientExternalSelectedTeku) GoBack() {
	ChangePage(config.PageID.ExecutionClientExternalSelection)
}

func (p *ExecutionClientExternalSelectedTeku) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	if key == tcell.KeyEsc {
		p.GoBack()
	}

	return event
}

func (p *ExecutionClientExternalSelectedTeku) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
