package config

type MEVBoostOption struct {
	LocallyManaged    string
	ExternallyManaged string
}

var MEVBoost = struct {
	Option       MEVBoostOption
	Options      []string
	Descriptions map[string]string
}{
	Option: MEVBoostOption{
		LocallyManaged:    "Locally managed",
		ExternallyManaged: "Externally managed",
	},
}

func init() {
	MEVBoost.Options = []string{
		MEVBoost.Option.LocallyManaged,
		MEVBoost.Option.ExternallyManaged,
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
