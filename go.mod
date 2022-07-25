module github.com/Planxnx/eth-wallet-gen

go 1.15

require (
	eth-wallet-gen/common v0.0.0-00010101000000-000000000000
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/ethereum/go-ethereum v1.10.4
	github.com/fatih/color v1.12.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/pkg/errors v0.9.1
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.10
)

replace eth-wallet-gen/common => ./common
