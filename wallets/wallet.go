package wallets

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

type (
	// Generator is a function that generates a wallet.
	Generator func() (*Wallet, error)

	// Wallet is a struct that contains the information of a wallet.
	Wallet struct {
		Address    string
		PrivateKey string
		Mnemonic   string
		HDPath     string
		gorm.Model
		Bits int
	}
)

const (
	// DefaultMnemonicBits is the default number of bits to use when generating a mnemonic. default is 128 bits (12 words).
	DefaultMnemonicBits = 128
)

// DefaultGenerator is the default wallet generator.
var DefaultGenerator = NewGeneratorMnemonic(DefaultMnemonicBits)

// NewWallet returns a new wallet using the default generator.
func NewWallet() (*Wallet, error) {
	return DefaultGenerator()
}

// NewFromPrivatekey returns a new wallet from a given private key.
func NewFromPrivatekey(privateKey *ecdsa.PrivateKey) (*Wallet, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	publicKey := &privateKey.PublicKey

	// toString PrivateKey
	priveKeyBytes := crypto.FromECDSA(privateKey)
	privHex := make([]byte, len(priveKeyBytes)*2)
	hex.Encode(privHex, priveKeyBytes)
	privString := b2s(privHex)

	// toString PublicKey
	publicKeyBytes := crypto.Keccak256(crypto.FromECDSAPub(publicKey)[1:])[12:]
	if len(publicKeyBytes) > common.AddressLength {
		publicKeyBytes = publicKeyBytes[len(publicKeyBytes)-common.AddressLength:]
	}
	pubHex := make([]byte, len(publicKeyBytes)*2+2)
	copy(pubHex[:2], "0x")
	hex.Encode(pubHex[2:], publicKeyBytes)
	pubString := b2s(pubHex)

	return &Wallet{
		Address:    pubString,
		PrivateKey: privString,
	}, nil
}
