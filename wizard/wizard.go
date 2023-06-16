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
)

var (
	log = logger.Log
)

var previousWidth int
var previousHeight int

// Run starts the app
// If settings is nil, it will start the app with default settings
func Run(s *pages.SettingsType) (
	confirmed bool,
	openConfiguration bool,
	settings *pages.SettingsType,
) {
	state.CurrentApp = tview.NewApplication().EnableMouse(true)
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

	// Create the main grid
	grid := tview.NewGrid().
		SetColumns(1, 0, 1).   // 1-unit border
		SetRows(1, 1, 1, 0, 1) // Also 1-unit border

	pages.Version = s.Version
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

	startPageID := config.PageID.Network
	if state.StartPageID != "" {
		startPageID = state.StartPageID
	}
	pages.ChangePage(startPageID)
	firstElement := pages.All[startPageID].GetFirstElement()

	log.Info("Starting app")
	// Set up the resize warning
	state.CurrentApp.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		w, h := screen.Size()

		if w == previousWidth && h == previousHeight {
			return false
		}
		if w < 150 || h < 40 {
			grid.RemoveItem(allPages)
			grid.AddItem(smallScreenAlert, 3, 1, 1, 1, 0, 0, false)
		} else {
			grid.RemoveItem(smallScreenAlert)
			grid.AddItem(allPages, 3, 1, 1, 1, 0, 0, true)
		}

		previousWidth = w
		previousHeight = h
		return false
	})
	grid.AddItem(allPages, 3, 1, 1, 1, 0, 0, true)
	appErr := state.CurrentApp.SetRoot(grid, true).SetFocus(firstElement).Run()

	if appErr != nil {
		panic("Error running app")
	}

	ps := pages.GetSettings()
	return state.Confirmed, state.OpenConfigurationUI, &ps
}
