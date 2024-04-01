package state

import (
	"github.com/stader-labs/ethcli-ui/wizard/config"

	"github.com/rivo/tview"
)

var (
	CurrentApp *tview.Application

	StartPageID string

	// If user successfully completed the Wizard or not
	// Initialised with false, and becomes true on Confirmation page
	Confirmed bool

	// If user wants to open the Configuration Manager or not
	OpenConfigurationUI bool

	Network = struct {
		SelectedOption string
	}{
		SelectedOption: config.Network.Option.HoleskyTestnet,
	}

	ETHClient = struct {
		SelectedOption string
	}{
		SelectedOption: config.ETHClient.Option.LocallyManaged,
	}

	ExecutionClient = struct {
		SelectedOption string
	}{
		SelectedOption: config.ExecutionClient.Option.SystemRecommended,
	}

	ExecutionClientExternal = struct {
		HTTPBasedRpcApi      string
		WebsocketBasedRpcApi string
	}{
		HTTPBasedRpcApi:      "",
		WebsocketBasedRpcApi: "",
	}

	ConsensusClientExternalSelection = struct {
		SelectedOption string
	}{
		SelectedOption: config.ConsensusClientExternalSelection.Option.Lighthouse,
	}

	ConsensusClient = struct {
		SelectionSelectedOption             string
		SelectedStage                       string
		DopelgangerProtectionSelectedOption string
		Graffiti                            string
		CheckpointUrl                       string
	}{
		SelectionSelectedOption:             config.ConsensusClient.Stage.Selection.Option.SystemRecommended,
		SelectedStage:                       config.ConsensusClient.Stage.Selection.Name,
		DopelgangerProtectionSelectedOption: config.ConsensusClient.Stage.DopelgangerProtection.Option.Yes,
		Graffiti:                            "",
		CheckpointUrl:                       "",
	}

	ConsensusClientExternalSelectedLighthouse = struct {
		HTTPUrl string
	}{
		HTTPUrl: "",
	}

	ConsensusClientExternalSelectedNimbus = struct {
		HTTPUrl string
	}{
		HTTPUrl: "",
	}

	ConsensusClientExternalSelectedPrysm = struct {
		HTTPUrl    string
		JSONRpcUrl string
	}{
		HTTPUrl:    "",
		JSONRpcUrl: "",
	}

	ConsensusClientExternalSelectedLodestar = struct {
		HTTPUrl string
	}{
		HTTPUrl: "",
	}

	ConsensusClientExternalSelectedTeku = struct {
		HTTPUrl string
	}{
		HTTPUrl: "",
	}

	FallbackClients = struct {
		SelectedOption string
	}{
		SelectedOption: config.FallbackClients.Option.Yes,
	}

	FallbackClientsLighthouse = struct {
		ExecutionClientUrl string
		BeaconNodeHttpUrl  string
	}{
		ExecutionClientUrl: "",
		BeaconNodeHttpUrl:  "",
	}

	FallbackClientsLodestar = struct {
		ExecutionClientUrl string
		BeaconNodeHttpUrl  string
	}{
		ExecutionClientUrl: "",
		BeaconNodeHttpUrl:  "",
	}

	FallbackClientsPrysm = struct {
		ExecutionClientUrl    string
		BeaconNodeHttpUrl     string
		BeaconNodeJsonRpcpUrl string
	}{
		ExecutionClientUrl:    "",
		BeaconNodeHttpUrl:     "",
		BeaconNodeJsonRpcpUrl: "",
	}

	FallbackClientsTeku = struct {
		ExecutionClientUrl string
		BeaconNodeHttpUrl  string
	}{
		ExecutionClientUrl: "",
		BeaconNodeHttpUrl:  "",
	}

	FallbackClientsNimbus = struct {
		ExecutionClientUrl string
		BeaconNodeHttpUrl  string
	}{
		ExecutionClientUrl: "",
		BeaconNodeHttpUrl:  "",
	}

	Monitoring = struct {
		SelectedOption string
	}{
		SelectedOption: config.StaderNode.Option.Yes,
	}

	MEVBoost = struct {
		SelectedOption string
	}{
		SelectedOption: config.MEVBoost.Option.ExternallyManaged,
	}

	MEVBoostLocal = struct {
		Unregulated bool
		Regulated   bool
	}{
		Unregulated: false,
		Regulated:   false,
	}

	MEVBoostExternal = struct {
		MevUrl string
	}{
		MevUrl: "",
	}

	Confirmation = struct {
		SelectedOption string
	}{
		SelectedOption: config.Confirmation.Option.ReviewSettings,
	}
)

func init() {
	Confirmed = false
}
