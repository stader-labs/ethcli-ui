package state

import (
	"github.com/stader-labs/ethcli-ui/config"

	"github.com/rivo/tview"
)

var (
	CurrentApp *tview.Application

	Network = struct {
		SelectedOption string
	}{
		SelectedOption: config.Network.Option.GoerliTestnet,
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

	ExecutionClientExternalSelection = struct {
		SelectedOption string
	}{
		SelectedOption: config.ExecutionClientExternalSelection.Option.Lighthouse,
	}

	ExecutionClientExternalSelectedLighthouse = struct {
		HTTPUrl string
	}{
		HTTPUrl: "",
	}

	ExecutionClientExternalSelectedPrysm = struct {
		HTTPUrl    string
		JSONRpcUrl string
	}{
		HTTPUrl:    "",
		JSONRpcUrl: "",
	}

	ExecutionClientExternalSelectedTeku = struct {
		HTTPUrl string
	}{
		HTTPUrl: "",
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
		SelectedOption: config.Confirmation.Option.GoBack,
	}
)
