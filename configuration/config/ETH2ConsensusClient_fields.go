package config

import (
	"github.com/stader-labs/ethcli-ui/configuration/utils"
)

func init() {
	FieldKey := GetFieldKey()

	locallyManagedFields := []FormFieldType{
		{
			Label: "Consensus client",
			Key:   FieldKey.E2cc_lc_consensus_client,
			Type:  "select",
			Options: []string{
				"Teku",
				"Lighthouse",
				"Nimbus",
				"Prysm",
				"Lodestar",
			},
			Description: utils.AddNewLines(`ETH 2 - Consensus Client

Choose the ETH 2 - Consensus Client that you wish to use.`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				"Teku": utils.AddNewLines(`Teku

Teku is an Ethereum consensus client developed by PegaSys, a branch of ConsenSys that focuses on building high-quality clients for Ethereum. Written in Java, Teku offers impressive security and scalability features, although it requires substantial RAM and CPU resources to operate efficiently.`, descriptionSidebarWidth),
				"Lighthouse": utils.AddNewLines(`Lighthouse

Lighthouse is a software client for the Ethereum 2.0 blockchain that is developed by Sigma Prime, a blockchain engineering firm based in Australia. It is written in the Rust programming language and is designed to be fast, efficient, and secure`, descriptionSidebarWidth),
				"Nimbus": utils.AddNewLines(`Nimbus

Nimbus is an Ethereum consensus client that prioritizes minimal resource usage, and is written in Nim - a language with Python-like syntax that compiles to C. Its efficiency enables it to perform well on any system.`, descriptionSidebarWidth),
				"Prysm": utils.AddNewLines(`Prysm

Prysm, an Ethereum proof-of-stake client written in Go, is developed by Prysmatic Labs. It prioritizes usability, security and reliability in the implementation of its consensus protocol.`, descriptionSidebarWidth),
				"Lodestar": utils.AddNewLines(`"Lodestar

Lodestar is the fifth open-source Ethereum consensus client. It is written in Typescript maintained by ChainSafe Systems. Lodestar, their flagship product, is a production-capable Beacon Chain and Validator Client uniquely situated as the go-to for researchers and developers for rapid prototyping and browser usage.",`, descriptionSidebarWidth),
			},
			Children: makeLocalManagerChildren(),
		},
	}

	externallyManagedFields := []FormFieldType{
		{
			Label: "Consensus client",
			Key:   FieldKey.E2cc_em_consensus_client,
			Type:  "select",
			Options: []string{
				"Teku",
				"Lighthouse",
				"Nimbus",
				"Prysm",
				"Lodestar",
			},
			Description: utils.AddNewLines(`Consensus Client

Choose the ETH 2 - Consensus Client that you manage externally.`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				"Teku": utils.AddNewLines(`Teku
Choose this option if you wish to use Teku as your externally manager ETH 2 - Consensus Client.`, descriptionSidebarWidth),
				"Lighthouse": utils.AddNewLines(`Lighthouse
Choose this option if you wish to use Lighthouse as your externally manager ETH 2 - Consensus Client.`, descriptionSidebarWidth),
				"Prysm": utils.AddNewLines(`Prysm
Choose this option if you wish to use Prysm as your externally manager ETH 2 - Consensus Client.`, descriptionSidebarWidth),
				"Lodestar": utils.AddNewLines(`Lodestar
Choose this option if you wish to use Lodestar as your externally manager ETH 2 - Consensus Client.`, descriptionSidebarWidth),
				"Nimbus": utils.AddNewLines(`Nimbus
Choose this option if you wish to use Nimbus as your externally manager ETH 2 - Consensus Client.`, descriptionSidebarWidth),
			},
			Children: makeExternalChildren(),
		},
	}

	ConfigurationFields[Categories.Option.ETH2ConsensusClient] = []FormFieldType{
		{
			Label: "Consensus Client Preference",
			Key:   FieldKey.E1ec_execution_and_consensus_mode,
			Type:  "select",
			Options: []string{
				"Locally Managed",
				"Externally Managed",
			},
			Description: utils.AddNewLines(`Consensus Client Preference

Select your preferred method for managing your Consensus Client.
Default: Locally Managed`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				"Locally Managed": utils.AddNewLines(`Locally Managed
	
Choose this option to let the Stader Node take care of your Execution Client.`, descriptionSidebarWidth),
				"Externally Managed": utils.AddNewLines(`Externally Managed

Choose this option if you prefer to use an existing client that you are already running and managing outside of Stader.`, descriptionSidebarWidth),
			},
			Children: map[string][]FormFieldType{
				"Locally Managed":    locallyManagedFields,
				"Externally Managed": externallyManagedFields,
			},
		},
	}
}

func makeMaxPeerField(key string) FormFieldType {
	FieldKey := GetFieldKey()
	defaultValue := "100"
	if key == FieldKey.E2cc_lc_max_peer_lighthouse {
		defaultValue = "80"
	} else if key == FieldKey.E2cc_lc_max_peer_prysm || key == FieldKey.E2cc_lc_max_peer_lodestar {
		defaultValue = "50"
	}

	return FormFieldType{
		Label: "Max Peers",
		Key:   key,
		Type:  "int",
		Description: utils.AddNewLines(`Max Peers

Please specify your ETH2 - Consensus client peer limit. The number can be decreased to enhance performance on systems with limited specs.
Default: `+defaultValue, descriptionSidebarWidth),
	}
}

func makeContainerTagField(key string) FormFieldType {
	fieldKey := GetFieldKey()
	description := ""
	switch key {
	case fieldKey.E2cc_em_container_tag_prysm:
		description = utils.AddNewLines(`Container tag
Please enter the label you wish to use on Docker Hub for the Prysm container. The label will be utilized for the Validator Client that Stader manages with your validator keys.
Default:`, descriptionSidebarWidth)
	case fieldKey.E2cc_em_container_tag_lighthouse:
		description = utils.AddNewLines(`Container tag
Please enter the label you wish to use on Docker Hub for the Lighthouse container. The label will be utilized for the Validator Client that Stader manages with your validator keys.
Default:`, descriptionSidebarWidth)
	case fieldKey.E2cc_em_container_tag_teku:
		description = utils.AddNewLines(`Container tag
Please enter the label you wish to use on Docker Hub for the Teku container. The label will be utilized for the Validator Client that Stader manages with your validator keys.
Default:`, descriptionSidebarWidth)
	}

	return FormFieldType{
		Label:       "Container Tag",
		Key:         key,
		Type:        "text",
		Description: description,
	}
}

func makeHTTPField(key string) FormFieldType {
	return FormFieldType{
		Label: "HTTP URL",
		Key:   key,
		Type:  "text",
		Description: utils.AddNewLines(`HTTP URL

Please enter the HTTP API URL of your externally managed Prysm client.
Note: When running this client on the same machine as the Stader Node, use your machine's LAN IP address instead of localhost or 127.0.0.1`, descriptionSidebarWidth),
	}
}

func makeJSONRPCField(key string) FormFieldType {
	return FormFieldType{
		Label: "JSON-RPC URL",
		Key:   key,
		Type:  "text",
		Description: utils.AddNewLines(`JSON-RPC URL
Please enter the JSON-RPC URL of your externally managed Prysm client.
Note: When running this client on the same machine as the Stader Node, use your machine's LAN IP address instead of localhost or 127.0.0.1.Default`, descriptionSidebarWidth),
	}
}

func makeCustomGraffitiField(key string) FormFieldType {
	return FormFieldType{
		Label:    "Custom Graffiti",
		Key:      key,
		Type:     "text",
		MaxChars: 16,
		Description: utils.AddNewLines(`Custom Graffiti

Want to add a personal touch to your proposed blocks? You have the option to add a short custom message, or "graffiti", of up to 16 characters to each block proposed by your validators.
Default: Stader`, descriptionSidebarWidth),
	}
}

func makeDoppelgängerField(key string) FormFieldType {
	return FormFieldType{
		Label: "Enable Doppelgänger Detection",
		Key:   key,
		Type:  "checkbox",
		Description: utils.AddNewLines(`Enable Doppelgänger Detection

Doppleganger protection safeguards your validators from being slashed and removed from the Beacon Chain, just in case you accidentally run your validator keys on multiple machines simultaneously.Here's how it works: whenever your validator client restarts, it will deliberately skip 2-3 attestations for each validator. If all of them are skipped successfully, you can be rest assured that you're good to start attesting again.
Default: True`, descriptionSidebarWidth),
	}
}

func makeAdditionValidatorClientField(key string) FormFieldType {
	return FormFieldType{
		Label: "Additional Validator Client Flags",
		Key:   key,
		Type:  "text",
		Description: utils.AddNewLines(`Additional Validator Client Flags

Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Validator Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
	}
}

func makeConsensusLocalManagerClientField(ccClient string) []FormFieldType {
	FieldKey := GetFieldKey()

	commonsField := []FormFieldType{
		{
			Label:    "Consensus Graffiti",
			Key:      FieldKey.E2cc_lc_common_graffiti,
			Type:     "text",
			MaxChars: 16,
			Description: utils.AddNewLines(`Consensus Graffiti
Want to add a personal touch to your proposed blocks? You have the option to add a short custom message, or "graffiti", of up to 16 characters to each block proposed by your validators.
Default: Stader`, descriptionSidebarWidth),
		},
		{
			Label: "Checkpoint Sync URL",
			Key:   FieldKey.E2cc_lc_common_checkpoint_sync_url,
			Type:  "text",
			Description: utils.AddNewLines(`Checkpoint Sync URL
Your client is equipped with the Checkpoint Sync feature, which can significantly reduce the time and effort required to sync from scratch. With this powerful functionality, your client can instantly copy the most recent state from a trusted Consensus client that you specify.`, descriptionSidebarWidth),
		},
		{
			Label: "P2P Port",
			Key:   FieldKey.E2cc_lc_common_p2p_port,
			Type:  "int",
			Description: utils.AddNewLines(`P2P Port
Enter the port number that the Consensus Client will use for P2P communication with other blockchain nodes.
Default: 9001`, descriptionSidebarWidth),
		},
		{
			Label: "HTTP API Port",
			Key:   FieldKey.E2cc_lc_common_http_api_port,
			Type:  "int",
			Description: utils.AddNewLines(`HTTP API Port

Enter the port your HTTP API should use
Default: 5052`, descriptionSidebarWidth),
		},
		{
			Label: "Expose API Port",
			Key:   FieldKey.E2cc_lc_common_expose_api_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose API Port

Choose this option to make the API ports visible on your local network, granting access to your Consensus Client's API endpoint for other local devices.
Default: False`, descriptionSidebarWidth),
		},
		{
			Label: "Enable Doppelgänger Detection",
			Key:   FieldKey.E2cc_lc_common_doppelganger_detection,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Doppelgänger Detection

Doppleganger protection safeguards your validators from being slashed and removed from the Beacon Chain, just in case you accidentally run your validator keys on multiple machines simultaneously.Here's how it works: whenever your validator client restarts, it will deliberately skip 2-3 attestations for each validator. If all of them are skipped successfully, you can be rest assured that you're good to start attesting again.
Default - True`, descriptionSidebarWidth),
		},
	}

	switch ccClient {
	case "Lighthouse":
		commonsField = appendFieldLightHouse(commonsField)
	case "Nimbus":
		commonsField = appendFieldNimbus(commonsField)
	case "Prysm":
		commonsField = appendFieldPrysm(commonsField)
	case "Teku":
		commonsField = appendFieldTeku(commonsField)
	case "Lodestar":
		commonsField = appendFieldLodestar(commonsField)
	default:
	}

	return commonsField
}

func appendFieldLightHouse(commonsField []FormFieldType) []FormFieldType {
	FieldKey := GetFieldKey()

	commonsField = append(
		commonsField,
		makeMaxPeerField(FieldKey.E2cc_lc_max_peer_lighthouse),
		FormFieldType{
			Label: "Container Tag",
			Key:   FieldKey.E2cc_lc_container_tag_lighthouse,
			Type:  "text",
			Description: utils.AddNewLines(`Container Tag
Please enter the label you wish to use on Docker Hub for the Lighthouse container. The label will be utilized for the Validator Client that Stader manages with your validator keys.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Beacon Node Flags",
			Key:   FieldKey.E2cc_lc_additional_beacon_node_flags_lighthouse,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Beacon Node Flags
Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Beacon Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Validator Client Flags",
			Key:   FieldKey.E2cc_lc_additional_client_flags_lighthouse,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Validator Client Flags
Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Validator Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
	)

	return commonsField
}

func appendFieldNimbus(commonsField []FormFieldType) []FormFieldType {
	FieldKey := GetFieldKey()

	commonsField = append(
		commonsField,
		makeMaxPeerField(FieldKey.E2cc_lc_max_peer_nimbus),
		FormFieldType{
			Label: "Container Tag",
			Key:   FieldKey.E2cc_lc_container_tag_nimbus,
			Type:  "text",
			Description: utils.AddNewLines(`Container Tag
Please enter the label you wish to use on Docker Hub for the Nimbus container.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Flags",
			Key:   FieldKey.E2cc_lc_additional_flags_nimbus,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Flags
Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Beacon Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
	)

	return commonsField
}

func appendFieldLodestar(commonsField []FormFieldType) []FormFieldType {
	FieldKey := GetFieldKey()

	commonsField = append(
		commonsField,
		makeMaxPeerField(FieldKey.E2cc_lc_max_peer_lodestar),
		FormFieldType{
			Label: "Container Tag",
			Key:   FieldKey.E2cc_lc_container_tag_lodestar,
			Type:  "text",
			Description: utils.AddNewLines(`Container Tag
Please enter the label you wish to use on Docker Hub for the Lodestar container.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Flags",
			Key:   FieldKey.E2cc_lc_additional_beacon_node_flags_lodestar,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Flags
Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Beacon Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Validator client Flags",
			Key:   FieldKey.E2cc_lc_additional_client_flags_lodestar,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Validator client Flags

Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Validator Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
	)

	return commonsField
}

func appendFieldPrysm(commonsField []FormFieldType) []FormFieldType {
	FieldKey := GetFieldKey()

	commonsField = append(commonsField, makeMaxPeerField(FieldKey.E2cc_lc_max_peer_prysm))
	commonsField = append(
		commonsField,
		FormFieldType{
			Label: "RPC Port",
			Key:   FieldKey.E2cc_lc_rpc_port_prysm,
			Type:  "int",
			Description: utils.AddNewLines(`RPC Port

Please enter the port on which Prysm should run its JSON-RPC API.
Default: 5053`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Expose RPC Port",
			Key:   FieldKey.E2cc_lc_expose_rpc_port_prysm,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose RPC Port

Choose this option to make the RPC ports visible on your local network, granting access to your Consensus Client's RPC endpoint for other local devices.
Default: False`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Beacon Node Container Tag",
			Key:   FieldKey.E2cc_lc_beacon_container_tag_prysm,
			Type:  "text",
			Description: utils.AddNewLines(`Beacon Node Container Tag

Please enter the label you wish to use on Docker Hub for the Prysm Beacon Node container.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Validator Client Container Tag",
			Key:   FieldKey.E2cc_lc_validator_client_container_tag_prysm,
			Type:  "text",
			Description: utils.AddNewLines(`Validator Client Container Tag

Please enter the label you wish to use on Docker Hub for the Prysm Validator Node container.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Beacon Node Flags",
			Key:   FieldKey.E2cc_lc_additional_beacon_node_flags_prysm,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Beacon Node Flags

Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Beacon Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Validator client Flags",
			Key:   FieldKey.E2cc_lc_additional_client_flags_prysm,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Validator client Flags

Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Validator Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth),
		},
	)

	//
	return commonsField
}

func appendFieldTeku(commonsField []FormFieldType) []FormFieldType { //JVM Heap Size
	FieldKey := GetFieldKey()

	commonsField = append(commonsField,
		FormFieldType{
			Label: "JVM Heap Size",
			Key:   FieldKey.E2cc_lc_jvm_heap_size,
			Type:  "int",
			Description: utils.AddNewLines(`JVM Heap Size
The upper threshold of the RAM, expressed in MB, that the JVM in Teku should restrict itself to. Lowering this value would prompt Teku to consume less RAM, although it would inevitably exceed this cap. If you prefer an automatic allocation, input 0.
Default: 2048`, descriptionSidebarWidth),
		},
		makeMaxPeerField(FieldKey.E2cc_lc_max_peer_teku),
		FormFieldType{
			Label: "Enable Archive mode",
			Key:   FieldKey.E2cc_lc_enable_archive_mode_teku,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Archive mode
Enabling Archive Mode triggers Teku to operate in "archival" mode, allowing it to replicate the state of the Beacon chain for a preceding block. This feature is essential for producing the Merkle rewards tree manually. If you are confident that you will never have to manually create a tree, you can opt-out of the archival mode.
Default: False`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Container Tag",
			Key:   FieldKey.E2cc_lc_container_tag_teku,
			Type:  "text",
			Description: utils.AddNewLines(`Container Tag
Please enter the label you wish to use on Docker Hub for the Teku container. The label will be utilized for the Validator Client that Stader manages with your validator keys.
Default:`, descriptionSidebarWidth),
		},
		FormFieldType{
			Label: "Additional Beacon Node Flags",
			Key:   FieldKey.E2cc_lc_additional_beacon_node_flags_teku,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Beacon Node Flags
Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Beacon Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth)},
		FormFieldType{
			Label: "Additional Validator Client Flags",
			Key:   FieldKey.E2cc_lc_additional_client_flags_teku,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Validator Client Flags
Please enter other flags you might use in conjunction with your ETH2 - Consensus Client Validator Node to activate added settings that the Stader Node configuration overlooks.
Default:`, descriptionSidebarWidth)})

	return commonsField
}

func makeConsensusExternalField(ccClient string) []FormFieldType {
	FieldKey := GetFieldKey()

	var commonsField []FormFieldType

	switch ccClient {
	case "Prysm":
		commonsField = append(commonsField, makeHTTPField(FieldKey.E2cc_em_http_prysm))
		commonsField = append(commonsField, makeJSONRPCField(FieldKey.E2cc_em_json_rpc_prysm))
		commonsField = append(commonsField, makeCustomGraffitiField(FieldKey.E2cc_em_custom_graffiti_prysm))
		commonsField = append(commonsField, makeDoppelgängerField(FieldKey.E2cc_em_doppelganger_detection_prysm))
		commonsField = append(commonsField, makeContainerTagField(FieldKey.E2cc_em_container_tag_prysm))
		commonsField = append(commonsField, makeAdditionValidatorClientField(FieldKey.E2cc_em_container_tag_prysm))
	case "Lighthouse":
		commonsField = append(commonsField, makeHTTPField(FieldKey.E2cc_em_http_lighthouse))
		commonsField = append(commonsField, makeCustomGraffitiField(FieldKey.E2cc_em_custom_graffiti_lighthouse))
		commonsField = append(commonsField, makeDoppelgängerField(FieldKey.E2cc_em_doppelganger_detection_lighthouse))
		commonsField = append(commonsField, makeContainerTagField(FieldKey.E2cc_em_container_tag_lighthouse))
		commonsField = append(commonsField, makeAdditionValidatorClientField(FieldKey.E2cc_em_additional_client_flags_lighthouse))
	case "Nimbus":
		commonsField = append(commonsField, makeHTTPField(FieldKey.E2cc_em_http_nimbus))
		commonsField = append(commonsField, makeCustomGraffitiField(FieldKey.E2cc_em_custom_graffiti_nimbus))
		commonsField = append(commonsField, makeDoppelgängerField(FieldKey.E2cc_em_doppelganger_detection_nimbus))
		commonsField = append(commonsField, makeContainerTagField(FieldKey.E2cc_em_container_tag_nimbus))
		commonsField = append(commonsField, makeAdditionValidatorClientField(FieldKey.E2cc_em_additional_client_flags_nimbus))
	case "Teku":
		commonsField = append(commonsField, makeHTTPField(FieldKey.E2cc_em_http_teku))
		commonsField = append(commonsField, makeCustomGraffitiField(FieldKey.E2cc_em_custom_graffiti_teku))
		commonsField = append(commonsField, makeDoppelgängerField(FieldKey.E2cc_em_doppelganger_detection_teku))
		commonsField = append(commonsField, makeContainerTagField(FieldKey.E2cc_em_container_tag_teku))
		commonsField = append(commonsField, makeAdditionValidatorClientField(FieldKey.E2cc_em_additional_client_flags_teku))
	case "Lodestar":
		commonsField = append(commonsField, makeHTTPField(FieldKey.E2cc_em_http_lodestar))
		commonsField = append(commonsField, makeCustomGraffitiField(FieldKey.E2cc_em_custom_graffiti_lodestar))
		commonsField = append(commonsField, makeDoppelgängerField(FieldKey.E2cc_em_doppelganger_detection_lodestar))
		commonsField = append(commonsField, makeContainerTagField(FieldKey.E2cc_em_container_tag_lodestar))
		commonsField = append(commonsField, makeAdditionValidatorClientField(FieldKey.E2cc_em_additional_client_flags_lodestar))
	}

	return commonsField
}

func makeLocalManagerChildren() map[string][]FormFieldType {
	locallyManagedLightHourseFields := makeConsensusLocalManagerClientField("Lighthouse")
	locallyManagedNimbusFields := makeConsensusLocalManagerClientField("Nimbus")
	locallyManagedPrysmFields := makeConsensusLocalManagerClientField("Prysm")
	locallyManagedTekuFields := makeConsensusLocalManagerClientField("Teku")
	locallyManagedLodestarFields := makeConsensusLocalManagerClientField("Lodestar")
	return map[string][]FormFieldType{
		"Lighthouse": locallyManagedLightHourseFields,
		"Nimbus":     locallyManagedNimbusFields,
		"Prysm":      locallyManagedPrysmFields,
		"Teku":       locallyManagedTekuFields,
		"Lodestar":   locallyManagedLodestarFields,
	}
}

func makeExternalChildren() map[string][]FormFieldType {
	locallyManagedLightHourseFields := makeConsensusExternalField("Lighthouse")
	locallyManagedPrysmFields := makeConsensusExternalField("Prysm")
	locallyManagedTekuFields := makeConsensusExternalField("Teku")
	locallyManagedLodestarFields := makeConsensusExternalField("Lodestar")
	locallyManagedNimbusFields := makeConsensusExternalField("Nimbus")

	return map[string][]FormFieldType{
		"Lighthouse": locallyManagedLightHourseFields,
		"Nimbus":     locallyManagedNimbusFields,
		"Prysm":      locallyManagedPrysmFields,
		"Teku":       locallyManagedTekuFields,
		"Lodestar":   locallyManagedLodestarFields,
	}
}
