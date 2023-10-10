package config

import (
	"reflect"
)

type FieldKeyType = struct {
	// App specific, shouldn't be used outside of this package
	App___selected_category string

	Sn_node_network     string
	Sn_project_title    string
	Sn_storage_location string
	// Sn_max_fee_limit                string
	// Sn_priority_fee                 string
	// Sn_contract_stake_gas_threshold string
	// Sn_archive_mode_ec_url          string

	// Fee and Reward
	Fr_max_fee_limit       string
	Fr_priority_fee        string
	Fr_archive_mode_ec_url string

	// ETH1ExecutionClient
	E1ec_execution_and_consensus_mode string

	// ETH1ExecutionClient/Locally Managed
	E1ec_lm_execution_client string

	// ETH1ExecutionClient/Locally Managed/Geth
	E1ec_lm_geth_cache_size       string
	E1ec_lm_geth_max_peers        string
	E1ec_lm_geth_container_tag    string
	E1ec_lm_geth_additional_flags string

	E1ec_lm_common_http_port       string
	E1ec_lm_common_websocket_port  string
	E1ec_lm_common_engine_api_port string
	E1ec_lm_common_expose_rpc_port string
	E1ec_lm_common_p2p_port        string
	E1ec_lm_common_ethstats_label  string
	E1ec_lm_common_ethstats_login  string

	// ETH1ExecutionClient/Locally Managed/Nethermind
	E1ec_lm_nethermind_cache_size         string
	E1ec_lm_nethermind_max_peers          string
	E1ec_lm_nethermind_pruning_cache_size string
	E1ec_lm_nethermind_additional_modules string
	E1ec_lm_nethermind_additional_urls    string
	E1ec_lm_nethermind_container_tag      string
	E1ec_lm_nethermind_additional_flags   string

	// ETH1ExecutionClient/Locally Managed/Besu
	E1ec_lm_besu_jvm_heap_size                 string
	E1ec_lm_besu_max_peers                     string
	E1ec_lm_besu_historical_block_replay_limit string
	E1ec_lm_besu_container_tag                 string
	E1ec_lm_besu_additional_flags              string

	// ETH1ExecutionClient/Externally Managed
	E1ec_em_http_url      string
	E1ec_em_websocket_url string

	// ETH2ConsensusClient
	E2cc_lc_consensus_client string
	E2cc_em_consensus_client string

	E2cc_http_url                          string
	E2cc_custom_graffiti                   string
	E2cc_custom_tag                        string
	E2cc_additional_validator_client_flags string

	// ETH2ConsensusClient | Local
	E2cc_lc_common_graffiti            string
	E2cc_lc_common_checkpoint_sync_url string
	E2cc_lc_common_p2p_port            string
	E2cc_lc_common_http_api_port       string
	E2cc_lc_common_expose_api_port     string

	E2cc_lc_common_doppelganger_detection string

	E2cc_lc_max_peer_lighthouse string
	E2cc_lc_max_peer_nimbus     string
	E2cc_lc_max_peer_prysm      string
	E2cc_lc_max_peer_teku       string
	E2cc_lc_max_peer_lodestar   string

	// ETH2ConsensusClient | Local Prysm
	E2cc_lc_rpc_port_prysm                       string
	E2cc_lc_expose_rpc_port_prysm                string
	E2cc_lc_beacon_container_tag_prysm           string
	E2cc_lc_validator_client_container_tag_prysm string
	E2cc_lc_additional_beacon_node_flags_prysm   string
	E2cc_lc_additional_client_flags_prysm        string

	// ETH2ConsensusClient | Local | Lighthouse
	E2cc_lc_container_tag_lighthouse                string
	E2cc_lc_additional_beacon_node_flags_lighthouse string
	E2cc_lc_additional_client_flags_lighthouse      string

	// ETH2ConsensusClient | Local | Nimbus
	E2cc_lc_container_tag_nimbus    string
	E2cc_lc_additional_flags_nimbus string

	// ETH2ConsensusClient | Local |Teku
	E2cc_lc_jvm_heap_size                     string
	E2cc_lc_enable_archive_mode_teku          string
	E2cc_lc_container_tag_teku                string
	E2cc_lc_additional_beacon_node_flags_teku string
	E2cc_lc_additional_client_flags_teku      string

	// ETH2ConsensusClient |External
	E2cc_em_custom_graffiti_teku               string
	E2cc_em_custom_graffiti_lighthouse         string
	E2cc_em_custom_graffiti_nimbus             string
	E2cc_em_custom_graffiti_prysm              string
	E2cc_em_doppelganger_detection_lighthouse  string
	E2cc_em_doppelganger_detection_nimbus      string
	E2cc_em_doppelganger_detection_prysm       string
	E2cc_em_container_tag_prysm                string
	E2cc_em_container_tag_lighthouse           string
	E2cc_em_container_tag_nimbus               string
	E2cc_em_container_tag_teku                 string
	E2cc_em_additional_client_flags_prysm      string
	E2cc_em_additional_client_flags_lighthouse string
	E2cc_em_additional_client_flags_nimbus     string
	E2cc_em_additional_client_flags_teku       string
	E2cc_em_http_teku                          string
	E2cc_em_http_prysm                         string
	E2cc_em_json_rpc_prysm                     string
	E2cc_em_http_lighthouse                    string
	E2cc_em_http_nimbus                        string

	// ETH2ConsensusClient | Local | Lodestar
	E2cc_lc_container_tag_lodestar                string
	E2cc_lc_additional_beacon_node_flags_lodestar string
	E2cc_lc_additional_client_flags_lodestar      string

	// ETH2ConsensusClient | External | Lodestar lodestar
	E2cc_em_custom_graffiti_lodestar         string
	E2cc_em_doppelganger_detection_lodestar  string
	E2cc_em_container_tag_lodestar           string
	E2cc_em_additional_client_flags_lodestar string
	E2cc_em_http_lodestar                    string

	// FallBackClients
	Fc_use_fallback_clients string

	// FallBackClients/true
	Fc_true_reconnect_delay          string
	Fc_true_execution_client_url     string
	Fc_true_beacon_node_url          string
	Fc_true_beacon_node_json_rpc_url string

	// MEV Boost
	Mev_boost_enabled        string
	Mev_boost_mode           string
	Mev_boost_selection_mode string

	// MEV Boost / Externally Managed
	Mev_boost_em_external_url string

	// MEV Boost / Locally Managed
	Mev_boost_lm_selection_mode string

	// MEV Boost / Profile Mode
	Mev_boost_pm_enable_unregulated string
	Mev_boost_pm_enable_regulated   string
	Mev_boost_pm_port               string
	Mev_boost_pm_expose_api_port    string
	Mev_boost_pm_container_tag      string
	Mev_boost_pm_additional_flags   string

	// MEV Boost / Relay Mode
	Mev_boost_rm_enable_flashbots            string
	Mev_boost_rm_enable_bloXroute_regulated  string
	Mev_boost_rm_enable_bloXroute_max_profit string
	Mev_boost_rm_enable_eden_network         string
	Mev_boost_rm_enable_ultra_sound          string
	Mev_boost_rm_enable_aestus               string
	Mev_boost_rm_port                        string
	Mev_boost_rm_expose_api_port             string
	Mev_boost_rm_container_tag               string
	Mev_boost_rm_additional_flags            string

	// Node Monitoring
	Nm_enable_metrics                  string
	Nm_expose_guardian_port            string
	Nm_enable_oracle_dao_metrics       string
	Nm_execution_client_metrics_port   string
	Nm_beacon_node_metrics_port        string
	Nm_validator_client_metrics_port   string
	Nm_node_metrics_port               string
	Nm_exporter_metrics_port           string
	Nm_guardian_oracle_port            string
	Nm_grafana_port                    string
	Nm_grafana_container_tag           string
	Nm_prometheus_port                 string
	Nm_expose_prometheus_port          string
	Nm_prometheus_container_tag        string
	Nm_additional_prometheus_flags     string
	Nm_allow_root_filesystem_access    string
	Nm_exporter_container_tag          string
	Nm_additional_exporter_flags       string
	Nm_enable_beaconchain_node_metrics string
	// Node Monitoring / Enable Beaconchain Node Metrics / true
	Nm_beaconchain_api_key                   string
	Nm_beaconchain_node_metrics_machine_name string
}

func GetFieldKey() FieldKeyType {
	FieldKey := FieldKeyType{}
	v := reflect.ValueOf(&FieldKey).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		field := v.Field(i)
		field.SetString(fieldName)
	}
	return FieldKey
}
