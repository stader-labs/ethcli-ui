package config

type ConsensusClientExternalSelectionOption struct {
	Lighthouse string
	Prysm      string
	Teku       string
}

var ConsensusClientExternalSelection = struct {
	Option  ConsensusClientExternalSelectionOption
	Options []string
}{
	Option: ConsensusClientExternalSelectionOption{
		Lighthouse: "lighthouse",
		Prysm:      "prysm",
		Teku:       "teku",
	},
}

func init() {
	ConsensusClientExternalSelection.Options = []string{
		ConsensusClientExternalSelection.Option.Lighthouse,
		ConsensusClientExternalSelection.Option.Prysm,
		ConsensusClientExternalSelection.Option.Teku,
	}
}
