package config

type MEVBoostLocalOption struct {
	Unregulated string
	Regulated   string
}

var MEVBoostLocal = struct {
	Option       MEVBoostLocalOption
	Options      []string
	Descriptions map[string]string
}{
	Option: MEVBoostLocalOption{
		Unregulated: "Unregulated",
		Regulated:   "Regulated",
	},
}

func init() {
	MEVBoostLocal.Options = []string{
		MEVBoostLocal.Option.Unregulated,
		MEVBoostLocal.Option.Regulated,
	}

	MEVBoostLocal.Descriptions = map[string]string{
		MEVBoostLocal.Option.Unregulated: `Choose this option to
activate relays that
don't adhere to any
sanctions lists and won't
censor transactions.
Unregulated (All MEV
Types) permits for all
forms of MEV, including
sandwich attacks.

Relays: Ultra Sound, Titan (unregulated) and Aestus`,
		MEVBoostLocal.Option.Regulated: `Choose this option to
activate relays that
adhere to government
regulations such as OFAC
sanctions. "Regulated
(All MEV Types)" permits
all forms of MEV,
including sandwich
attacks.

Relays: BloXroute regulated, BloXroute Max Profit, Flashbot, Agnostic, Titan (regulated) and Eden Network`,
	}
}
