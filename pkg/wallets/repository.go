package wallets

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DefaultRepositoryConfig = RepositoryConfig{
	BitSize: 128,
	AddresValidator: func(address string) bool {
		return true
	},
	DBTransactionsSize: 100,
}

var (
	ErrorInvalidAddress = errors.New("invalid wallet address")
	ErrorCanNotGenerate = errors.New("can not generate wallet")
	ErrorCanNotStore    = errors.New("can not store generated wallet")
)

type Repository struct {
	// StdOut
	result strings.Builder

	// DB
	mu        sync.Mutex
	tx        *gorm.DB
	txWallets uint64

	config RepositoryConfig
}

type RepositoryConfig struct {
	AddresValidator    func(address string) bool
	DB                 *gorm.DB
	DBTransactionsSize uint64
	BitSize            int
}

func NewRepository(config ...RepositoryConfig) *Repository {
	r := &Repository{
		config: DefaultRepositoryConfig,
	}

	if len(config) > 0 {
		r.config = config[0]
	}

	if r.config.BitSize != 128 && r.config.BitSize != 256 {
		r.config.BitSize = 128
	}

	if r.config.AddresValidator == nil {
		r.config.AddresValidator = DefaultRepositoryConfig.AddresValidator
	}

	if r.config.DBTransactionsSize == 0 {
		r.config.DBTransactionsSize = DefaultRepositoryConfig.DBTransactionsSize
	}

	return r
}

func (r *Repository) Generate() (ok bool, err error) {
	wallet, err := NewWallet(r.config.BitSize)
	if err != nil {
		return false, errors.Wrapf(ErrorCanNotGenerate, err.Error())
	}

	if !r.config.AddresValidator(wallet.Address) {
		return false, nil
	}

	if r.config.DB != nil {
		if err := r.dbInsert(wallet); err != nil {
			return false, errors.Wrapf(ErrorCanNotStore, err.Error())
		}
		return true, nil
	}

	if _, err := fmt.Fprintf(&r.result, "%-18s %s\n", wallet.Address, wallet.Mnemonic); err != nil {
		return false, errors.Wrapf(ErrorCanNotStore, err.Error())
	}

	return true, nil
}

func (r *Repository) Close() error {
	if err := r.Commit(); err != nil {
		return errors.WithStack(err)
	}
	r.result.Reset()
	return nil
}

func (r *Repository) Commit() error {
	if r.config.DB == nil {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *Repository) Results() string {
	if r.result.Len() > 0 {
		return r.result.String()
	}
	return ""
}

func (r *Repository) dbInsert(wallet *Wallet) error {
	if r.config.DB == nil {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.tx == nil {
		r.tx = r.config.DB.Begin()
	}

	if err := r.tx.Create(wallet).Error; err != nil {
		return errors.WithStack(err)
	}
	r.txWallets++

	if r.txWallets >= r.config.DBTransactionsSize {
		if err := r.commit(); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// commit unsafe method, should be called inside package only
func (r *Repository) commit() error {
	if r.config.DB == nil {
		return nil
	}

	if r.tx != nil {
		if err := r.tx.Commit().Error; err != nil {
			return errors.WithStack(err)
		}
		r.txWallets = 0
		r.tx = nil
	}

	return nil
}
