package config

type MEVBoostOption struct {
	LocallyManaged    string
	ExternallyManaged string
}

var MEVBoost = struct {
	Option       MEVBoostOption
	Options      []string
	OptionLabels map[string]string
	Descriptions map[string]string
}{
	Option: MEVBoostOption{
		LocallyManaged:    "local",
		ExternallyManaged: "external",
	},
}

func init() {
	MEVBoost.Options = []string{
		MEVBoost.Option.LocallyManaged,
		MEVBoost.Option.ExternallyManaged,
	}

	MEVBoost.OptionLabels = map[string]string{
		MEVBoost.Option.LocallyManaged:    "Locally managed",
		MEVBoost.Option.ExternallyManaged: "Externally managed",
	}

	MEVBoost.Descriptions = map[string]string{
		MEVBoost.Option.LocallyManaged: `Choose this option if you
would like Stader Node to
take care of the MEV Boost
client for you`,
		MEVBoost.Option.ExternallyManaged: `Enter the externally
managed MEV Boost client URL

e.g. http://192.168.1.46:18550`,
	}
}
