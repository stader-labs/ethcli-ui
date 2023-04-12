package config

type ConfirmationOption struct {
	ReviewSettings string
	SaveAndExit    string
}

var Confirmation = struct {
	Option       ConfirmationOption
	Options      []string
	OptionLabels map[string]string
	Descriptions map[string]string
}{
	Option: ConfirmationOption{
		ReviewSettings: "Review Settings",
		SaveAndExit:    "Save and exit",
	},
}

func init() {
	Confirmation.Options = []string{
		Confirmation.Option.ReviewSettings,
		Confirmation.Option.SaveAndExit,
	}

	Confirmation.OptionLabels = map[string]string{
		Confirmation.Option.ReviewSettings: "Review Settings",
		Confirmation.Option.SaveAndExit:    "Save and exit",
	}
}
