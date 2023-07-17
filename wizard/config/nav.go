package config

var PageID = struct {
	Network                                   string
	EthClient                                 string
	ExecutionClient                           string
	ExecutionClientExternal                   string
	ExecutionClientExternalSelection          string
	ConsensusClientExternalSelectedLighthouse string
	ConsensusClientExternalSelectedPrysm      string
	ConsensusClientExternalSelectedLodestar   string
	ConsensusClientExternalSelectedTeku       string
	ConsensusClientSelection                  string
	ConsensusClientGraffiti                   string
	ConsensusClientCheckpointSync             string
	ConsensusClientDopelgangerProtection      string
	FallbackClients                           string
	FallbackClientsPrysm                      string
	FallbackClientsLighthouse                 string
	FallbackClientsLodestar                   string
	FallbackClientsTeku                       string
	FallbackClientsNimbus                     string
	Monitoring                                string
	MEVBoost                                  string
	MEVBoostLocal                             string
	MEVBoostExternal                          string
	Confirmation                              string
}{
	Network:                          "network",
	EthClient:                        "eth_client",
	ExecutionClient:                  "execution_client",
	ExecutionClientExternal:          "execution_client_external",
	ExecutionClientExternalSelection: "execution_client_external_selection",
	ConsensusClientExternalSelectedLighthouse: "consensus_client_external_selected_lighthouse",
	ConsensusClientExternalSelectedPrysm:      "consensus_client_external_selected_prysm",
	ConsensusClientExternalSelectedTeku:       "consensus_client_external_selected_teku",
	ConsensusClientExternalSelectedLodestar:   "consensus_client_external_selected_lodestar",
	ConsensusClientSelection:                  "consensus_client_selection",
	ConsensusClientGraffiti:                   "consensus_client_graffiti",
	ConsensusClientCheckpointSync:             "consensus_client_checkpoint_sync",
	ConsensusClientDopelgangerProtection:      "consensus_client_dopelganger_protection",
	FallbackClients:                           "fallback_clients",
	FallbackClientsPrysm:                      "fallback_clients_prysm",
	FallbackClientsLighthouse:                 "fallback_clients_lighthouse",
	FallbackClientsTeku:                       "fallback_clients_teku",
	FallbackClientsNimbus:                     "fallback_clients_nimbus",
	FallbackClientsLodestar:                   "fallback_clients_lodestar",
	Monitoring:                                "monitoring",
	MEVBoost:                                  "mev_boost",
	MEVBoostLocal:                             "mev_boost_local",
	MEVBoostExternal:                          "mev_boost_external",
	Confirmation:                              "confirmation",
}

var TopNav = struct {
	Network         string
	ETHClient       string
	ExecutionClient string
	ConsensusClient string
	FallbackClients string
	Monitoring      string
	MEVBoost        string
	Confirmation    string
}{
	Network:         "Network",
	ETHClient:       "ETH Client",
	ExecutionClient: "Execution client",
	ConsensusClient: "Consensus client",
	FallbackClients: "Fallback clients",
	Monitoring:      "Monitoring",
	MEVBoost:        "MEV-Boost",
	Confirmation:    "Confirmation",
}
