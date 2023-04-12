package state

import "github.com/stader-labs/ethcli-ui/configuration/config"

type categoriesOptionType struct {
	Selected string
}

type categoriesType struct {
	Option categoriesOptionType
}

var (
	// It becomes true when user save configuration (save & exit button in footer)
	Saved = false

	// It becomes true when user open configuration (open config button in footer)
	OpenWizard = false

	Categories = categoriesType{
		Option: categoriesOptionType{
			Selected: config.Categories.Option.StaderNode,
		},
	}

	// Value can be string or boolean
	Configuration = make(map[string]interface{})
)

func init() {
	FieldKey := config.GetFieldKey()
	Configuration[FieldKey.E1ec_execution_client_mode] = "Locally Managed"
}
