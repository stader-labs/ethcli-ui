package config

type ConfirmationOption struct {
	GoBack      string
	SaveAndExit string
}

var Confirmation = struct {
	Option       ConfirmationOption
	Options      []string
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
}
