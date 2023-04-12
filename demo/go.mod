module github.com/stader-labs/ethcli-ui/demo

go 1.19

require github.com/stader-labs/ethcli-ui/ui v0.0.0-00010101000000-000000000000

require (
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.6.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/rivo/tview v0.0.0-20230406072732-e22ce9588bb4 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/stader-labs/ethcli-ui/configuration v0.0.0-20230412104132-f4eecaffdded // indirect
	github.com/stader-labs/ethcli-ui/wizard v0.0.0-20230412104132-f4eecaffdded // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/term v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

replace github.com/stader-labs/ethcli-ui/ui => ../ui

replace github.com/stader-labs/ethcli-ui/wizard => ../wizard

replace github.com/stader-labs/ethcli-ui/configuration => ../configuration

replace github.com/rivo/tview => github.com/hamidraza/tview v0.0.0-20230406022610-62c55f74c35f
