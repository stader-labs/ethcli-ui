package wizard

// package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/wizard/components"
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/logger"
	"github.com/stader-labs/ethcli-ui/wizard/pages"
	"github.com/stader-labs/ethcli-ui/wizard/state"
	"golang.org/x/term"
)

var (
	log = logger.Log
)

// Run starts the app
// If settings is nil, it will start the app with default settings
func Run(s *pages.SettingsType) (
	confirmed bool,
	openConfiguration bool,
	settings *pages.SettingsType,
) {
	state.CurrentApp = tview.NewApplication().EnableMouse(true)
	startPageID := config.PageID.Network

	if s != nil {
		pages.SetSettings(*s)
		state.Confirmed = false
		state.OpenConfigurationUI = false
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
		ps := pages.GetSettings()
		return state.Confirmed, state.OpenConfigurationUI, &ps
	}

	log.Infof("Terminal size is %d x %d", w, h)
	if w < 150 || h < 40 {
		state.CurrentApp.
			SetRoot(smallScreenAlert, true).
			Run()
		ps := pages.GetSettings()
		return state.Confirmed, state.OpenConfigurationUI, &ps
	}

	pages.Setup(state.CurrentApp)
	allPages := pages.Pages

	state.CurrentApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()

		if key == tcell.KeyCtrlS {
			state.Confirmed = true
			state.OpenConfigurationUI = false
			state.CurrentApp.Stop()
			return nil
		}

		if key == tcell.KeyCtrlU {
			state.Confirmed = false
			state.OpenConfigurationUI = true
			state.CurrentApp.Stop()
			return nil
		}

		pageName, _ := allPages.GetFrontPage()
		newEvent := pages.All[pageName].HandleEvents(event)
		return newEvent
	})

	pages.ChangePage(startPageID)
	firstElement := pages.All[startPageID].GetFirstElement()

	log.Info("Starting app")
	appErr := state.CurrentApp.SetRoot(allPages, true).SetFocus(firstElement).Run()

	if appErr != nil {
		panic("Error running app")
	}

	ps := pages.GetSettings()
	return state.Confirmed, state.OpenConfigurationUI, &ps
}
