// package ethcliui

package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/components"
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/logger"
	"github.com/stader-labs/ethcli-ui/pages"
	"github.com/stader-labs/ethcli-ui/state"
	"golang.org/x/term"
)

var (
	log = logger.Log
)

// Run starts the app
// If settings is nil, it will start the app with default settings
func Run(settings *pages.SettingsType) (func() pages.SettingsType, error) {
	state.CurrentApp = tview.NewApplication().EnableMouse(true)
	startPageID := config.PageID.ExecutionClient

	if settings != nil {
		pages.SetSettings(*settings)
	}

	smallScreenAlert := components.Alert(
		`Greetings! It seems that your terminal may not be large enough to run the service configuration app. To ensure proper functionality, please resize your terminal window to a larger size. This will allow you to view the app as intended and make any necessary configurations`,
		[]string{"Ok (Exit)"},
		map[string]func(){
			"Ok (Exit)": func() {
				state.CurrentApp.Stop()
			},
		},
		nil,
	)

	w, h, err := term.GetSize(0)
	if err != nil {
		log.Error("Error getting terminal size")
		return pages.GetSettings, err
	}

	log.Infof("Terminal size is %d x %d", w, h)
	if w < 150 || h < 40 {
		state.CurrentApp.
			SetRoot(smallScreenAlert, true).
			Run()
		return pages.GetSettings, nil
	}

	pages.Setup(state.CurrentApp)
	allPages := pages.Pages

	state.CurrentApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pageName, _ := allPages.GetFrontPage()
		newEvent := pages.All[pageName].HandleEvents(event)
		return newEvent
	})

	pages.ChangePage(startPageID)
	firstElement := pages.All[startPageID].GetFirstElement()

	log.Info("Starting app")
	return pages.GetSettings, state.CurrentApp.
		SetRoot(allPages, true).
		SetFocus(firstElement).Run()
}
