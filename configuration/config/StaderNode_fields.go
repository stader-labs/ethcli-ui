package config

func init() {
	FieldKey := GetFieldKey()

	ConfigurationFields[Categories.Option.StaderNode] = []FormFieldType{
		{
			Label: "Node Network",
			Key:   FieldKey.Sn_node_network,
			Type:  "select",
			Options: []string{
				EthereumMainnet,
				GoerliTestnet,
				HoleskyTestnet,
			},
			OptionDescriptions: map[string]string{
				GoerliTestnet: `Goerli

Opt for Goerli test network for a secure
and zero-cost environment environment to
executing Stader Node operations. By
choosing this network, you can create
Demo validators using testnet ETH.`,
				EthereumMainnet: `Ethereum Mainnet

Opt for Ethereum primary network,
employing real ETH to perform real
transactions and to create genuine
validators on the network.`,
				HoleskyTestnet: `Holešky

Opt for  Holešky (Holešovice) test network,
which is the next generation of long-lived
testnets for Ethereum. It uses free fake
ETH to make fake
validators.
`,
			},
			Description: `Node Network

Choose the Ethereum network of your
preference
Default: Goerli Testnet`,
		},
		{
			Label: "Project Title",
			Key:   FieldKey.Sn_project_title,
			Type:  "text",
			Description: `Project Title
			
All Docker containers managed by the
Stader Node will have this prefix
attached to them.
Default: Stader`,
		},
		{
			Label: "Storage Location",
			Key:   FieldKey.Sn_storage_location,
			Type:  "text",
			Description: `Storage Location

Create the precise path leading to the
data directory, which accommodates the
encrypted file of your node wallet,
password, and validator keys for all
validators. Incorporating environment
variables in this string is permissible.`,
		},
		// {
		// 	Label: "Maximum Fee Limit (in gwei)",
		// 	Key:   FieldKey.Sn_max_fee_limit,
		// 	Type:  "text",
		// },
		// {
		// 	Label: "Priority Fee (in gwei)",
		// 	Key:   FieldKey.Sn_priority_fee,
		// 	Type:  "text",
		// },
		// {
		// 	Label: "Contract Stake Gas Threshold",
		// 	Key:   FieldKey.Sn_contract_stake_gas_threshold,
		// 	Type:  "text",
		// },
		// {
		// 	Label: "Archive-Mode EC URL",
		// 	Key:   FieldKey.Sn_archive_mode_ec_url,
		// 	Type:  "text",
		// },
	}
}
