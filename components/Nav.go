package components

import (
	"github.com/stader-labs/ethcli-ui/config"

	"github.com/rivo/tview"
)

func Nav(selectedItem string) tview.Primitive {
	navItemNames := []string{
		config.TopNav.Network,
		config.TopNav.ETHClient,
		config.TopNav.ExecutionClient,
		config.TopNav.ConsensusClient,
		config.TopNav.FallbackClients,
		config.TopNav.Monitoring,
		config.TopNav.MEVBoost,
	}

	navItems := tview.NewFlex().
		SetDirection(tview.FlexColumn)

	for _, name := range navItemNames {
		navItem := NavText(name, name == selectedItem).SetTextAlign(tview.AlignCenter)
		navItems.AddItem(navItem, 0, 1, false)
	}

	return navItems
}
