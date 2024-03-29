package config

type CCSelectionOption struct {
	SystemRecommended string
	LightHouse        string
	Nimbus            string
	Prysm             string
	Teku              string
	Lodestar          string
}

type CCDopelgangerProtectionOption struct {
	Yes string
	No  string
}

type CCStageDopelgangerProtection struct {
	Name         string
	Option       CCDopelgangerProtectionOption
	Options      []string
	OptionLabels map[string]string
}

type CCStageSelection struct {
	Name         string
	Option       CCSelectionOption
	Options      []string
	OptionLabels map[string]string
	Descriptions map[string]string
}

type CSStageForm struct {
	Name string
}

type CCStage struct {
	Selection             CCStageSelection
	Graffiti              CSStageForm
	CheckpointSync        CSStageForm
	DopelgangerProtection CCStageDopelgangerProtection
}

var ConsensusClient = struct {
	Stage  CCStage
	Stages []string
}{
	Stage: CCStage{
		Selection: CCStageSelection{
			Name: "Selection",
			Option: CCSelectionOption{
				SystemRecommended: "System-recommended",
				LightHouse:        "lighthouse",
				Nimbus:            "nimbus",
				Prysm:             "prysm",
				Teku:              "teku",
				Lodestar:          "lodestar",
			},
		},
		Graffiti: CSStageForm{
			Name: "Graffiti",
		},
		CheckpointSync: CSStageForm{
			Name: "Checkpoint sync",
		},
		DopelgangerProtection: CCStageDopelgangerProtection{
			Name: "Doppelganger protection",
			Option: CCDopelgangerProtectionOption{
				Yes: "Yes",
				No:  "No",
			},
		},
	},
}

func init() {
	ConsensusClient.Stages = []string{
		ConsensusClient.Stage.Selection.Name,
		ConsensusClient.Stage.Graffiti.Name,
		ConsensusClient.Stage.CheckpointSync.Name,
		ConsensusClient.Stage.DopelgangerProtection.Name,
	}

	ConsensusClient.Stage.DopelgangerProtection.Options = []string{
		ConsensusClient.Stage.DopelgangerProtection.Option.Yes,
		ConsensusClient.Stage.DopelgangerProtection.Option.No,
	}

	ConsensusClient.Stage.DopelgangerProtection.OptionLabels = map[string]string{
		ConsensusClient.Stage.DopelgangerProtection.Option.Yes: "Yes",
		ConsensusClient.Stage.DopelgangerProtection.Option.No:  "No",
	}

	ConsensusClient.Stage.Selection.Options = []string{
		ConsensusClient.Stage.Selection.Option.SystemRecommended,
		ConsensusClient.Stage.Selection.Option.LightHouse,
		ConsensusClient.Stage.Selection.Option.Nimbus,
		ConsensusClient.Stage.Selection.Option.Prysm,
		ConsensusClient.Stage.Selection.Option.Teku,
		ConsensusClient.Stage.Selection.Option.Lodestar,
	}

	ConsensusClient.Stage.Selection.OptionLabels = map[string]string{
		ConsensusClient.Stage.Selection.Option.SystemRecommended: "System-recommended",
		ConsensusClient.Stage.Selection.Option.LightHouse:        "Lighthouse",
		ConsensusClient.Stage.Selection.Option.Nimbus:            "Nimbus",
		ConsensusClient.Stage.Selection.Option.Prysm:             "Prysm",
		ConsensusClient.Stage.Selection.Option.Teku:              "Teku",
		ConsensusClient.Stage.Selection.Option.Lodestar:          "Lodestar",
	}

	ConsensusClient.Stage.Selection.Descriptions = map[string]string{
		ConsensusClient.Stage.Selection.Option.SystemRecommended: `Let Stader node
arbitrarily choose from a
wide range of network
clients . This will
enhance the network
diversity and resilience
of the Ethereum ecosystem.`,
		ConsensusClient.Stage.Selection.Option.LightHouse: `Lighthouse is a software
client for the Ethereum
2.0 blockchain that is
developed by Sigma Prime
a blockchain engineering
firm based in Australia.
It is written in the Rust
programming language and
is designed to be fast,
efficient, and secure`,
		ConsensusClient.Stage.Selection.Option.Nimbus: `Nimbus is an Ethereum
consensus client that
prioritizes minimal
resource usage, and
is written in Nim - a
language with Python-like
syntax that compiles to C.
Its efficiency enables it
to perform well on any
system.`,
		ConsensusClient.Stage.Selection.Option.Prysm: `Prysm, an Ethereum proof-
of-stake client written in
Go, is developed by
Prysmatic Labs. It
prioritizes usability,
security and reliability
in the implementation of
its consensus protocol.`,
		ConsensusClient.Stage.Selection.Option.Teku: `Teku is an Ethereum consensus client
developed by PegaSys, a branch of
ConsenSys that focuses on building
high-quality clients for Ethereum.
Written in Java, Teku offers impressive
security and scalability features,
although it requires substantial RAM and
CPU resources to operate efficiently.

NOTE: Please note that Teku, while a
powerful client, is quite resource-
intensive. Given your system has limited
RAM, it may not perform as well as you'd
like. We recommend considering a lighter
client option instead to ensure optimal
performance.`,
		ConsensusClient.Stage.Selection.Option.Lodestar: `Lodestar is the 
fifth open-source 
Ethereum consensus client. 
It is written in Typescript 
maintained by ChainSafe Systems. 

Lodestar, their flagship product, 
is a production-capable
Beacon Chain and Validator 
Client uniquely situated as 
the go-to for researchers and 
developers for rapid prototyping 
and browser usage.`,
	}
}
