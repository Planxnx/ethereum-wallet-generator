package repository

import (
	"sync"

	"github.com/planxnx/ethereum-wallet-generator/pkg/wallets"
)

type InMemoryRepository struct {
	walletsMu sync.Mutex
	wallets   []*wallets.Wallet
}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		wallets: make([]*wallets.Wallet, 0),
	}
}

func (r *InMemoryRepository) Insert(wallet *wallets.Wallet) error {
	r.walletsMu.Lock()
	defer r.walletsMu.Unlock()
	r.wallets = append(r.wallets, wallet)
	return nil
}

func (r *InMemoryRepository) Result() []*wallets.Wallet {
	return r.wallets
}

func (r *InMemoryRepository) Close() error {
	return nil
}
