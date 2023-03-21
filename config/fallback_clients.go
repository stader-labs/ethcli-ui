package config

type FllbackClientsOption struct {
	Yes string
	No  string
}

var FallbackClients = struct {
	Option  FllbackClientsOption
	Options []string
}{
	Option: FllbackClientsOption{
		Yes: "Yes",
		No:  "No",
	},
}

func init() {
	FallbackClients.Options = []string{
		FallbackClients.Option.Yes,
		FallbackClients.Option.No,
	}
}
