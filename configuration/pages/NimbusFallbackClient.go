package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/components"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/utils"
)

type NimbusFallbackClient struct {
	*PageType
	firstElement tview.Primitive
}

func (p *NimbusFallbackClient) Page() tview.Primitive {

	message := utils.AddNewLines("\nYour ETH2 - Consensus client is Numbus which is not equipped with fallback clients feature.", 98)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 1, 1, false).
		AddItem(tview.NewTextView().SetText(message), 100, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(components.VerticalLine(tcell.ColorGray), 1, 1, false).
		AddItem(nil, 2, 1, false).
		AddItem(nil, 60, 1, false).
		AddItem(nil, 0, 1, false)

	p.firstElement = body

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header("Fallback Clients"), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(
			components.Footer(config.PageID.ConfigurationForm, p.App, p.GoBack),
			5, 1, false,
		)

	return root
}

func (p *NimbusFallbackClient) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	return event
}

func (p *NimbusFallbackClient) GoBack() {
	log.Infof("Going back to [%s]", config.PageID.Categories)
	ChangePage(config.PageID.Categories, p.App)
}
