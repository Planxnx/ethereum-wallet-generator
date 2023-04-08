package wallets

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/pkg/errors"
	"github.com/planxnx/ethereum-wallet-generator/bip39"
)

var (
	// DefaultBaseDerivationPath is the base path from which custom derivation endpoints
	// are incremented. As such, the first account will be at m/44'/60'/0'/0, the second
	// at m/44'/60'/0'/1, etc
	DefaultBaseDerivationPath       = accounts.DefaultBaseDerivationPath
	DefaultBaseDerivationPathString = DefaultBaseDerivationPath.String()
)

// NewGeneratorMnemonic returns a new wallet generator that uses a mnemonic(BIP39) and a derivation path(BIP44) to generate a wallet.
func NewGeneratorMnemonic(bitSize int) Generator {
	return func() (*Wallet, error) {
		mnemonic, err := NewMnemonic(bitSize)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to generate mnemonic with bit size: %d", bitSize)
		}

		privateKey, err := deriveWallet(bip39.NewSeed(mnemonic, ""), DefaultBaseDerivationPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to derive wallet from mnemonic: %s", mnemonic)
		}

		wallet, err := NewFromPrivatekey(privateKey)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create wallet from private key: %s", privateKey)
		}

		wallet.Bits = bitSize
		wallet.Mnemonic = mnemonic
		wallet.HDPath = DefaultBaseDerivationPathString

		return wallet, nil
	}
}

// NewGeneratorFastMnemonic returns a new wallet generator that uses a mnemonic(BIP39 with cumlative entropy) and a derivation path(BIP44) to generate a wallet.
func NewGeneratorFastMnemonic(bitSize int) Generator {
	fastEntropy, err := bip39.NewFastEntropy(bitSize)
	if err != nil {
		panic(err)
	}

	return func() (*Wallet, error) {
		entropy, err := fastEntropy.Next()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get next entropy")
		}

		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to generate mnemonic from given entropy, entropy: %x", entropy)
		}

		privateKey, err := deriveWallet(bip39.NewSeed(mnemonic, ""), DefaultBaseDerivationPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to derive wallet from mnemonic: %s", mnemonic)
		}

		wallet, err := NewFromPrivatekey(privateKey)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create wallet from private key: %s", privateKey)
		}

		wallet.Bits = bitSize
		wallet.Mnemonic = mnemonic
		wallet.HDPath = DefaultBaseDerivationPathString

		return wallet, nil
	}
}

// NewMnemonic returns a new mnemonic(BIP39) with the given bit size.
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

func deriveWallet(seed []byte, path accounts.DerivationPath) (*ecdsa.PrivateKey, error) {
	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, n := range path {
		key, err = key.Derive(n)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return privateKeyECDSA, nil
}
