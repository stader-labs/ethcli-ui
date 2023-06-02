package config

type NetworkOption struct {
	EthereumMainnet string
	GoerliTestnet   string
}

var Network = struct {
	Option       NetworkOption
	Options      []string
	OptionLabels map[string]string
	Descriptions map[string]string
}{
	Option: NetworkOption{
		GoerliTestnet:   "prater",
		EthereumMainnet: "mainnet",
	},
}

func init() {
	Network.Options = []string{
		Network.Option.EthereumMainnet,
		Network.Option.GoerliTestnet,
	}

	Network.OptionLabels = map[string]string{
		Network.Option.EthereumMainnet: "Ethereum Mainnet",
		Network.Option.GoerliTestnet:   "Goerli Testnet",
	}

	Network.Descriptions = map[string]string{
		Network.Option.EthereumMainnet: `Ethereum Mainnet is the primary
		production network of Ethereum.
		Running a Stader Node takes some
		know-how and resources, but it
		also offers significant rewards
		and benefits for those who
		contribute to the network's
		security and scalability. If
		you're ready to take on this
		challenge, Ethereum Mainnet
		welcomes you!`,
		Network.Option.GoerliTestnet: `Goerli is a test network that provides
a secure and zero-cost environment for
executing Stader Node operations. By
choosing this network, you can create
Demo validators using testnet ETH.`,
	}
}
