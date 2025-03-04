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
Relays: Ultra Sound, Titan (unregulated) and Aestus`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Regulated (All MEV Types)",
			Key:   FieldKey.Mev_boost_pm_enable_regulated,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Regulated (All MEV Types)

Choose this option to activate relays that adhere to government regulations such as OFAC sanctions. "Regulated (All MEV Types)" permits all forms of MEV, including sandwich attacks.
Relays: BloXroute regulated, BloXroute Max Profit, Flashbot, Agnostic, Titan (regulated) and Eden Network`, descriptionSidebarWidth),
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

Activate Flashbot relay which is in accordance with regulations and permits sandwich attack`, descriptionSidebarWidth),
		},
		{
			Label: "Enable bloXroute Regulated",
			Key:   FieldKey.Mev_boost_rm_enable_bloXroute_regulated,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable bloXroute Regulated

Activate bloXroute Regulated relay which is in accordance with regulations and permits sandwich attack. 

All bloXroute relays will reject  block bids if they contain OFAC transactions.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable bloXroute Max Profit",
			Key:   FieldKey.Mev_boost_rm_enable_bloXroute_max_profit,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable bloXroute Max Profit

Activate bloXroute Max Profit relay which is in accordance with regulations and permits sandwich attack. 

All bloXroute relays will reject  block bids if they contain OFAC transactions.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Eden",
			Key:   FieldKey.Mev_boost_rm_enable_eden_network,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Eden

Activate Eden Network relay which is in accordance with regulations and permits sandwich attack`, descriptionSidebarWidth),
		},

		{
			Label: "Enable Ultrasound",
			Key:   FieldKey.Mev_boost_rm_enable_ultra_sound,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Ultrasound

Activate Ultra Sound Ethical relay which is not in accordance with regulations and permits sandwich attack`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Aestus",
			Key:   FieldKey.Mev_boost_rm_enable_aestus,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Aestus

The Aestus MEV-Boost Relay is an independent and non-censoring relay. It is committed to neutrality and the development of a healthy MEV-Boost ecosystem`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Agnostic",
			Key:   FieldKey.Mev_boost_rm_enable_agnostic,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Agnostic

Agnostic Relay is an open-source MEV Boost relay available to anyone, anywhere in the world, without prejudice or privilege. It is an ideal relay for block producers and block builders trying to provide neutral features. Agnostic Relay is powered by the know-how and experience of Gnosis communities and contributors.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Titan Regional (regulated)",
			Key:   FieldKey.Mev_boost_rm_enable_titan_regulated,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Titan Regional (regulated)

The Titan Relay is a high-performance, censorship-resistant relay supporting both regulated and unregulated MEV. It prioritizes efficiency, neutrality, and validator rewards while complying with regulations based on the selected mode.`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Titan Global (unregulated)",
			Key:   FieldKey.Mev_boost_rm_enable_titan_unRegulated,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Titan Global (unregulated)

The Titan Relay is a high-performance, censorship-resistant relay supporting both regulated and unregulated MEV. It prioritizes efficiency, neutrality, and validator rewards while complying with regulations based on the selected mode.`, descriptionSidebarWidth),
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
				"Standard Mode",
				"Custom Mode",
			},
			Description: utils.AddNewLines(`Selection Mode

Choose how the Terminal User Interface displays your alternatives for activating MEV relays.`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				"Standard Mode": utils.AddNewLines(`Standard Mode

Choose this to group the relays based on whether they are subject to regulation and if they are susceptible to sandwich attacks. You can use this if you need to identify the kind of relay you wish to use without going through all of them one by one.`, descriptionSidebarWidth),
				"Custom Mode": utils.AddNewLines(`Custom Mode

Choose this to display each relay and activate them selectively if you already have the knowledge of relays. Use this if you want to gain further insight into the relays and utilize them accordingly.`, descriptionSidebarWidth),
			},
			Children: map[string][]FormFieldType{
				"Standard Mode": profileModeFields,
				"Custom Mode":   relayModeFields,
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
			Description: utils.AddNewLines(`MEV-Boost Mode
			
Would you prefer Stader Node to manage your MEV-Boost (Locally Managed), or would you like to handle it yourself (Externally Managed)?`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				"Locally Managed": utils.AddNewLines(`Locally Managed

Choose this option if you would like Stader Node to take care of the MEV Boost client for you`, descriptionSidebarWidth),
				"Externally Managed": utils.AddNewLines(`Externally Managed

Choose this option if you would like to utilize an existing MEV-Boost client that you already manage.`, descriptionSidebarWidth),
			},
			Children: map[string][]FormFieldType{
				"Locally Managed":    locallyManagedFields,
				"Externally Managed": externallyManagedFields,
			},
		},
	}
}
