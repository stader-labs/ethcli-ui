package config

import "github.com/stader-labs/ethcli-ui/configuration/utils"

func init() {
	FieldKey := GetFieldKey()

	locallyManagedGethFields := []FormFieldType{
		{
			Label: "HTTP Port",
			Key:   FieldKey.E1ec_lm_common_http_port,
			Type:  "int",
			Description: utils.AddNewLines(`HTTP Port
		
Enter the port your HTTP RPC API should use
Default: 8545`, descriptionSidebarWidth),
		},
		{
			Label: "Websocket Port",
			Key:   FieldKey.E1ec_lm_common_websocket_port,
			Type:  "int",
			Description: utils.AddNewLines(`Websocket Port

Enter the port your Websocket RPC API should use
Default: 8546`, descriptionSidebarWidth),
		},
		{
			Label: "Engine API port",
			Key:   FieldKey.E1ec_lm_common_engine_api_port,
			Type:  "int",
			Description: utils.AddNewLines(`Engine API port

Enter the port your Engine API should use. After the merge, the Consensus client will establish a connection with this endpoint.
Default: 8551`, descriptionSidebarWidth),
		},
		{
			Label: "Expose RPC port",
			Key:   FieldKey.E1ec_lm_common_expose_rpc_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose RPC Ports

Choose this option to make the HTTP and Websocket RPC ports visible on your local config.Network, granting access to your Execution Client's RPC endpoint for other local devices.
Default: False`, descriptionSidebarWidth),
		},
		{
			Label: "P2P port",
			Key:   FieldKey.E1ec_lm_common_p2p_port,
			Type:  "int",
			Description: utils.AddNewLines(`P2P port

Enter the port number that the Execution Client will use for P2P communication with other blockchain nodes.
Default: 30303`, descriptionSidebarWidth),
		},
		{
			Label: "ETHStats Label",
			Key:   FieldKey.E1ec_lm_common_ethstats_label,
			Type:  "text",
			Description: utils.AddNewLines(`ETHStats Label

Please provide the label of your choice if you wish to share your Execution Client's statistics on https://ethstats.net/
Default: ""`, descriptionSidebarWidth),
		},
		{
			Label: "ETHStats Login",
			Key:   FieldKey.E1ec_lm_common_ethstats_login,
			Type:  "text",
			Description: utils.AddNewLines(`ETHStats Login

Please provide a login you plan to use if you wish to share your Execution Client's statistics on https://ethstats.net/
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "Cache Size (in MB)",
			Key:   FieldKey.E1ec_lm_geth_cache_size,
			Type:  "int",
			Description: utils.AddNewLines(`Cache Size (in MB)

Specify the desired cache size for your ETH1 - Execution Client. The present value is determined by the overall amount of RAM on your system, but you have the option to modify it manually. Note that a higher cache size will result in a slower rise in disk space usage and will reduce the frequency at which pruning is required.
Default: 256`, descriptionSidebarWidth),
		},
		{
			Label: "Max Peers",
			Key:   FieldKey.E1ec_lm_geth_max_peers,
			Type:  "int",
			Description: utils.AddNewLines(`Max Peers

Please specify your ETH1 - Execution client peer limit. The number can be decreased to enhance performance on systems with limited specs. However, We suggest you to keep it at at least 12.
Default: 25`, descriptionSidebarWidth),
		},
		{
			Label: "Container tag",
			Key:   FieldKey.E1ec_lm_geth_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Container tag

Please specify the name of your tag you intend to utilize for the ETH1 - Execution Client container
Default: ethereum/client-go:v1.10.26`, descriptionSidebarWidth),
		},
		{
			Label: "Additional flags",
			Key:   FieldKey.E1ec_lm_geth_additional_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional flags

Are there any extra personalized terminal flags you intend to give to your ETH1 - Execution Client, in order to make use of supplementary configurations that are not covered by the Stader Node's setup?`, descriptionSidebarWidth),
		},
	}

	locallyManagedNethermindFields := []FormFieldType{
		{
			Label: "HTTP Port",
			Key:   FieldKey.E1ec_lm_common_http_port,
			Type:  "int",
			Description: utils.AddNewLines(`HTTP Port

Enter the port your HTTP RPC API should use
Default: 8545`, descriptionSidebarWidth),
		},
		{
			Label: "Websocket Port",
			Key:   FieldKey.E1ec_lm_common_websocket_port,
			Type:  "int",
			Description: utils.AddNewLines(`Websocket Port

Enter the port your Websocket RPC API should use
Default: 8546`, descriptionSidebarWidth),
		},
		{
			Label: "Engine API port",
			Key:   FieldKey.E1ec_lm_common_engine_api_port,
			Type:  "int",
			Description: utils.AddNewLines(`Engine API port

Enter the port your Engine API should use. After the merge, the Consensus client will establish a connection with this endpoint.
Default: 8551`, descriptionSidebarWidth),
		},
		{
			Label: "Expose RPC port",
			Key:   FieldKey.E1ec_lm_common_expose_rpc_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose RPC Ports

Choose this option to make the HTTP and Websocket RPC ports visible on your local config.Network, granting access to your Execution Client's RPC endpoint for other local devices.
Default: False`, descriptionSidebarWidth),
		},
		{
			Label: "P2P port",
			Key:   FieldKey.E1ec_lm_common_p2p_port,
			Type:  "int",
			Description: utils.AddNewLines(`P2P port

Enter the port number that the Execution Client will use for P2P communication with other blockchain nodes.
Default: 30303`, descriptionSidebarWidth),
		},
		{
			Label: "ETHStats Label",
			Key:   FieldKey.E1ec_lm_common_ethstats_label,
			Type:  "text",
			Description: utils.AddNewLines(`ETHStats Label

Please provide the label of your choice if you wish to share your Execution Client's statistics on https://ethstats.net/
Default: 8546`, descriptionSidebarWidth),
		},
		{
			Label: "ETHStats Login",
			Key:   FieldKey.E1ec_lm_common_ethstats_login,
			Type:  "text",
			Description: utils.AddNewLines(`ETHStats Login

Please provide a login you plan to use if you wish to share your Execution Client's statistics on https://ethstats.net/
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "Cache Size (in MB)",
			Key:   FieldKey.E1ec_lm_nethermind_cache_size,
			Type:  "int",
			Description: utils.AddNewLines(`Cache Size (in MB)

Specify the desired cache size for your ETH1 - Execution Client. The present value is determined by the overall amount of RAM on your system, but you have the option to modify it manually. Note that a higher cache size will result in a slower rise in disk space usage and will reduce the frequency at which pruning is required.
Default: 256`, descriptionSidebarWidth),
		},
		{
			Label: "Max Peers",
			Key:   FieldKey.E1ec_lm_nethermind_max_peers,
			Type:  "int",
			Description: utils.AddNewLines(`Max Peers

Please specify your ETH1 - Execution client peer limit. The number can be decreased to enhance performance on systems with limited specs. However, We suggest you to keep it at at least 12.
Default: 25`, descriptionSidebarWidth),
		},
		{
			Label: "In-Memory Pruning Cache Size",
			Key:   FieldKey.E1ec_lm_nethermind_pruning_cache_size,
			Type:  "int",
			Description: utils.AddNewLines(`In-Memory Pruning Cache Size

Specify the desired RAM you wish to allocate for your ETH1 - Execution Client for its in-memory pruning mechanism. The present value is determined by the overall amount of RAM on your system, but you have the option to modify it manually. Note that a larger size imply less writing to your SSD and gradual growth of the database as a whole.
Default: 512`, descriptionSidebarWidth),
		},
		{
			Label: "Additional Modules",
			Key:   FieldKey.E1ec_lm_nethermind_additional_modules,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Modules

Include extra modules to the primary JSON-RPC route beyond the default ones (Personal,Net,Eth,Web3). Add any other required modules by separating them with commas and not using spaces.
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "Additional URLs",
			Key:   FieldKey.E1ec_lm_nethermind_additional_urls,
			Type:  "text",
			Description: utils.AddNewLines(`Additional URLs

To run additional JSON-RPC URLs with the primary URL, place each extra URL in quotes by separating them with commas and not using spaces. The URLs will get added to the "--JsonRpc.AdditionalRpcUrls" argument. Please note that this is meant for experienced users only and for more details regarding this flag, its usage, and formatting, refer to the ETH Client documentation.
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "Container tag",
			Key:   FieldKey.E1ec_lm_nethermind_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Container tag

Please specify the name of your tag you intend to utilize for the ETH1 - Execution Client container
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "Additional flags",
			Key:   FieldKey.E1ec_lm_nethermind_additional_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional flags

Are there any extra personalized terminal flags you intend to give to your ETH1 - Execution Client, in order to make use of supplementary configurations that are not covered by the Stader Node's setup?
Default:`, descriptionSidebarWidth),
		},
	}

	locallyManagedBesuFields := []FormFieldType{
		{
			Label: "HTTP Port",
			Key:   FieldKey.E1ec_lm_common_http_port,
			Type:  "int",
			Description: utils.AddNewLines(`HTTP Port

Enter the port your HTTP RPC API should use
Default: 8545`, descriptionSidebarWidth),
		},
		{
			Label: "Websocket port",
			Key:   FieldKey.E1ec_lm_common_websocket_port,
			Type:  "int",
			Description: utils.AddNewLines(`Websocket port

Enter the port your Websocket RPC API should use
Default: 8546`, descriptionSidebarWidth),
		},
		{
			Label: "Engine API port",
			Key:   FieldKey.E1ec_lm_common_engine_api_port,
			Type:  "int",
			Description: utils.AddNewLines(`Engine API port

Enter the port your Engine API should use. After the merge, the Consensus client will establish a connection with this endpoint.
Default: 8551`, descriptionSidebarWidth),
		},
		{
			Label: "Expose RPC ports",
			Key:   FieldKey.E1ec_lm_common_expose_rpc_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose RPC ports

Choose this option to make the HTTP and Websocket RPC ports visible on your local config.Network, granting access to your Execution Client's RPC endpoint for other local devices.
Default: False`, descriptionSidebarWidth),
		},
		{
			Label: "P2P port",
			Key:   FieldKey.E1ec_lm_common_p2p_port,
			Type:  "int",
			Description: utils.AddNewLines(`P2P port

Enter the port number that the Execution Client will use for P2P communication with other blockchain nodes.
Default: 30303`, descriptionSidebarWidth),
		},
		{
			Label: "ETHStats Label",
			Key:   FieldKey.E1ec_lm_common_ethstats_label,
			Type:  "text",
			Description: utils.AddNewLines(`Ethstats

Please provide the label of your choice if you wish to share your Execution Client's statistics on https://ethstats.net/
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "ETHStats Login",
			Key:   FieldKey.E1ec_lm_common_ethstats_login,
			Type:  "text",
			Description: utils.AddNewLines(`ETHStats Login

Please provide a login you plan to use if you wish to share your Execution Client's statistics on https://ethstats.net/
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "JVM Heap Size",
			Key:   FieldKey.E1ec_lm_besu_jvm_heap_size,
			Type:  "int",
			Description: utils.AddNewLines(`JVM Heap Size

The upper threshold of the RAM, expressed in MB, that the JVM in Besu should restrict itself to. Lowering this value would prompt Besu to consume less RAM, although it would inevitably exceed this cap. If you prefer an automatic allocation, input 0.
Default: 0.`, descriptionSidebarWidth),
		},
		{
			Label: "Max peers",
			Key:   FieldKey.E1ec_lm_besu_max_peers,
			Type:  "int",
			Description: utils.AddNewLines(`Max peers

Please specify your ETH1 - Execution client peer limit. The number can be decreased to enhance performance on systems with limited specs. However, We suggest you to keep it at at least 12.
Default: 25.`, descriptionSidebarWidth),
		},
		{
			Label: "In-Memory Pruning Cache Size",
			Key:   FieldKey.E1ec_lm_besu_in_memory_pruning_cache_size,
			Type:  "int",
			Description: utils.AddNewLines(`In-Memory Pruning Cache Size

Please specify your ETH1 - Execution client peer limit. The number can be decreased to enhance performance on systems with limited specs. However, We suggest you to keep it at at least 12.
Default: 25`, descriptionSidebarWidth),
		},
		{
			Label: "Container tag",
			Key:   FieldKey.E1ec_lm_besu_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Container tag

Please specify the name of your tag you intend to utilize for the ETH1 - Execution Client container
Default:`, descriptionSidebarWidth),
		},
		{
			Label: "Additional flags",
			Key:   FieldKey.E1ec_lm_besu_additional_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional flags

Are there any extra personalized terminal flags you intend to give to your ETH1 - Execution Client, in order to make use of supplementary configurations that are not covered by the Stader Node's setup?
Default:`, descriptionSidebarWidth),
		},
	}

	locallyManagedFields := []FormFieldType{
		{
			Label: "Execution client",
			Key:   FieldKey.E1ec_lm_execution_client,
			Type:  "select",
			Options: []string{
				// "Teku",
				// "Lighthouse",
				// "Nimbus",
				// "Prysm",
				"Geth",
				"Nethermind",
				"Besu",
			},
			Description: utils.AddNewLines(`Execution Client

Choose the ETH 1 - Execution Client that you wish to use.`, descriptionSidebarWidth),
			OptionDescriptions: map[string]string{
				// "Teku":       `Teku`,
				// "Lighthouse": `Lighthouse`,
				// "Nimbus":     `Nimbus`,
				// "Prysm":      `Prysm`,
				"Geth": utils.AddNewLines(`Geth

One of the most popular software clients maintained by the Ethereum Foundation, Geth is an open source CLI developed in the Go Programming Language. It is designed to be flexible and customizable, and it supports a wide range of functionalities such as secure key management, consensus mechanisms etc.`, descriptionSidebarWidth),
				"Nethermind": utils.AddNewLines(`Nethermind

Nethermind is a high-performance Ethereum client built on .NET that offers fast sync speeds and advanced features for developers and end-users. While requiring over 8GB of RAM, it remains a reliable choice for running Ethereum nodes.`, descriptionSidebarWidth),
				"Besu": utils.AddNewLines(`Besu

Besu, developed by ConsenSys and written in Java, is a comprehensive Ethereum protocol client. It utilizes an innovative storage system called "Bonsai Trees" to store its chain data effectively, enabling it to retrieve historical block states without the need for pruning.`, descriptionSidebarWidth),
			},
			Children: map[string][]FormFieldType{
				"Geth":       locallyManagedGethFields,
				"Nethermind": locallyManagedNethermindFields,
				"Besu":       locallyManagedBesuFields,
			},
		},
	}

	externallyManagedFields := []FormFieldType{
		{
			Label: "HTTP URL",
			Key:   FieldKey.E1ec_em_http_url,
			Type:  "text",
			Description: utils.AddNewLines(`HTTP URL
			
Enter the HTTP-based RPC API URL for your current external client. Remember to use your machine's LAN IP address in the place of localhost or 127.0.0.1 when running this client on the same machine as the Stader Node.,`, descriptionSidebarWidth),
		},
		{
			Label:       "Websocket URL",
			Key:         FieldKey.E1ec_em_websocket_url,
			Type:        "text",
			Description: `Websocket URL`,
		},
	}

	ConfigurationFields[Categories.Option.ETH1ExecutionClient] = []FormFieldType{
		{
			Label: "Execution client Preference",
			Key:   FieldKey.E1ec_execution_and_consensus_mode,
			Type:  "select",
			Options: []string{
				"Locally Managed",
				"Externally Managed",
			},
			Description: `Execution client Preference

Select your preferred method for
managing your Execution Client.

Default: Locally Managed`,
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
