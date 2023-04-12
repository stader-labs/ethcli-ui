package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/components"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/state"
	"github.com/stader-labs/ethcli-ui/configuration/utils"
)

type Categories struct {
	*PageType
	btns            map[string]tview.Primitive
	descriptionView *tview.TextView
}

func (p *Categories) Page() tview.Primitive {
	p.btns = map[string]tview.Primitive{}
	p.descriptionView = tview.NewTextView()

	btns := tview.NewFlex().
		SetDirection(tview.FlexRow)

	for _, option := range config.Categories.Options {
		btn := components.CategoryButton(option)

		p.btns[option] = btn
		btns.AddItem(nil, 1, 1, false).
			AddItem(btn, 3, 1, false)

		func(option string) {
			btn.SetSelectedFunc(func() {
				p.onSubmit(option)
			})
		}(option)
	}

	descriptionContent := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 2, 1, false).
		AddItem(p.descriptionView, 0, 1, false).
		AddItem(nil, 2, 1, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 1, 1, false).
		AddItem(btns, 100, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(components.VerticalLine(tcell.ColorGray), 1, 1, false).
		AddItem(nil, 2, 1, false).
		AddItem(descriptionContent, 60, 1, false).
		AddItem(nil, 0, 1, false)

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header("Configuration Summary"), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(
			components.Footer(config.PageID.Categories, p.App, nil),
			5, 1, false,
		)

	p.updateDescription()

	return root
}

func (p *Categories) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyTab, tcell.KeyDown:
		p.selectNext()
	case tcell.KeyBacktab, tcell.KeyUp:
		p.selectPrevious()
	}
	return event
}

func (p *Categories) updateDescription() {
	desc := config.Categories.Descriptions[state.Categories.Option.Selected]
	p.descriptionView.SetText(desc)
}

func (p *Categories) selectOption(option string) {
	state.Categories.Option.Selected = option
	p.App.SetFocus(p.btns[option])
}

func (p *Categories) onSubmit(option string) {
	p.selectOption(option)
	p.updateDescription()
	log.Info("Selected option: ", option)
	ChangePage(config.PageID.ConfigurationForm, p.App)
}

func (p *Categories) selectNext() {
	option, _ := utils.GetNextItem(config.Categories.Options, state.Categories.Option.Selected)
	p.selectOption(option)
	p.updateDescription()
}

func (p *Categories) selectPrevious() {
	option, _ := utils.GetPrevItem(config.Categories.Options, state.Categories.Option.Selected)
	p.selectOption(option)
	p.updateDescription()
}

func (p *Categories) GetFirstElement() tview.Primitive {
	return p.btns[state.Categories.Option.Selected]
}
