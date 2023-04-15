package state

import "github.com/stader-labs/ethcli-ui/configuration/config"

type categoriesOptionType struct {
	Selected string
}

type categoriesType struct {
	Option categoriesOptionType
}

var (
	// It becomes true when user save configuration (save & exit button in footer)
	Saved = false

	// It becomes true when user open configuration (open config button in footer)
	OpenWizard = false

	Categories = categoriesType{
		Option: categoriesOptionType{
			Selected: config.Categories.Option.StaderNode,
		},
	}

	// Value can be string or boolean
	Configuration = make(map[string]interface{})
)

func init() {
	FieldKey := config.GetFieldKey()
	Configuration[FieldKey.E1ec_execution_and_consensus_mode] = "Locally Managed"
	Configuration[FieldKey.Sn_node_network] = "Goerli Testnet"
	Configuration[FieldKey.Mev_boost_enabled] = true

	Configuration[FieldKey.Nm_execution_client_metrics_port] = 9105
	Configuration[FieldKey.Nm_beacon_node_metrics_port] = 9100
	Configuration[FieldKey.Nm_validator_client_metrics_port] = 9101
	Configuration[FieldKey.Nm_node_metrics_port] = 9102
}
