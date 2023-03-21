package config

type MonitoringOption struct {
	Yes string
	No  string
}

var Monitoring = struct {
	Option       MonitoringOption
	Options      []string
	Descriptions map[string]string
}{
	Option: MonitoringOption{
		Yes: "Yes",
		No:  "No",
	},
}

func init() {
	Monitoring.Options = []string{
		Monitoring.Option.Yes,
		Monitoring.Option.No,
	}
}
