package main

import (
	"encoding/json"
	"fmt"

	"github.com/stader-labs/ethcli-ui/ui"
)

func main() {
	a, wizard, configuration := ui.Run(nil, nil)
	fmt.Printf("saved: %v\n\n", a)

	// Wizard json
	ws, _ := json.MarshalIndent(wizard, "", "  ")
	fmt.Printf("wizard: %s\n\n", string(ws))

	// Configuration json
	cs, _ := json.MarshalIndent(configuration, "", "  ")
	fmt.Printf("configuration: %s\n\n", string(cs))
}
