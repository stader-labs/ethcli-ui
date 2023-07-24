package config

type ConsensusClientExternalSelectionOption struct {
	Lighthouse string
	Prysm      string
	Teku       string
	Lodestar   string
	Nimbus     string
}

var ConsensusClientExternalSelection = struct {
	Option       ConsensusClientExternalSelectionOption
	Options      []string
	OptionLabels map[string]string
}{
	Option: ConsensusClientExternalSelectionOption{
		Lighthouse: "lighthouse",
		Nimbus:     "nimbus",
		Prysm:      "prysm",
		Teku:       "teku",
		Lodestar:   "lodestar",
	},
}

func init() {
	ConsensusClientExternalSelection.Options = []string{
		ConsensusClientExternalSelection.Option.Lighthouse,
		ConsensusClientExternalSelection.Option.Nimbus,
		ConsensusClientExternalSelection.Option.Prysm,
		ConsensusClientExternalSelection.Option.Teku,
		ConsensusClientExternalSelection.Option.Lodestar,
	}

	ConsensusClientExternalSelection.OptionLabels = map[string]string{
		ConsensusClientExternalSelection.Option.Lighthouse: "Lighthouse",
		ConsensusClientExternalSelection.Option.Nimbus:     "Nimbus",
		ConsensusClientExternalSelection.Option.Prysm:      "Prysm",
		ConsensusClientExternalSelection.Option.Teku:       "Teku",
		ConsensusClientExternalSelection.Option.Lodestar:   "Lodestar",
	}
}
