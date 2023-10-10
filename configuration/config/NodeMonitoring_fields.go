package config

import "github.com/stader-labs/ethcli-ui/configuration/utils"

func init() {
	FieldKey := GetFieldKey()

	metricsFields := []FormFieldType{
		{
			Label: "Expose Metrics Port",
			Key:   FieldKey.Nm_expose_guardian_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Metrics

Please choose this option to expose Stader Node guardian.
Default: True`, descriptionSidebarWidth),
		},
		{
			Label: "Execution Client Metrics Port",
			Key:   FieldKey.Nm_execution_client_metrics_port,
			Type:  "int",
			Description: utils.AddNewLines(`Execution Client Metrics Port

			Enter the port your ETH1- Execution Client should use for the API.
Default: 9105`, descriptionSidebarWidth),
		},
		{
			Label: "Beacon Node Metrics Port",
			Key:   FieldKey.Nm_beacon_node_metrics_port,
			Type:  "int",
			Description: utils.AddNewLines(`Beacon Node Metrics Port

Enter the port your Beacon node should use for the API.
Default: 9100`, descriptionSidebarWidth),
		},
		{
			Label: "Validator Client Metrics Port",
			Key:   FieldKey.Nm_validator_client_metrics_port,
			Type:  "int",
			Description: utils.AddNewLines(`Validator Client Metrics Port

Enter the port your Validator Client should use for the API.
Default: 9101`, descriptionSidebarWidth),
		},
		{
			Label: "Node Metrics Port",
			Key:   FieldKey.Nm_node_metrics_port,
			Type:  "int",
			Description: utils.AddNewLines(`Node Metrics Port

			Enter the port your Node Container should use for the API.
Default: 9102`, descriptionSidebarWidth),
		},
		{
			Label: "Exporter Metrics Port",
			Key:   FieldKey.Nm_exporter_metrics_port,
			Type:  "int",
			Description: utils.AddNewLines(`Exporter Metrics Port

Enter the port your Node Exporter should use for the API.
Default: 9103`, descriptionSidebarWidth),
		},
		{
			Label: "Guardian Oracle Port",
			Key:   FieldKey.Nm_guardian_oracle_port,
			Type:  "int",
			Description: utils.AddNewLines(`Guardian Oracle Port

Enter the port your Guardian Oracle Container should use for the API.
Default: 9104`, descriptionSidebarWidth),
		},
		{
			Label: "Grafana Port",
			Key:   FieldKey.Nm_grafana_port,
			Type:  "int",
			Description: utils.AddNewLines(`Grafana Port

Enter the port which Grafana will use to host its HTTP server; this is also the port you will be using in your browser for access.
Default: 3100`, descriptionSidebarWidth),
		},
		{
			Label: "Grafana Container Tag",
			Key:   FieldKey.Nm_grafana_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Grafana Container Tag

Please enter the label you wish to use on Docker Hub for the Grafana container.
Default: grafana/grafana:9.3.1`, descriptionSidebarWidth),
		},
		{
			Label: "Prometheus Port",
			Key:   FieldKey.Nm_prometheus_port,
			Type:  "int",
			Description: utils.AddNewLines(`Prometheus Port

Enter the port where Promethus metrics will be presented to you.
Default: 9091`, descriptionSidebarWidth),
		},
		{
			Label: "Expose Prometheus Port",
			Key:   FieldKey.Nm_expose_prometheus_port,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Expose Prometheus Port

Choose this option to make the ports visible on your local network, granting access to your Promethus ports for other local devices.
Default: false`, descriptionSidebarWidth),
		},
		{
			Label: "Prometheus Container Tag",
			Key:   FieldKey.Nm_prometheus_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Prometheus Container Tag

Please enter the label you wish to use on Docker Hub for the Promethus container.`, descriptionSidebarWidth),
		},
		{
			Label: "Additional Prometheus Flags",
			Key:   FieldKey.Nm_additional_prometheus_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Prometheus Flags

Please enter other flags you might use in conjunction with Promethus to activate added settings that the Stader Node configuration overlooks.`, descriptionSidebarWidth),
		},
		{
			Label: "Allow Root Filesystem Access",
			Key:   FieldKey.Nm_allow_root_filesystem_access,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Allow Root Filesystem Access

Choose this option to alllow Prometheus's Node Exporter to view your root filesystem. This is necessary if you want the Grafana dashboard to display the utilized disk space of your SSD.`, descriptionSidebarWidth),
		},
		{
			Label: "Exporter Container Tag",
			Key:   FieldKey.Nm_exporter_container_tag,
			Type:  "text",
			Description: utils.AddNewLines(`Exporter Container Tag

Please enter the label you wish to use on Docker Hub for the Promethus node exporter container.`, descriptionSidebarWidth),
		},
		{
			Label: "Additional Exporter Flags",
			Key:   FieldKey.Nm_additional_exporter_flags,
			Type:  "text",
			Description: utils.AddNewLines(`Additional Exporter Flags

Please enter other custom command line flags you might use in conjunction with Node Exporter to activate added settings that the Stader Node configuration overlooks.`, descriptionSidebarWidth),
		},
	}

	ConfigurationFields[Categories.Option.NodeMonitoring] = []FormFieldType{
		{
			Label: "Enable Metrics",
			Key:   FieldKey.Nm_enable_metrics,
			Type:  "checkbox",
			Description: utils.AddNewLines(`Enable Metrics

Please choose this option to activate Stader Node Metric Monitoring.
Default: True`, descriptionSidebarWidth),

			Children: map[string][]FormFieldType{
				"true": metricsFields,
				"false": {{
					Label: "Expose Metrics Port",
					Key:   FieldKey.Nm_expose_guardian_port,
					Type:  "checkbox",
					Description: utils.AddNewLines(`Enable Metrics
		
		Please choose this option to expose Stader Node guardian.
		Default: True`, descriptionSidebarWidth),
				}, {
					Label: "Node Metrics Port",
					Key:   FieldKey.Nm_node_metrics_port,
					Type:  "int",
					Description: utils.AddNewLines(`Node Metrics Port
		
					Enter the port your Node Container should use for the API.
		Default: 9102`, descriptionSidebarWidth),
				}},
			},
		},
	}
}
