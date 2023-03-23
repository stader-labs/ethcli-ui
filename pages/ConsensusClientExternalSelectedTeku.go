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

type ConsensusClientExternalSelectedTeku struct {
	*PageType
	firstElement tview.Primitive
	leftSidebar  *tview.Flex
	body         *tview.Flex
}

func (p *ConsensusClientExternalSelectedTeku) Page() tview.Primitive {
	p.PageType.ID = config.PageID.ConsensusClientExternalSelectedTeku

	form := tview.NewForm().
		AddInputField("HTTP URL", "", 0, nil, func(text string) {
			state.ConsensusClientExternalSelectedTeku.HTTPUrl = text
		}).
		AddButton("Next", func() {
			p.onSumit()
		})

	formHeight := 3 + 2

	p.firstElement = form

	bodyText := `Enter the HTTP API URL of your Teku client.
e.g. http://192.168.1.40:5052

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

// func (p *ConsensusClientExternalSelectedTeku) OnResume() {
// 	leftSidebar := components.PageLeftNav(
// 		append([]string{state.ConsensusClientExternalSelection.SelectedOption}, config.ConsensusClient.Stages[1:]...),
// 		state.ConsensusClientExternalSelection.SelectedOption,
// 	)
// 	p.leftSidebar.Clear().AddItem(leftSidebar, 0, 1, false)
// }

func (p *ConsensusClientExternalSelectedTeku) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ConsensusClientExternalSelectedTeku)
	ChangePage(config.PageID.ConsensusClientGraffiti)
}

func (p *ConsensusClientExternalSelectedTeku) GoBack() {
	ChangePage(config.PageID.ExecutionClientExternalSelection)
}

func (p *ConsensusClientExternalSelectedTeku) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	if key == tcell.KeyEsc {
		p.GoBack()
	}

	return event
}

func (p *ConsensusClientExternalSelectedTeku) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
