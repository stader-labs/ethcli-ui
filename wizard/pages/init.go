package pages

import (
	"github.com/stader-labs/ethcli-ui/wizard/config"
	"github.com/stader-labs/ethcli-ui/wizard/logger"

	"github.com/rivo/tview"
)

var (
	All     = map[string]PageInterface{}
	Pages   *tview.Pages
	log     = logger.Log
	Version string
)

// TODO: This is a bit of a mess, I'm sure there's a better way to do this.
// Setup should be called once to initialize the pages
func Setup(app *tview.Application) {
	All[config.PageID.Network] = &Network{PageType: &PageType{}}
	All[config.PageID.EthClient] = &ETHClient{PageType: &PageType{}}
	All[config.PageID.ExecutionClient] = &ExecutionClient{PageType: &PageType{}}
	All[config.PageID.ExecutionClientExternal] = &ExecutionClientExternal{PageType: &PageType{}}
	All[config.PageID.ExecutionClientExternalSelection] = &ExecutionClientExternalSelection{PageType: &PageType{}}
	All[config.PageID.ConsensusClientExternalSelectedLighthouse] = &ConsensusClientExternalSelectedLighthouse{PageType: &PageType{}}
	All[config.PageID.ConsensusClientExternalSelectedNimbus] = &ConsensusClientExternalSelectedNimbus{PageType: &PageType{}}
	All[config.PageID.ConsensusClientExternalSelectedPrysm] = &ConsensusClientExternalSelectedPrysm{PageType: &PageType{}}
	All[config.PageID.ConsensusClientExternalSelectedTeku] = &ConsensusClientExternalSelectedTeku{PageType: &PageType{}}
	All[config.PageID.ConsensusClientExternalSelectedLodestar] = &ConsensusClientExternalSelectedLodestar{PageType: &PageType{}}
	All[config.PageID.ConsensusClientSelection] = &ConsensusClientSelection{PageType: &PageType{}}
	All[config.PageID.ConsensusClientGraffiti] = &ConsensusClientGraffiti{PageType: &PageType{}}
	All[config.PageID.ConsensusClientCheckpointSync] = &ConsensusClientCheckpointSync{PageType: &PageType{}}
	All[config.PageID.ConsensusClientDopelgangerProtection] = &ConsensusClientDopelgangerProtection{PageType: &PageType{}}
	All[config.PageID.FallbackClients] = &FallbackClients{PageType: &PageType{}}
	All[config.PageID.FallbackClientsPrysm] = &FallbackClientsPrysm{PageType: &PageType{}}
	All[config.PageID.FallbackClientsNimbus] = &FallbackClientsNimbus{PageType: &PageType{}}

	All[config.PageID.FallbackClientsLighthouse] = &FallbackClientsLighthouse{PageType: &PageType{}}
	All[config.PageID.FallbackClientsTeku] = &FallbackClientsTeku{PageType: &PageType{}}
	All[config.PageID.FallbackClientsLodestar] = &FallbackClientsLodestar{PageType: &PageType{}}
	All[config.PageID.Monitoring] = &Monitoring{PageType: &PageType{}}
	All[config.PageID.MEVBoost] = &MEVBoost{PageType: &PageType{}}
	All[config.PageID.MEVBoostLocal] = &MEVBoostLocal{PageType: &PageType{}}
	All[config.PageID.MEVBoostExternal] = &MEVBoostExternal{PageType: &PageType{}}
	All[config.PageID.Confirmation] = &Confirmation{PageType: &PageType{}}

	Pages = tview.NewPages()
	for pageID, page := range All {
		page.SetPageID(pageID)
		page.SetApp(app)
		Pages.AddPage(pageID, page.Page(), true, false)
	}
}
