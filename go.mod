module github.com/Planxnx/eth-wallet-gen

go 1.15

require (
	eth-wallet-gen/common v0.0.0-00010101000000-000000000000
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/cheggaaa/pb/v3 v3.0.8
	github.com/ethereum/go-ethereum v1.10.3 // indirect
	github.com/fatih/color v1.12.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/miguelmota/go-ethereum-hdwallet v0.0.1
	github.com/pkg/errors v0.9.1
	github.com/schollz/progressbar/v3 v3.8.3
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.10
)

replace eth-wallet-gen/common => ./common
