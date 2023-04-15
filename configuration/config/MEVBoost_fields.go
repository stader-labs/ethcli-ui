package config

import "github.com/stader-labs/ethcli-ui/configuration/utils"

func init() {
	FieldKey := GetFieldKey()

	profileModeFields := []FormFieldType{
		{
			Label: "Enable Unregulated (All MEV Types)",
			Key:   FieldKey.Mev_boost_pm_enable_unregulated,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Unregulated (All MEV Types)

Choose this option to activate relays that don't adhere to any sanctions lists and won't censor transactions. Unregulated (All MEV Types) permits for all forms of MEV, including sandwich attacks.
Relays: Ultra Sound and bloXroute Max Profit`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Regulated (All MEV Types)",
			Key:   FieldKey.Mev_boost_pm_enable_regulated,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Regulated (All MEV Types)

Choose this option to activate relays that adhere to government regulations such as OFAC sanctions. "Regulated (All MEV Types)" permits all forms of MEV, including sandwich attacks.
Relays: Blocknative, Flashbots and Eden Network.`, descriptionSidebarWidth),
		},
		{
			Label: "Port",
			Key:   FieldKey.Mev_boost_pm_port,
			Type:  "int",
			Description: utils.AddNewLines(`Port

Enter the port your MEV-Boost should use for the API.`, descriptionSidebarWidth),
		},
		{
			Label: "Expose API Port",
			Key:   FieldKey.Mev_boost_pm_expose_api_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose API Port

Choose this option to make the ports visible on your local network, granting access to your MEV-Boost API for other local devices.`, descriptionSidebarWidth),
		},
		{
			Label: "Container Tag",
			Key:   FieldKey.Mev_boost_pm_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Container Tag

Please enter the label you wish to use on Docker Hub for the MEV-Boost container.
Default`, descriptionSidebarWidth),
		},
		{
			Label: "Additional Flags",
			Key:   FieldKey.Mev_boost_pm_additional_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Flags

Please enter other flags you might use in conjunction with your MEV-Boost to activate added settings that the Stader Node configuration overlooks.`, descriptionSidebarWidth),
		},
	}

	relayModeFields := []FormFieldType{
		{
			Label: "Enable Flashbots",
			Key:   FieldKey.Mev_boost_rm_enable_flashbots,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Flashbots

Activate Flashbots relay which is in accordance with regulations and permits sandwich attack.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable bloXroute Max Profit",
			Key:   FieldKey.Mev_boost_rm_enable_bloXroute_max_profit,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable bloXroute Max Profit

Activate bloxRoute Max Profit relay which is not in accordance with regulations and permits sandwich attack.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Blocknative",
			Key:   FieldKey.Mev_boost_rm_enable_blocknative,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Blocknative

Activate Blocknative relay which is in accordance with regulations and permits sandwich attack.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Eden Network",
			Key:   FieldKey.Mev_boost_rm_enable_eden_network,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Eden Network

Activate Eden Network relay which is in accordance with regulations and permits sandwich attack.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Ultra Sound",
			Key:   FieldKey.Mev_boost_rm_enable_ultra_sound,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Ultra Sound

Activate Ultra Sound relay which is not in accordance with regulations and permits sandwich attack.`, descriptionSidebarWidth),
		},
		{
			Label: "Port",
			Key:   FieldKey.Mev_boost_rm_port,
			Type:  "int",
			Description: utils.AddNewLines(`Port

Enter the port your MEV-Boost should use for the API.`, descriptionSidebarWidth),
		},
		{
			Label: "Expose API Port",
			Key:   FieldKey.Mev_boost_rm_expose_api_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose API Port

Choose this option to make the ports visible on your local network, granting access to your MEV-Boost API for other local devices.`, descriptionSidebarWidth),
		},
		{
			Label: "Container Tag",
			Key:   FieldKey.Mev_boost_rm_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Container Tag

Please enter the label you wish to use on Docker Hub for the MEV-Boost container.
Default`, descriptionSidebarWidth),
		},
		{
			Label: "Additional Flags",
			Key:   FieldKey.Mev_boost_rm_additional_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Flags

Please enter other flags you might use in conjunction with your MEV-Boost to activate added settings that the Stader Node configuration overlooks.`, descriptionSidebarWidth),
		},
	}

	locallyManagedFields := []FormFieldType{
		{
			Label: "Selection Mode",
			Key:   FieldKey.Mev_boost_lm_selection_mode,
			Type:  "select",
			Options: []string{
				"Profile mode",
				"Relay mode",
			},
			Description: utils.AddNewLines(`Selection Mode

Choose how the Terminal User Interface displays your alternatives for activating MEV relays.`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				"Profile mode": utils.AddNewLines(`Profile mode

Choose this to group the relays based on whether they are subject to regulation and if they are susceptible to sandwich attacks. You can use this if you need to identify the kind of relay you wish to use without going through all of them one by one.`, descriptionSidebarWidth),
				"Relay mode": utils.AddNewLines(`Relay mode

Choose this to display each relay and activate them selectively if you already have the knowledge of relays. Use this if you want to gain further insight into the relays and utilize them accordingly.`, descriptionSidebarWidth),
			},
			Children: map[string][]FormFieldType{
				"Profile mode": profileModeFields,
				"Relay mode":   relayModeFields,
			},
		},
	}

	externallyManagedFields := []FormFieldType{
		{
			Label: "MEV URL",
			Key:   FieldKey.Mev_boost_em_external_url,
			Type:  "text",
			Description: utils.AddNewLines(`MEV URL
			
Enter the externally managed MEV Boost client URL`, descriptionSidebarWidth),
		},
	}

	ConfigurationFields[Categories.Option.MEVBoost] = []FormFieldType{
		{
			Label: "MEV-Boost Mode",
			Key:   FieldKey.Mev_boost_mode,
			Type:  "select",
			Options: []string{
				"Locally Managed",
				"Externally Managed",
			},
			OptionDescriptions: map[string]string{
				"Locally Managed": utils.AddNewLines(`Locally Managed

Choose this option if you would like Stader Node to take care of the MEV Boost client for you`, descriptionSidebarWidth),
				"Externally Managed": utils.AddNewLines(`Externally Managed

Choose this option if you would like to manage the MEV Boost client yourself`, descriptionSidebarWidth),
			},
			Children: map[string][]FormFieldType{
				"Locally Managed":    locallyManagedFields,
				"Externally Managed": externallyManagedFields,
			},
		},
	}

	// 	ConfigurationFields[Categories.Option.MEVBoost] = []FormFieldType{
	// 		{
	// 			Label: "Enable MEV-Boost",
	// 			Key:   FieldKey.Mev_boost_enabled,
	// 			Type:  "checkbox",
	// 			Description: utils.AddNewLines(`Enable MEV-Boost.

	// MEV-Boost gives your validator the opportunity to link up with relays that act as a go-between for you and experienced block builders. These experts are able to hunt out and make use of MEV openings, and in return for your services, could give you a fair reward that is often more than what you could have made on your own.
	// Default: True`, descriptionSidebarWidth),
	// 			Children: map[string][]FormFieldType{
	// 				"true": {
	// 					{
	// 						Label: "MEV-Boost Mode",
	// 						Key:   FieldKey.Mev_boost_mode,
	// 						Type:  "select",
	// 						Options: []string{
	// 							"Locally Managed",
	// 							"Externally Managed",
	// 						},
	// 						OptionDescriptions: map[string]string{
	// 							"Locally Managed": utils.AddNewLines(`Locally Managed

	// Choose this option if you would like Stader Node to take care of the MEV Boost client for you`, descriptionSidebarWidth),
	// 							"Externally Managed": utils.AddNewLines(`Externally Managed

	// Choose this option if you would like to manage the MEV Boost client yourself`, descriptionSidebarWidth),
	//
	//						},
	//						Children: map[string][]FormFieldType{
	//							"Locally Managed":    locallyManagedFields,
	//							"Externally Managed": externallyManagedFields,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	}
}
