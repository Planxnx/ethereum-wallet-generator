package wallets

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func NewGeneratorPrivatekey() Generator {
	return func() (*Wallet, error) {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		wallet, err := NewFromPrivatekey(privateKey)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return wallet, nil
	}
}
