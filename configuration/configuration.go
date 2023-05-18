package configuration

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/components"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/logger"
	"github.com/stader-labs/ethcli-ui/configuration/pages"
	"github.com/stader-labs/ethcli-ui/configuration/state"
)

var (
	log            = logger.Log
	previousWidth  int
	previousHeight int
)

func Run(settings *map[string]interface{}) (
	saved bool,
	openWizard bool,
	configurations *map[string]interface{},
) {
	app := tview.NewApplication().EnableMouse(true)
	fieldKeys := config.GetFieldKey()

	if settings != nil {
		state.Configuration = *settings
		state.Saved = false
		state.OpenWizard = false
	}
	// Create the main grid
	grid := tview.NewGrid().
		SetColumns(1, 0, 1).   // 1-unit border
		SetRows(1, 1, 1, 0, 1) // Also 1-unit border

	smallScreenAlert := components.Alert(
		`Greetings! It seems that your terminal may not be large enough to run the service configuration app. To ensure proper functionality, please resize your terminal window to a larger size. This will allow you to view the app as intended and make any necessary configurations`,
		[]string{"Ok (Exit)"},
		map[string]func(){
			"Ok (Exit)": func() {
				app.Stop()
			},
		},
		nil,
	)

	pages.Setup(app)
	rootPageID := config.PageID.Categories
	selectedCat, _ := state.Configuration[fieldKeys.App___selected_category].(string)
	if selectedCat != "" {
		state.Categories.Option.Selected = selectedCat
		rootPageID = config.PageID.ConfigurationForm
	}
	firstElement := pages.All[rootPageID].GetFirstElement()
	pages.ChangePage(rootPageID, app)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pageId, page := pages.Pages.GetFrontPage()
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

		newEvent := event
		if page.HasFocus() {
			newEvent = pages.All[pageId].HandleEvents(event)
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
	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		x, y := screen.Size()

		if x == previousWidth && y == previousHeight {
			return false
		}
		if x < 112 || y < 32 {
			grid.RemoveItem(pages.Pages)
			grid.AddItem(smallScreenAlert, 3, 1, 1, 1, 0, 0, false)
		} else {
			grid.RemoveItem(smallScreenAlert)
			grid.AddItem(pages.Pages, 3, 1, 1, 1, 0, 0, true)
		}

		previousWidth = x
		previousHeight = y
		return false
	})
	if err := app.SetRoot(grid, true).SetFocus(firstElement).Run(); err != nil {
		log.Fatal(err)
	}

	return state.Saved, state.OpenWizard, &state.Configuration
}
