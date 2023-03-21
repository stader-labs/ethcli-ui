package pages

import (
	"github.com/stader-labs/ethcli-ui/components"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/state"
	"github.com/stader-labs/ethcli-ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ExecutionClientExternalSelection struct {
	*PageType
	buttons             map[string]*tview.Button
	titleTextView       *tview.TextView
	descriptionTextView *tview.TextView
}

func (p *ExecutionClientExternalSelection) Page() tview.Primitive {
	cOptions := config.ExecutionClientExternalSelection.Options
	p.PageType.ID = config.PageID.ExecutionClientExternalSelection

	p.titleTextView = tview.NewTextView()
	p.descriptionTextView = tview.NewTextView()

	left, buttons := components.BodyWithOptions(
		`Select the consensus client that you are
managing externally. If your preferred client
is not listed here, it may not be compatible
with the Hybrid mode.

Note: Since each consensus client has its own
unique behavior, Stader Node must be informed
of the specific client you are using
externally. This way, it can adjust its
behavior accordingly.`,
		cOptions,
		p.onSumit,
	)
	p.buttons = buttons

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ExecutionClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *ExecutionClientExternalSelection) onSumit(option string) {
	state.ExecutionClientExternalSelection.SelectedOption = option
	log.Info("Selected option: ", option)

	if state.ExecutionClientExternalSelection.SelectedOption == config.ExecutionClientExternalSelection.Option.Teku {
		ChangePage(config.PageID.ExecutionClientExternalSelectedTeku)
	} else if state.ExecutionClientExternalSelection.SelectedOption == config.ExecutionClientExternalSelection.Option.Lighthouse {
		ChangePage(config.PageID.ExecutionClientExternalSelectedLighthouse)
	} else if state.ExecutionClientExternalSelection.SelectedOption == config.ExecutionClientExternalSelection.Option.Prysm {
		ChangePage(config.PageID.ExecutionClientExternalSelectedPrysm)
	}
}

func (p *ExecutionClientExternalSelection) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.ExecutionClientExternalSelection.Options, state.ExecutionClientExternalSelection.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.ExecutionClientExternalSelection.SelectedOption, prevItem)
	state.ExecutionClientExternalSelection.SelectedOption = prevItem
	p.App.SetFocus(p.buttons[prevItem])
}

func (p *ExecutionClientExternalSelection) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.ExecutionClientExternalSelection.Options, state.ExecutionClientExternalSelection.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.ExecutionClientExternalSelection.SelectedOption, nextItem)
	state.ExecutionClientExternalSelection.SelectedOption = nextItem
	p.App.SetFocus(p.buttons[nextItem])
}

func (p *ExecutionClientExternalSelection) GoBack() {
	ChangePage(config.PageID.ExecutionClientExternal)
}

func (p *ExecutionClientExternalSelection) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	if key == tcell.KeyLeft {
		p.selectPrevOption()
	}

	if key == tcell.KeyRight {
		p.selectNextOption()
	}

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (n *ExecutionClientExternalSelection) GetFirstElement() tview.Primitive {
	fb := n.buttons[state.ExecutionClientExternalSelection.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
