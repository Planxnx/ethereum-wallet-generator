package repository

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/planxnx/ethereum-wallet-generator/pkg/wallets"
	"gorm.io/gorm"
)

type GormRepository struct {
	db        *gorm.DB
	mu        sync.Mutex
	tx        *gorm.DB
	txSize    uint64
	maxTxSize uint64
}

func NewGormRepository(db *gorm.DB, maxTxSize uint64) Repository {
	return &GormRepository{
		db:        db,
		maxTxSize: maxTxSize,
	}
}

func (r *GormRepository) Insert(wallet *wallets.Wallet) error {
	if r.db == nil {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.tx == nil {
		r.tx = r.db.Begin()
	}

	if err := r.tx.Create(wallet).Error; err != nil {
		return errors.WithStack(err)
	}
	r.txSize++

	if r.txSize >= r.maxTxSize {
		if err := r.commit(); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (r *GormRepository) Result() []*wallets.Wallet {
	return nil
}

func (r *GormRepository) Close() error {
	if err := r.Commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *GormRepository) Commit() error {
	if r.db == nil {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// commit unsafe method, should be called inside package only
func (r *GormRepository) commit() error {
	if r.tx != nil {
		if err := r.tx.Commit().Error; err != nil {
			return errors.WithStack(err)
		}
		r.txSize = 0
		r.tx = nil
	}
	return nil
}
