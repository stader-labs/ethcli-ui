package config

type pageIDType struct {
	Categories           string
	ConfigurationForm    string
	NimbusFallbackClient string
}

type categoriesOptionType struct {
	StaderNode          string
	FeeAndReward        string
	ETH1ExecutionClient string
	ETH2ConsensusClient string
	FallbackClients     string
	MEVBoost            string
	NodeMonitoring      string
}

type categoriesType struct {
	Option       categoriesOptionType
	Options      []string
	Descriptions map[string]string
}
