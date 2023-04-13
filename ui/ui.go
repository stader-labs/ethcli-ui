package ui

import (
	"github.com/stader-labs/ethcli-ui/configuration"
	"github.com/stader-labs/ethcli-ui/wizard"
	wpages "github.com/stader-labs/ethcli-ui/wizard/pages"
)

func handleUI(
	initialUI string,
	wizardSettings *wpages.SettingsType,
	configurationSettings *map[string]interface{},
) (
	saved bool,
	wSettings *wpages.SettingsType,
	cSettings *map[string]interface{},
) {
	if initialUI == "wizard" {
		confirmed, openConfiguratin, wSettings := wizard.Run(wizardSettings)

		if openConfiguratin {
			return handleUI("configuration", wSettings, configurationSettings)
		}

		return confirmed, wSettings, configurationSettings
	}

	if initialUI == "configuration" {
		cSaved, cOpenWizard, cSettings := configuration.Run(configurationSettings)
		if cOpenWizard {
			return handleUI("wizard", wizardSettings, cSettings)
		}

		return cSaved, wizardSettings, cSettings
	}

	return false, wizardSettings, configurationSettings
}

func Run(
	wSettings *wpages.SettingsType,
	cSettings *map[string]interface{},
) (
	saved bool,
	wizardSettings *wpages.SettingsType,
	configurationSettings *map[string]interface{},
) {
	return handleUI("configuration", wSettings, cSettings)
}
