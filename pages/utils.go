package pages

import (
	"github.com/pbnjay/memory"
	"math/rand"
	"runtime"
	"time"

	"github.com/stader-labs/ethcli-ui/state"
)

func ChangePage(nextPage string) {
	currentPageName, _ := Pages.GetFrontPage()
	log.Infof("ChangePage: from [%s] to [%s]", currentPageName, nextPage)
	Pages.SwitchToPage(nextPage)
	pageInstance := All[nextPage]
	firstElement := pageInstance.GetFirstElement()
	state.CurrentApp.SetFocus(firstElement)
	pageInstance.OnResume()
}

// TODO - bchain and hamid - refactor the types to a repo to share types b/w wizard and stader node
func GetRandomEcClient() string {
	availableEcClients := []string{"geth", "nethermind", "besu"}

	// If is a low power device ignore nethermind
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024
	isLowPower := (totalMemoryGB < 15 || runtime.GOARCH == "arm64")

	if isLowPower {
		availableEcClients = []string{"geth", "besu"}
	}

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(availableEcClients))

	return availableEcClients[randIndex]
}

func GetRandomCcClient() string {
	availableCcClients := []string{"lighthouse", "nimbus", "prysm", "teku"}

	// If is a low power device ignore teku
	// TODO - bchain - these are heuristics for now. We need to properly come up with these numbers via benchmarks
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024
	isLowPower := (totalMemoryGB < 15 || runtime.GOARCH == "arm64")

	if isLowPower {
		availableCcClients = []string{"lighthouse", "nimbus", "prysm"}
	}

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
// 		nextPage = pid.ExecutionClientSettingsType
// 	} else if currentPage == pid.ExecutionClientSettingsType {
// 		nextPage = pid.EthClient
// 	} else if currentPage == pid.EthClient {
// 		nextPage = pid.Network
// 	}

// 	log.Infof("Next page: %s", nextPage)
// 	ChangePage(nextPage)
// }

type ConsensusClientExternalSelectedLighthouseType struct {
	HTTPUrl string `json:"httpUrl"`
}

type ConsensusClientExternalSelectedPrysmType struct {
	HTTPUrl    string `json:"httpUrl"`
	JSONRpcUrl string `json:"jsonRpcUrl"`
}

type ConsensusClientExternalSelectedTekuType struct {
	HTTPUrl string `json:"httpUrl"`
}

type ConsensusClientExternalType struct {
	Lighthouse ConsensusClientExternalSelectedLighthouseType `json:"lighthouse"`
	Prysm      ConsensusClientExternalSelectedPrysmType      `json:"prysm"`
	Teku       ConsensusClientExternalSelectedTekuType       `json:"teku"`
}

type ConsensusClientSettingsType struct {
	Selection              string                      `json:"selection"`
	ExternalSelection      string                      `json:"externalSelection"`
	Graffit                string                      `json:"graffit"`
	CheckpointUrl          string                      `json:"checkpointUrl"`
	DoppelgangerProtection string                      `json:"doppelgangerProtection"`
	External               ConsensusClientExternalType `json:"external"`
}

type FallbackClientsLighthouseType struct {
	ExecutionClientUrl string `json:"executionClientUrl"`
	BeaconNodeHttpUrl  string `json:"beaconNodeHttpUrl"`
}

type FallbackClientsPrysmType struct {
	ExecutionClientUrl    string `json:"executionClientUrl"`
	BeaconNodeHttpUrl     string `json:"beaconNodeHttpUrl"`
	BeaconNodeJsonRpcpUrl string `json:"beaconNodeJsonRpcpUrl"`
}

type FallbackClientsTekuType struct {
	ExecutionClientUrl string `json:"executionClientUrl"`
	BeaconNodeHttpUrl  string `json:"beaconNodeHttpUrl"`
}

type FallbackClientsSettingsType struct {
	SelectionOption string                        `json:"selectionOption"`
	Lighthouse      FallbackClientsLighthouseType `json:"lighthouse"`
	Prysm           FallbackClientsPrysmType      `json:"prysm"`
	Teku            FallbackClientsTekuType       `json:"teku"`
}

type ExecutionClientExternalType struct {
	HTTPBasedRpcApi      string `json:"httpBasedRpcApi"`
	WebsocketBasedRpcApi string `json:"websocketBasedRpcApi"`
}

type ExecutionClientSettingsType struct {
	SelectionOption string                      `json:"selectionOption"`
	External        ExecutionClientExternalType `json:"external"`
}

type SettingsType struct {
	Confirmed                bool                        `json:"confirmed"`
	Network                  string                      `json:"network"`
	EthClient                string                      `json:"ethClient"`
	ExecutionClient          ExecutionClientSettingsType `json:"executionClient"`
	ConsensusClient          ConsensusClientSettingsType `json:"consensusClient"`
	Monitoring               string                      `json:"monitoring"`
	MEVBoost                 string                      `json:"mevBoost"`
	MEVBoostExternalMevUrl   string                      `json:"mevBoostLocalMevUrl"`
	MEVBoostLocalRegulated   bool                        `json:"mevBoostLocalRegulated"`
	MEVBoostLocalUnregulated bool                        `json:"mevBoostLocalUnregulated"`
	FallbackClients          FallbackClientsSettingsType `json:"fallbackClients"`
}

func GetSettings() SettingsType {
	settings := SettingsType{
		Confirmed: state.Confirmed,
		Network:   state.Network.SelectedOption,
		EthClient: state.ETHClient.SelectedOption,
		ExecutionClient: ExecutionClientSettingsType{
			SelectionOption: state.ExecutionClient.SelectedOption,
			External: ExecutionClientExternalType{
				HTTPBasedRpcApi:      state.ExecutionClientExternal.HTTPBasedRpcApi,
				WebsocketBasedRpcApi: state.ExecutionClientExternal.WebsocketBasedRpcApi,
			},
		},
		ConsensusClient: ConsensusClientSettingsType{
			Selection:              state.ConsensusClient.SelectionSelectedOption,
			ExternalSelection:      state.ConsensusClientExternalSelection.SelectedOption,
			Graffit:                state.ConsensusClient.Graffiti,
			CheckpointUrl:          state.ConsensusClient.CheckpointUrl,
			DoppelgangerProtection: state.ConsensusClient.DopelgangerProtectionSelectedOption,
			External: ConsensusClientExternalType{
				Lighthouse: ConsensusClientExternalSelectedLighthouseType{
					HTTPUrl: state.ConsensusClientExternalSelectedLighthouse.HTTPUrl,
				},
				Prysm: ConsensusClientExternalSelectedPrysmType{
					HTTPUrl:    state.ConsensusClientExternalSelectedPrysm.HTTPUrl,
					JSONRpcUrl: state.ConsensusClientExternalSelectedPrysm.JSONRpcUrl,
				},
				Teku: ConsensusClientExternalSelectedTekuType{
					HTTPUrl: state.ConsensusClientExternalSelectedTeku.HTTPUrl,
				},
			},
		},
		Monitoring:               state.Monitoring.SelectedOption,
		MEVBoost:                 state.MEVBoost.SelectedOption,
		MEVBoostExternalMevUrl:   state.MEVBoostExternal.MevUrl,
		MEVBoostLocalRegulated:   state.MEVBoostLocal.Regulated,
		MEVBoostLocalUnregulated: state.MEVBoostLocal.Unregulated,
		FallbackClients: FallbackClientsSettingsType{
			SelectionOption: state.FallbackClients.SelectedOption,
			Lighthouse: FallbackClientsLighthouseType{
				ExecutionClientUrl: state.FallbackClientsLighthouse.ExecutionClientUrl,
				BeaconNodeHttpUrl:  state.FallbackClientsLighthouse.BeaconNodeHttpUrl,
			},
			Prysm: FallbackClientsPrysmType{
				ExecutionClientUrl:    state.FallbackClientsPrysm.ExecutionClientUrl,
				BeaconNodeHttpUrl:     state.FallbackClientsPrysm.BeaconNodeHttpUrl,
				BeaconNodeJsonRpcpUrl: state.FallbackClientsPrysm.BeaconNodeJsonRpcpUrl,
			},
			Teku: FallbackClientsTekuType{
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
