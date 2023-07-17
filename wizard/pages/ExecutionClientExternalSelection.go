package pages

import (
	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"github.com/stader-labs/ethcli-ui/wizard/utils"

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
		config.ConsensusClientExternalSelection.Options,
		config.ConsensusClientExternalSelection.OptionLabels,
		p.onSumit,
	)
	p.buttons = buttons

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *ExecutionClientExternalSelection) onSumit(option string) {
	state.ConsensusClientExternalSelection.SelectedOption = option
	log.Info("Selected option: ", option)

	if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Teku {
		ChangePage(config.PageID.ConsensusClientExternalSelectedTeku)
	} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Lighthouse {
		ChangePage(config.PageID.ConsensusClientExternalSelectedLighthouse)
	} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Prysm {
		ChangePage(config.PageID.ConsensusClientExternalSelectedPrysm)
	} else if state.ConsensusClientExternalSelection.SelectedOption == config.ConsensusClientExternalSelection.Option.Lodestar {
		ChangePage(config.PageID.ConsensusClientExternalSelectedLodestar)
	}
}

func (p *ExecutionClientExternalSelection) selectPrevOption() {
	prevItem, _ := utils.GetPrevItem(config.ConsensusClientExternalSelection.Options, state.ConsensusClientExternalSelection.SelectedOption)
	log.Infof("Select prev: [%s] to [%s]", state.ConsensusClientExternalSelection.SelectedOption, prevItem)
	state.ConsensusClientExternalSelection.SelectedOption = prevItem
	p.App.SetFocus(p.buttons[prevItem])
}

func (p *ExecutionClientExternalSelection) selectNextOption() {
	nextItem, _ := utils.GetNextItem(config.ConsensusClientExternalSelection.Options, state.ConsensusClientExternalSelection.SelectedOption)
	log.Infof("Select next: [%s] to [%s]", state.ConsensusClientExternalSelection.SelectedOption, nextItem)
	state.ConsensusClientExternalSelection.SelectedOption = nextItem
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
	fb := n.buttons[state.ConsensusClientExternalSelection.SelectedOption]
	log.Infof("%s GetFirstElement: [%s]", n.ID, fb.GetLabel())
	return fb
}
