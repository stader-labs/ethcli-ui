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

type ConsensusClientCheckpointSync struct {
	*PageType
	firstElement tview.Primitive
}

func (p *ConsensusClientCheckpointSync) Page() tview.Primitive {
	p.PageType.ID = config.PageID.ConsensusClientCheckpointSync

	form := tview.NewForm().
		AddInputField("Checkpoint URL", state.ConsensusClient.CheckpointUrl, 0, nil, func(text string) {
			state.ConsensusClient.CheckpointUrl = text
		}).
		AddButton("Next", func() {
			p.onSumit()
		})

	formHeight := 3 + 2

	p.firstElement = form

	bodyText := `Good news - Your client is equipped with the Checkpoint
Sync feature, which can significantly reduce the time and
effort required to sync from scratch. With this powerful
functionality, your client can instantly copy the most
recent state from a trusted Consensus client that you
specify.

If you wish to activate Checkpoint Sync, kindly provide
the provider URL. However, if it's not something you're
interested in, feel free to leave it blank.`

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
	formWrapHeight := formHeight + bodyTextHeight

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
				config.ConsensusClient.Stage.CheckpointSync.Name,
			),
			40, 1, false,
		).
		AddItem(content, 0, 1, false)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(components.Header(), 3, 1, false).
		AddItem(components.Nav(config.TopNav.ConsensusClient), 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(components.Footer(), 3, 1, false)
}

func (p *ConsensusClientCheckpointSync) onSumit() {
	log.Infof("onSumit: [%s]", config.PageID.ConsensusClientCheckpointSync)
	ChangePage(config.PageID.ConsensusClientDopelgangerProtection)
}

func (p *ConsensusClientCheckpointSync) GoBack() {
	ChangePage(config.PageID.ConsensusClientGraffiti)
}

func (p *ConsensusClientCheckpointSync) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	var key = event.Key()

	if key == tcell.KeyEsc {
		p.GoBack()
	}

	return event
}

func (p *ConsensusClientCheckpointSync) GetFirstElement() tview.Primitive {
	fb := p.firstElement
	log.Infof("%s GetFirstElement", p.ID)
	return fb
}
