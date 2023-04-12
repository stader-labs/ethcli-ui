package config

type FormFieldType struct {
	Label              string
	Key                string
	Type               string
	Readonly           bool
	Hidden             bool
	Options            []string
	OptionDescriptions map[string]string
	Description        string
	Meta               map[string]string
	Children           map[string][]FormFieldType
}

var ConfigurationFields = map[string][]FormFieldType{}

func init() {

}
