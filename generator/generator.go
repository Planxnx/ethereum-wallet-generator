package generator

import (
	"strings"
	"sync"
	"time"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//Wallet ethereum wallet data
type Wallet struct {
	Address    string
	PrivateKey string
	Mnemonic   string
	Bits       int
	HDPath     string
	CreatedAt  time.Time
	gorm.Model
}

//NewWallet .
func NewWalletGenerator(bits int, hdPath string) (*Wallet, error) {
	if bits != 256 && bits != 128 {
		return nil, errors.New("bits must be 128 or 256")
	}
	if hdPath == "" {
		hdPath = hdwallet.DefaultBaseDerivationPath.String()
	} else {
		_, err := hdwallet.ParseDerivationPath(hdPath)
		if err != nil {
			return nil, errors.Wrap(err, "wrong hdpath format")
		}
	}
	return &Wallet{
		Bits:   bits,
		HDPath: hdPath,
	}, nil
}

func (w *Wallet) GenerateGeneratorWallet() *Wallet {
	mnemonic, _ := hdwallet.NewMnemonic(w.Bits)
	w.Mnemonic = mnemonic

	wallet, _ := hdwallet.NewFromMnemonic(w.Mnemonic)

	path := hdwallet.DefaultBaseDerivationPath
	if w.HDPath != "" {
		path = hdwallet.MustParseDerivationPath(w.HDPath)
	}

	account, _ := wallet.Derive(path, false)
	pk, _ := wallet.PrivateKeyHex(account)

	w.Address = account.Address.Hex()
	w.PrivateKey = pk
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()

	return w
}

func (w *Wallet) GenerateWalletFromMnemonic(mnemonic string) *Wallet {
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	path := hdwallet.DefaultBaseDerivationPath
	if w.HDPath != "" {
		path = hdwallet.MustParseDerivationPath(w.HDPath)
	}

	account, _ := wallet.Derive(path, false)
	pk, _ := wallet.PrivateKeyHex(account)

	return &Wallet{
		Address:    account.Address.Hex(),
		PrivateKey: pk,
		Mnemonic:   mnemonic,
		HDPath:     account.URL.Path,
		CreatedAt:  time.Now(),
	}
}

func (w *Wallet) GenerateWallet() *Wallet {
	mnemonic, _ := hdwallet.NewMnemonic(w.Bits)
	return w.GenerateWalletFromMnemonic(mnemonic)
}

type Config struct {
	Number      int
	Concurrency int
	Bits        int
	Strict      bool
	Contains    []string
	Prefix      string
	Suffix      string
}

func (w *Wallet) GenerateMultipleWallets(cf Config) (index chan int, wallets chan *Wallet) {
	if cf.Concurrency < 1 {
		cf.Concurrency = 1
	}

	addreessValidator := func(address string) bool {
		isValid := false
		if len(cf.Contains) != 0 {
			cb := func(contain string) bool {
				return strings.Contains(address, contain)
			}
			if cf.Strict {
				if have(cf.Contains, cb) {
					isValid = true
				}
			} else {
				if some(cf.Contains, cb) {
					isValid = true
				}
			}
		}

		if cf.Prefix != "" {
			if !strings.HasPrefix(address, cf.Prefix) {
				isValid = false
			}
		}

		if cf.Suffix != "" {
			if !strings.HasSuffix(address, cf.Suffix) {
				isValid = false
			}
		}

		return isValid
	}
	if cf.Number < 0 {
		cf.Number = -1
	}

	var wg sync.WaitGroup
	index = make(chan int, cf.Number)
	wallets = make(chan *Wallet, cf.Number)

	go func() {
		defer func() {
			close(wallets)
			close(index)
		}()
		for i := 0; i < cf.Number || cf.Number < 0; i += cf.Concurrency {
			for j := 0; j < cf.Concurrency && (i+j < cf.Number || cf.Number < 0); j++ {
				wg.Add(1)

				go func(j int) {
					defer wg.Done()

					wallet := w.GenerateWallet()

					index <- i + j

					if !addreessValidator(wallet.Address) {
						return
					}

					wallets <- wallet
				}(j)
			}
			wg.Wait()
		}
	}()

	return index, wallets
}

// forked this methods from core-js
func some(arr []string, fn func(string) bool) bool {
	for _, v := range arr {
		if fn(v) {
			return true
		}
	}
	return false
}

func have(arr []string, fn func(string) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}
