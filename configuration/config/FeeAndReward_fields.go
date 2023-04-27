package config

import "github.com/stader-labs/ethcli-ui/configuration/utils"

func init() {
	FieldKey := GetFieldKey()

	ConfigurationFields[Categories.Option.FeeAndReward] = []FormFieldType{
		{
			Label: "Maximum Fee Limit (in gwei)",
			Key:   FieldKey.Fr_max_fee_limit,
			Type:  "int",
			Description: utils.AddNewLines(`Maximum Fee Limit

This value will represent the highest fee you are willing to pay and applies to both manual and automated transactions. Enable this option to set a maximum fee value that will be utilized for all transactions processed by the Stader Node. Note that the maximum fee limit you set will include the priority fee.
If you set the value to 0, the ongoing estimated max fee based on the present state of the network will be displayed, and you will need to state the fee every time you conduct a transaction.
If you set it to a number other than 0, then the estimated max fee by the network will be overridden and the number entered by you will be directly used.
Default: 0`, descriptionSidebarWidth),
		},
		{
			Label: "Priority Fee (in gwei)",
			Key:   FieldKey.Fr_priority_fee,
			Type:  "int",
			Description: utils.AddNewLines(`Priority Fee (in gwei)

This value represents the additional amount you are willing to pay for your transaction beyond the network's essential transaction fee. The higher the value, the faster the block inclusion, as you will provide more ETH to validators for incorporating your transaction.
It is important to note that you must set a Priority Fee higher than zero and ensure that the "Maximum Fee Limit" you have set is enough to match the current estimated transaction fee on the network.
Default: 2`, descriptionSidebarWidth),
		},
		{
			Label: "Archive-mode EC URL",
			Key:   FieldKey.Fr_archive_mode_ec_url,
			Type:  "text",
			Description: utils.AddNewLines(`Archive-mode EC URL

To conserve your disk capacity, the Archive mode is disabled on your primary and fallback execution clients. If you need to generate reward tree files for block intervals that occurred far in the past, you can input the URL of an execution client that has Archive accessibility. To get a free, lightweight client with Archive accessibility, please follow this link: https://www.alchemy.com/supernode.`, descriptionSidebarWidth),
		},
	}
}
