package pages

import (
	"github.com/stader-labs/ethcli-ui/config"
	"github.com/stader-labs/ethcli-ui/logger"

	"github.com/rivo/tview"
)

var (
	All   = make(map[string]PageInterface)
	Pages *tview.Pages
	log   = logger.Log
)

// Setup should be called once to initialize the pages
func Setup(app *tview.Application) {

	All[config.PageID.Network] = &Network{
		PageType: &PageType{
			ID:  config.PageID.Network,
			App: app,
		},
	}

	All[config.PageID.EthClient] = &ETHClient{
		PageType: &PageType{
			ID:  config.PageID.EthClient,
			App: app,
		},
	}

	All[config.PageID.ExecutionClient] = &ExecutionClient{
		PageType: &PageType{
			ID:  config.PageID.ExecutionClient,
			App: app,
		},
	}

	All[config.PageID.ExecutionClientExternal] = &ExecutionClientExternal{
		PageType: &PageType{
			ID:  config.PageID.ExecutionClientExternal,
			App: app,
		},
	}

	All[config.PageID.ExecutionClientExternalSelection] = &ExecutionClientExternalSelection{
		PageType: &PageType{
			ID:  config.PageID.ExecutionClientExternalSelection,
			App: app,
		},
	}

	All[config.PageID.ExecutionClientExternalSelectedLighthouse] = &ExecutionClientExternalSelectedLighthouse{
		PageType: &PageType{
			ID:  config.PageID.ExecutionClientExternalSelectedLighthouse,
			App: app,
		},
	}

	All[config.PageID.ExecutionClientExternalSelectedPrysm] = &ExecutionClientExternalSelectedPrysm{
		PageType: &PageType{
			ID:  config.PageID.ExecutionClientExternalSelectedPrysm,
			App: app,
		},
	}

	All[config.PageID.ExecutionClientExternalSelectedTeku] = &ExecutionClientExternalSelectedTeku{
		PageType: &PageType{
			ID:  config.PageID.ExecutionClientExternalSelectedTeku,
			App: app,
		},
	}

	All[config.PageID.ConsensusClientSelection] = &ConsensusClientSelection{
		PageType: &PageType{
			ID:  config.PageID.ConsensusClientSelection,
			App: app,
		},
	}

	All[config.PageID.ConsensusClientGraffiti] = &ConsensusClientGraffiti{
		PageType: &PageType{
			ID:  config.PageID.ConsensusClientGraffiti,
			App: app,
		},
	}

	All[config.PageID.ConsensusClientCheckpointSync] = &ConsensusClientCheckpointSync{
		PageType: &PageType{
			ID:  config.PageID.ConsensusClientCheckpointSync,
			App: app,
		},
	}

	All[config.PageID.ConsensusClientDopelgangerProtection] = &ConsensusClientDopelgangerProtection{
		PageType: &PageType{
			ID:  config.PageID.ConsensusClientDopelgangerProtection,
			App: app,
		},
	}

	All[config.PageID.FallbackClients] = &FallbackClients{
		PageType: &PageType{
			ID:  config.PageID.FallbackClients,
			App: app,
		},
	}

	All[config.PageID.FallbackClientsPrysm] = &FallbackClientsPrysm{
		PageType: &PageType{
			ID:  config.PageID.FallbackClientsPrysm,
			App: app,
		},
	}

	All[config.PageID.FallbackClientsLighthouse] = &FallbackClientsLighthouse{
		PageType: &PageType{
			ID:  config.PageID.FallbackClientsLighthouse,
			App: app,
		},
	}

	All[config.PageID.FallbackClientsTeku] = &FallbackClientsTeku{
		PageType: &PageType{
			ID:  config.PageID.FallbackClientsTeku,
			App: app,
		},
	}

	All[config.PageID.Monitoring] = &Monitoring{
		PageType: &PageType{
			ID:  config.PageID.Monitoring,
			App: app,
		},
	}

	All[config.PageID.MEVBoost] = &MEVBoost{
		PageType: &PageType{
			ID:  config.PageID.MEVBoost,
			App: app,
		},
	}

	All[config.PageID.MEVBoostLocal] = &MEVBoostLocal{
		PageType: &PageType{
			ID:  config.PageID.MEVBoostLocal,
			App: app,
		},
	}

	All[config.PageID.MEVBoostExternal] = &MEVBoostExternal{
		PageType: &PageType{
			ID:  config.PageID.MEVBoostExternal,
			App: app,
		},
	}

	All[config.PageID.Confirmation] = &Confirmation{
		PageType: &PageType{
			ID:  config.PageID.Confirmation,
			App: app,
		},
	}

	Pages = tview.NewPages()
	for pageID, page := range All {
		Pages.AddPage(pageID, page.Page(), true, false)
	}
}
