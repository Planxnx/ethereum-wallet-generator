package generator

import (
	"testing"
)

func TestNewGeneratorWithEmptyPath(t *testing.T) {
	bits := 256
	hdpath := ""
	generator, err := NewWalletGenerator(bits, hdpath)
	if err != nil {
		t.Errorf("Error with %v", err)
	}
	if generator == nil {
		t.Error("generator should not be nil")
	}
}

func TestNewGeneratorWithPath(t *testing.T) {
	bits := 256
	hdpath := "m/44'/60'/0'/0"
	generator, err := NewWalletGenerator(bits, hdpath)
	if err != nil {
		t.Errorf("Error with %v", err)
	}
	if generator == nil {
		t.Error("generator should not be nil")
	}
}

func TestNewGeneratorWithWrongConfig(t *testing.T) {
	bits := 512
	hdpath := "m/'44'/60'/0'/0"

	generator, err := NewWalletGenerator(bits, "")
	if err == nil {
		t.Errorf("error should not be nil")
	}
	if generator != nil {
		t.Error("generator should be nil")
	}

	generator, err = NewWalletGenerator(256, hdpath)
	if err == nil {
		t.Errorf("error should not be nil")
	}
	if generator != nil {
		t.Error("generator should be nil")
	}

}

func TestGenerate100Wallets(t *testing.T) {
	number := 100

	generator, err := NewWalletGenerator(256, "")
	if err != nil {
		t.Errorf("Error with %v", err)
	}
	if generator == nil {
		t.Error("generator should not be nil")
	}

	func() {
		index, wallets := generator.GenerateMultipleWallets(Config{
			Number:   number,
			Contains: []string{"0x"},
		})

		walletCounter := 0
		for wallet := range wallets {
			if wallet.Address == "" {
				t.Error("wallet.Address should not be empty")
			}
			if wallet.PrivateKey == "" {
				t.Error("wallet.PrivateKey should not be empty")
			}
			if wallet.Mnemonic == "" {
				t.Error("wallet.Mnemonic should not be empty")
			}
			walletCounter++
		}

		if walletCounter != number {
			t.Errorf("walletCounter got %v want %v", walletCounter, number)
		}

		indexCounter := 0
		for i := range index {
			if i >= number {
				t.Errorf("index should less than %v, got %v", number, i)
			}
			indexCounter++
		}

		if indexCounter != number {
			t.Errorf("indexCounter got %v want %v", indexCounter, number)
		}
	}()

	func() {
		index, wallets := generator.GenerateMultipleWallets(Config{
			Number:   number,
			Strict:   true,
			Contains: []string{"0x"},
			Prefix:   "0x",
			Suffix:   "0",
		})

		walletCounter := 0
		for wallet := range wallets {
			if wallet.Address == "" {
				t.Error("wallet.Address should not be empty")
			}
			if wallet.PrivateKey == "" {
				t.Error("wallet.PrivateKey should not be empty")
			}
			if wallet.Mnemonic == "" {
				t.Error("wallet.Mnemonic should not be empty")
			}
			walletCounter++
		}

		indexCounter := 0
		for i := range index {
			if i >= number {
				t.Errorf("index should less than %v, got %v", number, i)
			}
			indexCounter++
		}

		if indexCounter != number {
			t.Errorf("indexCounter got %v want %v", indexCounter, number)
		}
	}()

}
