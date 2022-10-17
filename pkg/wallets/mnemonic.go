package wallets

import (
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
)

func NewMnemonic(bitSize int) (string, error) {
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", errors.WithStack(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return mnemonic, nil
}
