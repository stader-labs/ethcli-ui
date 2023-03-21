package config

type StaderNodeOption struct {
	Yes string
	No  string
}

var StaderNode = struct {
	Option  StaderNodeOption
	Options []string
}{
	Option: StaderNodeOption{
		Yes: "Yes",
		No:  "No",
	},
}

func init() {
	StaderNode.Options = []string{
		StaderNode.Option.Yes,
		StaderNode.Option.No,
	}
}
