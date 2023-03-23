package pages

import (
	"math/rand"
	"time"

	"github.com/stader-labs/ethcli-ui/state"
)

func ChangePage(nextPage string) {
	Pages.SwitchToPage(nextPage)
	pageInstance := All[nextPage]
	firstElement := pageInstance.GetFirstElement()
	state.CurrentApp.SetFocus(firstElement)
	pageInstance.OnResume()
}

// TODO - bchain - abstract these to constants
func GetRandomEcClient() string {
	availableEcClients := []string{"geth", "nethermind", "besu"}

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(availableEcClients))

	return availableEcClients[randIndex]
}

// TODO - bchain - We can abstract the random selection logic out
func GetRandomCcClient() string {
	// TODO - bchain - use teku only if enough memory is not there
	availableCcClients := []string{"lighthouse", "nimbus", "prysm", "teku"}

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(availableCcClients))

	return availableCcClients[randIndex]
}

func GetEcClient(ecClient string) string {
	if ecClient == "System-recommended" {
		return GetRandomEcClient()
	}

	return ecClient
}

func GetCcClient(ccClient string) string {
	if ccClient == "System-recommended" {
		return GetRandomCcClient()
	}

	return ccClient
}

// func GoBack(app *tview.Application) {
// 	currentPage, _ := Pages.GetFrontPage()
// 	log.Infof("Current page: %s", currentPage)
// 	pid := config.PageID

// 	if currentPage == pid.Network {
// 		modal := tview.NewModal().
// 			SetText("Do you want to quit the application?").
// 			AddButtons([]string{"Quit", "Cancel"}).
// 			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 				if buttonLabel == "Quit" {
// 					app.Stop()
// 				} else {
// 					app.SetRoot(Pages, true).SetFocus(Pages)
// 				}
// 			})

// 		log.Info("Showing modal to confirm quit")
// 		app.SetRoot(modal, true).SetFocus(modal)
// 		return
// 	}

// 	nextPage := ""
// 	if currentPage == pid.Confirmation {
// 		if state.MEVBoost.SelectedOption == config.MEVBoost.Option.ExternallyManaged {
// 			nextPage = pid.MEVBoostExternal
// 		} else if state.MEVBoost.SelectedOption == config.MEVBoost.Option.LocallyManaged {
// 			nextPage = pid.MEVBoostLocal
// 		} else {
// 			// This should never happen
// 			nextPage = pid.MEVBoost
// 			log.Warnf("This else shouldn't be reached. Current page: %s, next page: %s", currentPage, nextPage)
// 		}
// 	} else if currentPage == pid.MEVBoostExternal || currentPage == pid.MEVBoostLocal {
// 		nextPage = pid.MEVBoost
// 	} else if currentPage == pid.MEVBoost {
// 		nextPage = pid.Monitoring
// 	} else if currentPage == pid.Monitoring {
// 		if state.FallbackClients.SelectedOption == config.FallbackClients.Option.Yes {
// 			if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Prysm {
// 				nextPage = pid.FallbackClientsPrysm
// 			} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.LightHouse {
// 				nextPage = pid.FallbackClientsLighthouse
// 			} else if state.ConsensusClient.SelectionSelectedOption == config.ConsensusClient.Stage.Selection.Option.Teku {
// 				nextPage = pid.FallbackClientsTeku
// 			} else {
// 				nextPage = pid.FallbackClients
// 			}
// 		} else {
// 			nextPage = pid.FallbackClients
// 		}
// 	} else if currentPage == pid.FallbackClientsPrysm || currentPage == pid.FallbackClientsLighthouse || currentPage == pid.FallbackClientsTeku {
// 		nextPage = pid.FallbackClients
// 	} else if currentPage == pid.FallbackClients {
// 		nextPage = pid.ConsensusClientDopelgangerProtection
// 	} else if currentPage == pid.ConsensusClientDopelgangerProtection {
// 		nextPage = pid.ConsensusClientCheckpointSync
// 	} else if currentPage == pid.ConsensusClientCheckpointSync {
// 		nextPage = pid.ConsensusClientGraffiti
// 	} else if currentPage == pid.ConsensusClientGraffiti {
// 		nextPage = pid.ConsensusClientSelection
// 	} else if currentPage == pid.ConsensusClientSelection {
// 		nextPage = pid.ExecutionClient
// 	} else if currentPage == pid.ExecutionClient {
// 		nextPage = pid.EthClient
// 	} else if currentPage == pid.EthClient {
// 		nextPage = pid.Network
// 	}

// 	log.Infof("Next page: %s", nextPage)
// 	ChangePage(nextPage)
// }

type consensusClientExternalSelectedLighthouseType struct {
	HTTPUrl string `json:"httpUrl"`
}

type consensusClientExternalSelectedPrysmType struct {
	HTTPUrl    string `json:"httpUrl"`
	JSONRpcUrl string `json:"jsonRpcUrl"`
}

type consensusClientExternalSelectedTekuType struct {
	HTTPUrl string `json:"httpUrl"`
}

type consensusClientExternalType struct {
	Lighthouse consensusClientExternalSelectedLighthouseType `json:"lighthouse"`
	Prysm      consensusClientExternalSelectedPrysmType      `json:"prysm"`
	Teku       consensusClientExternalSelectedTekuType       `json:"teku"`
}

type consensusClientType struct {
	Selection              string                      `json:"selection"`
	ExternalSelection      string                      `json:"externalSelection"`
	Graffit                string                      `json:"graffit"`
	CheckpointUrl          string                      `json:"checkpointUrl"`
	DoppelgangerProtection string                      `json:"doppelgangerProtection"`
	External               consensusClientExternalType `json:"external"`
}

type fallbackClientsLighthouseType struct {
	ExecutionClientUrl string `json:"executionClient"`
	BeaconNodeHttpUrl  string `json:"beaconNodeHttpUrl"`
}

type fallbackClientsPrysmType struct {
	ExecutionClientUrl    string `json:"executionClient"`
	BeaconNodeHttpUrl     string `json:"beaconNodeHttpUrl"`
	BeaconNodeJsonRpcpUrl string `json:"beaconNodeJsonRpcpUrl"`
}

type fallbackClientsTekuType struct {
	ExecutionClientUrl string `json:"executionClient"`
	BeaconNodeHttpUrl  string `json:"beaconNodeHttpUrl"`
}

type fallbackClientsType struct {
	SelectionOption string                        `json:"selectionOption"`
	Lighthouse      fallbackClientsLighthouseType `json:"lighthouse"`
	Prysm           fallbackClientsPrysmType      `json:"prysm"`
	Teku            fallbackClientsTekuType       `json:"teku"`
}

type executionClientExternal struct {
	HTTPBasedRpcApi      string `json:"httpBasedRpcApi"`
	WebsocketBasedRpcApi string `json:"websocketBasedRpcApi"`
}

type executionClient struct {
	SelectionOption string                  `json:"selectionOption"`
	External        executionClientExternal `json:"external"`
}

type SettingsType struct {
	Confirmed                bool                `json:"confirmed"`
	Network                  string              `json:"network"`
	EthClient                string              `json:"ethClient"`
	ExecutionClient          executionClient     `json:"executionClient"`
	ConsensusClient          consensusClientType `json:"consensusClient"`
	Monitoring               string              `json:"monitoring"`
	MEVBoost                 string              `json:"mevBoost"`
	MEVBoostExternalMevUrl   string              `json:"mevBoostLocalMevUrl"`
	MEVBoostLocalRegulated   bool                `json:"mevBoostLocalRegulated"`
	MEVBoostLocalUnregulated bool                `json:"mevBoostLocalUnregulated"`
	FallbackClients          fallbackClientsType `json:"fallbackClients"`
}

func GetSettings() SettingsType {
	settings := SettingsType{
		Confirmed: state.Confirmed,
		Network:   state.Network.SelectedOption,
		EthClient: state.ETHClient.SelectedOption,
		ExecutionClient: executionClient{
			SelectionOption: state.ExecutionClient.SelectedOption,
			External: executionClientExternal{
				HTTPBasedRpcApi:      state.ExecutionClientExternal.HTTPBasedRpcApi,
				WebsocketBasedRpcApi: state.ExecutionClientExternal.WebsocketBasedRpcApi,
			},
		},
		ConsensusClient: consensusClientType{
			Selection:              state.ConsensusClient.SelectionSelectedOption,
			ExternalSelection:      state.ConsensusClientExternalSelection.SelectedOption,
			Graffit:                state.ConsensusClient.Graffiti,
			CheckpointUrl:          state.ConsensusClient.CheckpointUrl,
			DoppelgangerProtection: state.ConsensusClient.DopelgangerProtectionSelectedOption,
			External: consensusClientExternalType{
				Lighthouse: consensusClientExternalSelectedLighthouseType{
					HTTPUrl: state.ConsensusClientExternalSelectedLighthouse.HTTPUrl,
				},
				Prysm: consensusClientExternalSelectedPrysmType{
					HTTPUrl:    state.ConsensusClientExternalSelectedPrysm.HTTPUrl,
					JSONRpcUrl: state.ConsensusClientExternalSelectedPrysm.JSONRpcUrl,
				},
				Teku: consensusClientExternalSelectedTekuType{
					HTTPUrl: state.ConsensusClientExternalSelectedTeku.HTTPUrl,
				},
			},
		},
		Monitoring:               state.Monitoring.SelectedOption,
		MEVBoost:                 state.MEVBoost.SelectedOption,
		MEVBoostExternalMevUrl:   state.MEVBoostExternal.MevUrl,
		MEVBoostLocalRegulated:   state.MEVBoostLocal.Regulated,
		MEVBoostLocalUnregulated: state.MEVBoostLocal.Unregulated,
		FallbackClients: fallbackClientsType{
			SelectionOption: state.FallbackClients.SelectedOption,
			Lighthouse: fallbackClientsLighthouseType{
				ExecutionClientUrl: state.FallbackClientsLighthouse.ExecutionClientUrl,
				BeaconNodeHttpUrl:  state.FallbackClientsLighthouse.BeaconNodeHttpUrl,
			},
			Prysm: fallbackClientsPrysmType{
				ExecutionClientUrl:    state.FallbackClientsPrysm.ExecutionClientUrl,
				BeaconNodeHttpUrl:     state.FallbackClientsPrysm.BeaconNodeHttpUrl,
				BeaconNodeJsonRpcpUrl: state.FallbackClientsPrysm.BeaconNodeJsonRpcpUrl,
			},
			Teku: fallbackClientsTekuType{
				ExecutionClientUrl: state.FallbackClientsTeku.ExecutionClientUrl,
				BeaconNodeHttpUrl:  state.FallbackClientsTeku.BeaconNodeHttpUrl,
			},
		},
	}

	return settings
}

func SetSettings(settings SettingsType) {
	// Should initialise with false value always
	// state.Confirmed = settings.Confirmed

	state.Network.SelectedOption = settings.Network

	state.ETHClient.SelectedOption = settings.EthClient

	state.ExecutionClient.SelectedOption = settings.ExecutionClient.SelectionOption

	state.ExecutionClientExternal.HTTPBasedRpcApi = settings.ExecutionClient.External.HTTPBasedRpcApi
	state.ExecutionClientExternal.WebsocketBasedRpcApi = settings.ExecutionClient.External.WebsocketBasedRpcApi

	state.ConsensusClient.SelectionSelectedOption = settings.ConsensusClient.Selection
	state.ConsensusClient.Graffiti = settings.ConsensusClient.Graffit
	state.ConsensusClient.CheckpointUrl = settings.ConsensusClient.CheckpointUrl
	state.ConsensusClient.DopelgangerProtectionSelectedOption = settings.ConsensusClient.DoppelgangerProtection

	state.ConsensusClientExternalSelection.SelectedOption = settings.ConsensusClient.ExternalSelection
	state.ConsensusClientExternalSelectedLighthouse.HTTPUrl = settings.ConsensusClient.External.Lighthouse.HTTPUrl
	state.ConsensusClientExternalSelectedPrysm.HTTPUrl = settings.ConsensusClient.External.Prysm.HTTPUrl
	state.ConsensusClientExternalSelectedPrysm.JSONRpcUrl = settings.ConsensusClient.External.Prysm.JSONRpcUrl
	state.ConsensusClientExternalSelectedTeku.HTTPUrl = settings.ConsensusClient.External.Teku.HTTPUrl

	state.Monitoring.SelectedOption = settings.Monitoring

	state.MEVBoost.SelectedOption = settings.MEVBoost

	state.MEVBoostExternal.MevUrl = settings.MEVBoostExternalMevUrl

	state.MEVBoostLocal.Regulated = settings.MEVBoostLocalRegulated
	state.MEVBoostLocal.Unregulated = settings.MEVBoostLocalUnregulated

	state.FallbackClients.SelectedOption = settings.FallbackClients.SelectionOption
	state.FallbackClientsLighthouse.ExecutionClientUrl = settings.FallbackClients.Lighthouse.ExecutionClientUrl
	state.FallbackClientsLighthouse.BeaconNodeHttpUrl = settings.FallbackClients.Lighthouse.BeaconNodeHttpUrl
	state.FallbackClientsPrysm.ExecutionClientUrl = settings.FallbackClients.Prysm.ExecutionClientUrl
	state.FallbackClientsPrysm.BeaconNodeHttpUrl = settings.FallbackClients.Prysm.BeaconNodeHttpUrl
	state.FallbackClientsPrysm.BeaconNodeJsonRpcpUrl = settings.FallbackClients.Prysm.BeaconNodeJsonRpcpUrl
	state.FallbackClientsTeku.ExecutionClientUrl = settings.FallbackClients.Teku.ExecutionClientUrl
}

// func saveSettings() error {
// 	settings := GetSettings()

// 	file, err := os.Create("ethcliui-settings.json")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	err = encoder.Encode(&settings)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON:", err)
// 		return err
// 	}

// 	return nil
// }
