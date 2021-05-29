package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var wg sync.WaitGroup

type Wallet struct {
	Address    string
	PrivateKey string
	Mnemonic   string
	Bits       int
	HDPath     string
	CreatedAt  time.Time
	gorm.Model
}

func NewWallet(bits int, hdPath string) *Wallet {
	mnemonic, _ := hdwallet.NewMnemonic(bits)

	return &Wallet{
		Mnemonic:  mnemonic,
		HDPath:    hdPath,
		CreatedAt: time.Now(),
	}
}

func (w *Wallet) createWallet(mnemonic string) *Wallet {
	wallet, _ := hdwallet.NewFromMnemonic(w.Mnemonic)

	path := hdwallet.DefaultBaseDerivationPath
	if w.HDPath != "" {
		path = hdwallet.MustParseDerivationPath(w.HDPath)
	}

	account, _ := wallet.Derive(path, false)
	pk, _ := wallet.PrivateKeyHex(account)

	w.Address = account.Address.Hex()
	w.PrivateKey = pk
	w.UpdatedAt = time.Now()

	return w
}

func generateNewWallet(bits int) *Wallet {
	mnemonic, _ := hdwallet.NewMnemonic(bits)
	wallet := createWallet(mnemonic)
	wallet.Bits = bits
	return wallet
}

func createWallet(mnemonic string) *Wallet {
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	account, _ := wallet.Derive(hdwallet.DefaultBaseDerivationPath, false)
	pk, _ := wallet.PrivateKeyHex(account)

	return &Wallet{
		Address:    account.Address.Hex(),
		PrivateKey: pk,
		Mnemonic:   mnemonic,
		HDPath:     account.URL.Path,
		CreatedAt:  time.Now(),
	}
}

func main() {
	fmt.Println("Starting...")
	fmt.Println("")

	number := flag.Int("n", 10, "set number of wallets to generate")
	dbPath := flag.String("db", "", "set sqlite output path eg. wallets.db")
	concurrency := flag.Int("c", 1, "set number of concurrency")
	bits := flag.Int("bit", 256, "set number of entropy bits [128, 256]")
	contain := flag.String("contain", "", "used to check the given letters present in the given string or not")
	isLog := flag.Bool("l", false, "show logging data in stdout")
	flag.Parse()

	now := time.Now()
	count := 0

	if *dbPath != "" {
		db, err := gorm.Open(sqlite.Open(*dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&Wallet{})
		bar := pb.StartNew(*number)
		bar.SetTemplate(pb.Default)
		bar.SetTemplateString(`{{string . "prefix"}}{{counters . }} {{bar . "" "█" "█" "" "" | rndcolor}} {{percent . }} {{speed . }}{{string . "suffix"}}`)

		for i := 0; i < *number; i += *concurrency {
			tx := db.Begin()
			for j := 0; j < *concurrency && i+j < *number; j++ {
				wg.Add(1)

				go func(j int) {
					defer wg.Done()

					wallet := generateNewWallet(*bits)
					bar.Increment()

					if *contain != "" && !strings.Contains(wallet.Address, *contain) {
						return
					}

					if *isLog {
						fmt.Printf("[%v] %v, %v:\n", i+j+1, time.Now().Local().String(), time.Since(now))
					}

					tx.Create(wallet)
					count++
				}(j)
			}
			wg.Wait()
			tx.Commit()
		}
		bar.Finish()

	} else {
		fmt.Printf("\n%-42s %s\n", "Address", "Seed (24 words)")
		fmt.Printf("%-42s %s\n", strings.Repeat("-", 42), strings.Repeat("-", 160))

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
	}

	fmt.Printf("\nResolved Speed: %.2f w/s\n", float64(count)/time.Since(now).Seconds())
	fmt.Printf("Total Duration: %v\n", time.Since(now))
	fmt.Printf("Total Wallet Resolved: %d w\n", count)
}
