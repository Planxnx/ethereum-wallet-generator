module github.com/Planxnx/eth-wallet-gen

go 1.20

require (
	github.com/btcsuite/btcd v0.23.1
	github.com/btcsuite/btcd/btcutil v1.1.2
	github.com/cheggaaa/pb/v3 v3.1.0
	github.com/ethereum/go-ethereum v1.10.23
	github.com/pkg/errors v0.9.1
	github.com/schollz/progressbar/v3 v3.9.0
	github.com/tyler-smith/go-bip39 v1.1.0
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.3.4 // indirect
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8 // indirect
	golang.org/x/sys v0.0.0-20220823224334-20c2bfdbfe24 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
)

replace eth-wallet-gen/common => ./common
