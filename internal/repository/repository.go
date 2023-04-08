package repository

import "github.com/planxnx/ethereum-wallet-generator/wallets"

type Repository interface {
	Insert(wallet *wallets.Wallet) error
	Result() []*wallets.Wallet
	Close() error
}
