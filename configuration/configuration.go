package configuration

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/components"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/logger"
	"github.com/stader-labs/ethcli-ui/configuration/pages"
	"github.com/stader-labs/ethcli-ui/configuration/state"
	"golang.org/x/term"
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
	fieldKeys := config.GetFieldKey()

	if settings != nil {
		state.Configuration = *settings
		state.Saved = false
		state.OpenWizard = false
	}

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

	w, h, err := term.GetSize(0)
	if err != nil {
		log.Error("Error getting terminal size")
		return state.Saved, state.OpenWizard, &state.Configuration
	}

	log.Infof("Terminal size is %d x %d", w, h)
	if w < 150 || h < 40 {
		app.SetRoot(smallScreenAlert, true).Run()
		return state.Saved, state.OpenWizard, &state.Configuration
	}

	pages.Setup(app)
	rootPageID := config.PageID.Categories
	selectedCat := state.Configuration[fieldKeys.App___selected_category].(string)
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

	if err := app.SetRoot(pages.Pages, true).SetFocus(firstElement).Run(); err != nil {
		log.Fatal(err)
	}

	return state.Saved, state.OpenWizard, &state.Configuration
}
