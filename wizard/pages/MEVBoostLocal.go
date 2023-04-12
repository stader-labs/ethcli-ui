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

type MEVBoostLocal struct {
	*PageType
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
	rightSide           *tview.Flex
	firstElement        tview.Primitive
}

func (p *MEVBoostLocal) Page() tview.Primitive {
	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	form := components.Form().
		AddFormItem(components.Checkbox("Regulated", state.MEVBoostLocal.Regulated, func(checked bool) {
			state.MEVBoostLocal.Regulated = checked
		})).
		AddFormItem(components.Checkbox("Unregulated", state.MEVBoostLocal.Unregulated, func(checked bool) {
			state.MEVBoostLocal.Unregulated = checked
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

	bodyText := `Read the MEV profile description and select the
one you wish to activate. If you're not
interested in using MEV-Boost at this time,
leave both options unchecked.`

	bodyTextHeight := strings.Count(bodyText, "\n") + 1

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			utils.CenterText(bodyText),
			bodyTextHeight, 1, false,
		).
		AddItem(nil, 2, 1, false).
		AddItem(formWrap, 0, 1, false).
		AddItem(nil, 0, 1, false)

	unregulatedTextView := tview.NewTextView().SetText(`Unregulated

Choose this option to activate relays that don't adhere to
any sanctions lists and won't censor transactions.
Unregulated (All MEV Types) permits for all forms of MEV,
including sandwich attacks.

Relays: Ultra Sound and bloXroute Max Profit`)

	unregulatedTextViewHeight := 8

	regulatedTextView := tview.NewTextView().SetText(`Regulated
	
Choose this option to activate relays that adhere to
government regulations such as OFAC sanctions. "Regulated
(All MEV Types)" permits all forms of MEV, including
sandwich attacks.

Relays: Blocknative, Flashbots and Eden Network.`)

	regulatedTextViewHeight := 8

	p.rightSide = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(unregulatedTextView, unregulatedTextViewHeight, 1, false).
		AddItem(nil, 3, 1, false).
		AddItem(regulatedTextView, regulatedTextViewHeight, 1, false).
		AddItem(nil, 0, 1, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(content, 0, 1, false).
		AddItem(components.VerticalLine(tcell.ColorDarkSlateGray), 3, 1, false).
		AddItem(p.rightSide, 60, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.MEVBoost), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *MEVBoostLocal) onSumit() {
	ChangePage(config.PageID.Confirmation)
}

func (p *MEVBoostLocal) GoBack() {
	ChangePage(config.PageID.MEVBoost)
}

func (p *MEVBoostLocal) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (p *MEVBoostLocal) OnResume() {
	log.Info("MEVBoost page changed")
}

func (p *MEVBoostLocal) GetFirstElement() tview.Primitive {
	return p.firstElement
}
