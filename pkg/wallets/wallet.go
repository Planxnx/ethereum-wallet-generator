package wallets

import (
	"crypto/ecdsa"
	"encoding/hex"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"gorm.io/gorm"
)

var (
	// DefaultBaseDerivationPath is the base path from which custom derivation endpoints
	// are incremented. As such, the first account will be at m/44'/60'/0'/0, the second
	// at m/44'/60'/0'/1, etc
	DefaultBaseDerivationPath       = accounts.DefaultBaseDerivationPath
	DefaultBaseDerivationPathString = DefaultBaseDerivationPath.String()
)

type Wallet struct {
	Address    string
	PrivateKey string
	Mnemonic   string
	Bits       int
	HDPath     string
	CreatedAt  time.Time
	gorm.Model
}

func NewWallet(bitSize int) (*Wallet, error) {
	mnemonic, err := NewMnemonic(bitSize)
	if err != nil {
		return nil, err
	}

	privateKey, publicKey, err := deriveWallet(bip39.NewSeed(mnemonic, ""), DefaultBaseDerivationPath)
	if err != nil {
		return nil, err
	}

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
		Mnemonic:   mnemonic,
		Bits:       bitSize,
		HDPath:     DefaultBaseDerivationPathString,
		CreatedAt:  time.Now(),
	}, nil
}

func deriveWallet(seed []byte, path accounts.DerivationPath) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, err
	}

	for _, n := range path {
		key, err = key.DeriveNonStandard(n)
		if err != nil {
			return nil, nil, err

		}
	}

	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, nil, err
	}

	return privateKeyECDSA, privateKeyECDSA.Public().(*ecdsa.PublicKey), nil
}
