package config

import "github.com/stader-labs/ethcli-ui/configuration/utils"

func init() {
	FieldKey := GetFieldKey()

	ConfigurationFields[Categories.Option.FallbackClients] = []FormFieldType{
		{
			Label: "Use Fallback Clients",
			Key:   FieldKey.Fc_use_fallback_clients,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Use Fallback Clients
			
Choose this option to add fallback client.`, descriptionSidebarWidth),
			Children: map[string][]FormFieldType{
				"true": {
					{
						Label: "Reconnect Delay",
						Key:   FieldKey.Fc_true_reconnect_delay,
						Type:  "text",
						Description: utils.AddNewLines(`Reconnect Delay

Please set a wait time after your primary Execution or Consensus client disconnects before attempting to reconnect.
Note: Please set the time in the given format: 12h50m40s
i.e. 12 hours, 50 minutes and 40 seconds.
Default: 60s`, descriptionSidebarWidth),
					},
					{
						Label: "Execution Client URL",
						Key:   FieldKey.Fc_true_execution_client_url,
						Type:  "text",
						Description: utils.AddNewLines(`Execution Client URL

Please enter the URLs of the HTTP APIs for your fallback clients.
Note: When running this client on the same machine as the Stader Node, use your machine's LAN IP address instead of localhost or 127.0.0.1.
Default:`, descriptionSidebarWidth),
					},
					{
						Label: "Beacon Node URL",
						Key:   FieldKey.Fc_true_beacon_node_url,
						Type:  "text",
						Description: utils.AddNewLines(`Beacon Node URL

Please enter the URLs of the HTTP APIs for your fallback clients.
Note: When running this client on the same machine as the Stader Node, use your machine's LAN IP address instead of localhost or 127.0.0.1.
Default:`, descriptionSidebarWidth),
					},
					{
						Label:       "Beacon Node JSON-RPC URL",
						Key:         FieldKey.Fc_true_beacon_node_json_rpc_url,
						Type:        "text",
						Description: utils.AddNewLines(`Beacon Node JSON-RPC URL`, descriptionSidebarWidth),
					},
				},
			},
		},
	}
}
