package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"sync"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	Address    string
	PrivateKey string
	Mnemonic   string
	CreatedAt  time.Time
}

func main() {
	fmt.Println("Starting...")

	isLog := flag.Bool("l", false, "show logging data in stdout")
	number := flag.Int("n", 10, "set number of wallets to generate")
	contain := flag.String("contain", "", "used to check the given letters present in the given string or not")
	concurrency := flag.Int("c", 1, "set number of concurrency")
	bits := flag.Int("bit", 256, "set number of entropy bits [128, 256]")
	flag.Parse()

	fmt.Printf("\n%-42s %s\n", "Address", "Seed (24 words)")
	fmt.Printf("%-42s %s\n", strings.Repeat("-", 42), strings.Repeat("-", 160))

	now := time.Now()

	var wg sync.WaitGroup
	count := 0
	for i := 0; i < *number; i += *concurrency {
		for j := 0; j < *concurrency && i+j < *number; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				wallet := generateNewWallet(*bits)
				if *contain != "" && !strings.Contains(wallet.Address, *contain) {
					return
				}
				if *isLog {
					fmt.Printf("[%v] %v, %v:\n", i+j+1, time.Now().Local().String(), time.Since(now))
				}
				fmt.Printf("%-18s %s\n", wallet.Address, wallet.Mnemonic)
				count++
			}(j)
		}
		wg.Wait()
	}
	fmt.Printf("\nDuration: %v\n", time.Since(now))
	fmt.Printf("Total Wallet: %d\n", count)
}

func generateNewWallet(bits int) *Wallet {
	mnemonic, _ := hdwallet.NewMnemonic(bits)
	return createWallet(mnemonic)
}

func createWallet(mnemonic string) *Wallet {
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	account, _ := wallet.Derive(hdwallet.DefaultBaseDerivationPath, false)
	pk, _ := wallet.PrivateKeyHex(account)

	return &Wallet{
		Address:    account.Address.Hex(),
		PrivateKey: pk,
		Mnemonic:   mnemonic,
		CreatedAt:  time.Now(),
	}
}
