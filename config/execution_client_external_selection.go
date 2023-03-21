package config

type ExecutionClientExternalSelectionOption struct {
	Lighthouse string
	Prysm      string
	Teku       string
}

var ExecutionClientExternalSelection = struct {
	Option  ExecutionClientExternalSelectionOption
	Options []string
}{
	Option: ExecutionClientExternalSelectionOption{
		Lighthouse: "Lighthouse",
		Prysm:      "Prysm",
		Teku:       "Teku",
	},
}

func init() {
	ExecutionClientExternalSelection.Options = []string{
		ExecutionClientExternalSelection.Option.Lighthouse,
		ExecutionClientExternalSelection.Option.Prysm,
		ExecutionClientExternalSelection.Option.Teku,
	}
}
