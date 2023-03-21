package config

type ExecutionClientOption struct {
	SystemRecommended string
	Geth              string
	Nethermind        string
	Besu              string
}

var ExecutionClient = struct {
	Option       ExecutionClientOption
	Options      []string
	Descriptions map[string]string
}{
	Option: ExecutionClientOption{
		SystemRecommended: "System-recommended",
		Geth:              "geth",
		Nethermind:        "nethermind",
		Besu:              "besu",
	},
}

func init() {
	ExecutionClient.Options = []string{
		ExecutionClient.Option.SystemRecommended,
		ExecutionClient.Option.Geth,
		ExecutionClient.Option.Nethermind,
		ExecutionClient.Option.Besu,
	}

	ExecutionClient.Descriptions = map[string]string{
		ExecutionClient.Option.SystemRecommended: `Let Stader node arbitrarily choose
from a wide range of network clients.
This will enhance the network
diversity and resilience of the
Ethereum ecosystem.`,
		ExecutionClient.Option.Geth: `One of the most popular software
clients maintained by the Ethereum
Foundation, Geth is an open source CLI
developed in the Go Programming
Language. It is designed to be
flexible and customizable, and it
supports a wide range of
functionalities such as secure key
management, consensus mechanisms etc.`,
		ExecutionClient.Option.Nethermind: `Nethermind is a high-performance
Ethereum client built on .NET that
offers fast sync speeds and advanced
features for developers and end-users.
While requiring over 8GB of RAM, it
remains a reliable choice for running
Ethereum nodes.`,
		ExecutionClient.Option.Besu: `Besu, developed by ConsenSys and
written in Java, is a comprehensive
Ethereum protocol client. It utilizes
an innovative storage system called
"Bonsai Trees" to store its chain data
effectively, enabling it to retrieve
historical block states without the
need for pruning.`,
	}
}
