package config

var PageID = struct {
	Network                                   string
	EthClient                                 string
	ExecutionClient                           string
	ExecutionClientExternal                   string
	ExecutionClientExternalSelection          string
	ExecutionClientExternalSelectedLighthouse string
	ExecutionClientExternalSelectedPrysm      string
	ExecutionClientExternalSelectedTeku       string
	ConsensusClientSelection                  string
	ConsensusClientGraffiti                   string
	ConsensusClientCheckpointSync             string
	ConsensusClientDopelgangerProtection      string
	FallbackClients                           string
	FallbackClientsPrysm                      string
	FallbackClientsLighthouse                 string
	FallbackClientsTeku                       string
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
	ExecutionClientExternalSelectedLighthouse: "execution_client_external_selected_lighthouse",
	ExecutionClientExternalSelectedPrysm:      "execution_client_external_selected_prysm",
	ExecutionClientExternalSelectedTeku:       "execution_client_external_selected_teku",
	ConsensusClientSelection:                  "consensus_client_selection",
	ConsensusClientGraffiti:                   "consensus_client_graffiti",
	ConsensusClientCheckpointSync:             "consensus_client_checkpoint_sync",
	ConsensusClientDopelgangerProtection:      "consensus_client_dopelganger_protection",
	FallbackClients:                           "fallback_clients",
	FallbackClientsPrysm:                      "fallback_clients_prysm",
	FallbackClientsLighthouse:                 "fallback_clients_lighthouse",
	FallbackClientsTeku:                       "fallback_clients_teku",
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
}{
	Network:         "Network",
	ETHClient:       "ETH Client",
	ExecutionClient: "Execution client",
	ConsensusClient: "Consensus client",
	FallbackClients: "Fallback clients",
	Monitoring:      "Monitoring",
	MEVBoost:        "MEV-Boost",
}
