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

func trimWrap(f func(text string)) func(text string) {
	return func(text string) {
		text = strings.TrimSpace(text)
		f(text)
	}
}

type ExecutionClientExternal struct {
	*PageType
	firstElement tview.Primitive
}

func (p *ExecutionClientExternal) Page() tview.Primitive {
	form := components.Form().
		AddInputField("HTTP-based RPC API", state.ExecutionClientExternal.HTTPBasedRpcApi, 0, nil, trimWrap(func(text string) {
			state.ExecutionClientExternal.HTTPBasedRpcApi = text
		})).
		AddInputField("Websocket-based RPC API", state.ExecutionClientExternal.WebsocketBasedRpcApi, 0, nil, trimWrap(func(text string) {
			state.ExecutionClientExternal.WebsocketBasedRpcApi = text
		})).
		AddButton("Next", func() {
			p.onSumit()
		})

	formHeight := 5 + 2

	p.firstElement = form

	bodyText := `Enter the HTTP-based RPC API URL
and Websocket-based RPC API URL for your
current clients.
E.g. http://127.0.0.1:8545 & ws://127.0.0.1`

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

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(content, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ExecutionClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *ExecutionClientExternal) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ExecutionClientExternal)
	ChangePage(config.PageID.ExecutionClientExternalSelection)
}

func (p *ExecutionClientExternal) GoBack() {
	ChangePage(config.PageID.EthClient)
}

func (p *ExecutionClientExternal) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (p *ExecutionClientExternal) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
