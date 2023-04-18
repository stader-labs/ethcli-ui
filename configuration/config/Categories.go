package config

var Categories = categoriesType{
	Option: categoriesOptionType{
		StaderNode:          "Stader Node",
		FeeAndReward:        "Fee and Reward",
		ETH1ExecutionClient: "ETH1 - Execution Client",
		ETH2ConsensusClient: "ETH2 - Consensus Client",
		FallbackClients:     "Fallback Clients",
		MEVBoost:            "MEV Boost",
		NodeMonitoring:      "Node Monitoring",
	},
}

func init() {
	Categories.Options = []string{
		Categories.Option.StaderNode,
		Categories.Option.FeeAndReward,
		Categories.Option.ETH1ExecutionClient,
		Categories.Option.ETH2ConsensusClient,
		Categories.Option.FallbackClients,
		Categories.Option.MEVBoost,
		Categories.Option.NodeMonitoring,
	}

	Categories.Descriptions = map[string]string{
		Categories.Option.StaderNode: `Stader Node
		
Choose this option to set up the
preferences for the Stader Node`,
		Categories.Option.FeeAndReward: `Fee and Reward

Choose this option to customize
the transaction fee and reward parameters.`,
		Categories.Option.ETH1ExecutionClient: `ETH1 - Execution Client
		
Please indicate your preferred Execution
client and adjust its configurations
accordingly.`,
		Categories.Option.ETH2ConsensusClient: `ETH2 - Consensus Client
		
Please indicate your preferred
Consensus client and adjust its
configurations accordingly.`,
		Categories.Option.FallbackClients: `Fallback Clients
		
Choose this option to add fallback
client. In case your primary Execution
and Consensus clients go offline for
any reason, you can use an externally-
managed client pair that you trust as
a backup option. This ensures that
your Validator Client and Stader Node
stay connected and continue
functioning properly.`,
		Categories.Option.MEVBoost: `MEV Boost
		
Please choose this option to customize
the settings of MEV-Boost, providing
Validators with MEV earnings.

Note: Refer https://writings.flashbots.net/
writings/why-run-mevboost/ for further
insights into MEV, MEV-Boost and
Flashbots.`,
		Categories.Option.NodeMonitoring: `Node Monitoring
		
Please choose this option to customize
the settings of Stader Node Metric
Monitoring`,
	}
}
