module github.com/Planxnx/eth-wallet-gen

go 1.15

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/btcsuite/btcd v0.23.1
	github.com/btcsuite/btcd/btcutil v1.1.1
	github.com/cheggaaa/pb/v3 v3.1.0
	github.com/ethereum/go-ethereum v1.10.20
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-sqlite3 v1.14.14 // indirect
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/schollz/progressbar/v3 v3.8.7
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

replace eth-wallet-gen/common => ./common
