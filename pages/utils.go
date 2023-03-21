package pages

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/stader-labs/ethcli-ui/state"
)

func ChangePage(nextPage string) {
	Pages.SwitchToPage(nextPage)
	pageInstance := All[nextPage]
	firstElement := pageInstance.GetFirstElement()
	state.CurrentApp.SetFocus(firstElement)
	pageInstance.OnResume()
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

type consensusClientType struct {
	Selection              string `json:"selection"`
	Graffit                string `json:"graffit"`
	CheckpointUrl          string `json:"checkpointUrl"`
	DoppelgangerProtection string `json:"doppelgangerProtection"`
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

type SettingsType struct {
	Network                  string              `json:"network"`
	EthClient                string              `json:"ethClient"`
	ExecutionClient          string              `json:"executionClient"`
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
		Network:         state.Network.SelectedOption,
		EthClient:       state.ETHClient.SelectedOption,
		ExecutionClient: state.ExecutionClient.SelectedOption,
		ConsensusClient: consensusClientType{
			Selection:              state.ConsensusClient.SelectionSelectedOption,
			Graffit:                state.ConsensusClient.Graffiti,
			CheckpointUrl:          state.ConsensusClient.CheckpointUrl,
			DoppelgangerProtection: state.ConsensusClient.DopelgangerProtectionSelectedOption,
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

func saveSettings() error {
	settings := GetSettings()

	file, err := os.Create("ethcliui-settings.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&settings)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	return nil
}
