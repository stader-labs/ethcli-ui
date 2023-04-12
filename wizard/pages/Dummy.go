package pages

import (
	"fmt"

	"github.com/stader-labs/ethcli-ui/wizard/components"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Dummy struct {
	*PageType
}

func (p *Dummy) Page() tview.Primitive {
	pageName := fmt.Sprintf("Page: %s", p.ID)

	body := tview.NewBox().
		SetBorder(true).SetTitle(pageName)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(pageName), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *Dummy) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	return event
}

func (p *Dummy) OnResume() {
	log.Info("Network page changed")
}
