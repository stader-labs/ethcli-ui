package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/logger"
	"github.com/stader-labs/ethcli-ui/pages"
	"github.com/stader-labs/ethcli-ui/state"
)

var (
	log = logger.Log
)

// Run starts the app
// If settings is nil, it will start the app with default settings
func Run(settings *pages.SettingsType) (func() pages.SettingsType, error) {
	state.CurrentApp = tview.NewApplication()
	startPageID := config.PageID.Network

	if settings != nil {
		pages.SetSettings(*settings)
	}

	pages.Setup(state.CurrentApp)
	allPages := pages.Pages

	state.CurrentApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pageName, _ := allPages.GetFrontPage()
		newEvent := pages.All[pageName].HandleEvents(event)

		// if newEvent != nil && event.Key() == tcell.KeyEsc {
		// 	log.Debug("Esc pressed")
		// 	pages.GoBack(state.CurrentApp)
		// }

		return newEvent
	})

	pages.ChangePage(startPageID)
	firstElement := pages.All[startPageID].GetFirstElement()

	log.Info("Starting app")
	return pages.GetSettings, state.CurrentApp.
		SetRoot(allPages, true).
		EnableMouse(true).
		SetFocus(firstElement).Run()
}

func main() {
	_, err := Run(nil)
	if err != nil {
		panic(err)
	}
}
