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

type ConsensusClientGraffiti struct {
	*PageType
	firstElement tview.Primitive
}

func (p *ConsensusClientGraffiti) Page() tview.Primitive {
	form := components.Form().
		AddInputField("Add graffiti", state.ConsensusClient.Graffiti, 0,
			func(textToCheck string, lastChar rune) bool {
				return len(textToCheck) <= 16
			},
			func(text string) {
				state.ConsensusClient.Graffiti = text
			},
		).
		AddButton("Next", func() {
			p.onSumit()
		})

	formHeight := 3 + 2

	p.firstElement = form

	bodyText := `Want to add a personal touch to your proposed blocks?
You have the option to add a short custom message, or
"graffiti", of up to 16 characters to each block
proposed by your validators. 

This feature is non-mandatory and is intended purely
for fun! If you do not wish to add any graffiti, simply
leave it blank.`

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

	formWrapHeight := formHeight + bodyTextHeight + 1

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(formWrap, formWrapHeight, 1, false).
		AddItem(nil, 0, 1, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(
			components.PageLeftNav(
				config.ConsensusClient.Stages,
				config.ConsensusClient.Stage.Graffiti.Name,
			),
			40, 1, false,
		).
		AddItem(content, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(Version), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(p.App), 5, 1, false)
}

func (p *ConsensusClientGraffiti) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ConsensusClientGraffiti)
	ChangePage(config.PageID.ConsensusClientCheckpointSync)
}

func (p *ConsensusClientGraffiti) GoBack() {
	nextPage := config.PageID.ConsensusClientSelection

	switch state.ETHClient.SelectedOption {
	case config.ETHClient.Option.LocallyManaged:
		nextPage = config.PageID.ConsensusClientSelection
	case config.ETHClient.Option.ExternallyManaged:
		selectedOption := state.ConsensusClientExternalSelection.SelectedOption
		if selectedOption == config.ConsensusClientExternalSelection.Option.Teku {
			nextPage = config.PageID.ConsensusClientExternalSelectedTeku
		} else if selectedOption == config.ConsensusClientExternalSelection.Option.Lighthouse {
			nextPage = config.PageID.ConsensusClientExternalSelectedLighthouse
		} else if selectedOption == config.ConsensusClientExternalSelection.Option.Prysm {
			nextPage = config.PageID.ConsensusClientExternalSelectedPrysm
		} else if selectedOption == config.ConsensusClientExternalSelection.Option.Nimbus {
			nextPage = config.PageID.ConsensusClientExternalSelectedNimbus
		}
	}

	ChangePage(nextPage)
}

func (p *ConsensusClientGraffiti) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	if key == tcell.KeyEscape {
		p.GoBack()
	}
	return event
}

func (p *ConsensusClientGraffiti) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
