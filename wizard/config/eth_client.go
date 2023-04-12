package config

type ETHClientOption struct {
	ExternallyManaged string
	LocallyManaged    string
}

var ETHClient = struct {
	Option       ETHClientOption
	Options      []string
	OptionLabels map[string]string
	Descriptions map[string]string
}{
	Option: ETHClientOption{
		ExternallyManaged: "external",
		LocallyManaged:    "local",
	},
}

func init() {
	ETHClient.Options = []string{
		ETHClient.Option.LocallyManaged,
		ETHClient.Option.ExternallyManaged,
	}

	ETHClient.OptionLabels = map[string]string{
		ETHClient.Option.ExternallyManaged: "Externally managed",
		ETHClient.Option.LocallyManaged:    "Locally managed",
	}

	ETHClient.Descriptions = map[string]string{
		ETHClient.Option.ExternallyManaged: `Choose this option if you prefer to
use an existing client that you are
already running and managing outside
of Stader Node.`,
		ETHClient.Option.LocallyManaged: `Choose this option to let the Stader
Node take care of your execution and
consensus client.`,
	}
}
