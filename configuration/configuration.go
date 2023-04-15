package configuration

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/logger"
	"github.com/stader-labs/ethcli-ui/configuration/pages"
	"github.com/stader-labs/ethcli-ui/configuration/state"
)

var (
	log = logger.Log
)

func Run(settings *map[string]interface{}) (
	saved bool,
	openWizard bool,
	configurations *map[string]interface{},
) {
	app := tview.NewApplication().EnableMouse(true)

	if settings != nil {
		state.Configuration = *settings
		state.Saved = false
		state.OpenWizard = false
	}

	pages.Setup(app)
	rootPageID := config.PageID.Categories
	firstElement := pages.All[rootPageID].GetFirstElement()
	pages.Pages.SwitchToPage(rootPageID)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pageId, page := pages.Pages.GetFrontPage()
		var newEvent *tcell.EventKey
		key := event.Key()

		if key == tcell.KeyCtrlS {
			state.Saved = true
			state.OpenWizard = false
			app.Stop()
			return nil
		}

		if key == tcell.KeyCtrlU {
			state.Saved = false
			state.OpenWizard = true
			app.Stop()
			return nil
		}

		if page.HasFocus() {
			newEvent = pages.All[pageId].HandleEvents(event)
		} else {
			newEvent = event
		}

		if newEvent == nil {
			return nil
		}

		if page.HasFocus() {
			switch newEvent.Key() {
			case tcell.KeyEscape:
				pages.All[pageId].GoBack()
			}
		}

		return newEvent
	})

	if err := app.SetRoot(pages.Pages, true).SetFocus(firstElement).Run(); err != nil {
		log.Fatal(err)
	}

	return state.Saved, state.OpenWizard, &state.Configuration
}
