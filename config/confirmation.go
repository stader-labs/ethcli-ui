package config

type ConfirmationOption struct {
	GoBack      string
	SaveAndExit string
}

var Confirmation = struct {
	Option       ConfirmationOption
	Options      []string
	OptionLabels map[string]string
	Descriptions map[string]string
}{
	Option: ConfirmationOption{
		GoBack:      "Go back",
		SaveAndExit: "Save and exit",
	},
}

func init() {
	Confirmation.Options = []string{
		Confirmation.Option.GoBack,
		Confirmation.Option.SaveAndExit,
	}

	Confirmation.OptionLabels = map[string]string{
		Confirmation.Option.GoBack:      "Go back",
		Confirmation.Option.SaveAndExit: "Save and exit",
	}
}
