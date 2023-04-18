package pages

import (
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/logger"

	"github.com/rivo/tview"
)

var (
	All   = make(map[string]PageInterface)
	Pages *tview.Pages
	log   = logger.Log
)

// Setup should be called once to initialize the pages
func Setup(app *tview.Application) {
	All[config.PageID.Categories] = &Categories{PageType: &PageType{}}
	All[config.PageID.ConfigurationForm] = &ConfigurationForm{PageType: &PageType{}}
	All[config.PageID.NimbusFallbackClient] = &NimbusFallbackClient{PageType: &PageType{}}

	Pages = tview.NewPages()
	for pageID, page := range All {
		page.SetApp(app)
		page.SetPageID(pageID)
		Pages.AddPage(pageID, page.Page(), true, false)
	}
}
