package config

import "github.com/stader-labs/ethcli-ui/configuration/logger"

var (
	PageID = pageIDType{
		Categories:           "Categories",
		ConfigurationForm:    "ConfigurationForm",
		NimbusFallbackClient: "NimbusFallbackClient",
	}
	descriptionSidebarWidth = 38
	log                     = logger.Log
)
