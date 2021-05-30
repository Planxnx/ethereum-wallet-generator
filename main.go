package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/cheggaaa/pb/v3"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/schollz/progressbar/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	wg     sync.WaitGroup
	result strings.Builder
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

//ProgressBar progressbar logging with 2 mode
type ProgressBar struct {
	StandardMode   *pb.ProgressBar
	CompatibleMode *progressbar.ProgressBar
}

//NewWallet .
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

//NewProgressBar .
func NewProgressBar(number int, isCompatible bool) (bar *ProgressBar) {
	if isCompatible {
		bar = &ProgressBar{
			CompatibleMode: progressbar.NewOptions(number,
				progressbar.OptionSetItsString("w"),
				progressbar.OptionSetPredictTime(true),
				progressbar.OptionShowIts(),
				progressbar.OptionShowCount(),
				progressbar.OptionFullWidth(),
			),
		}
		return
	}
	bar = &ProgressBar{
		StandardMode: pb.StartNew(number),
	}
	bar.StandardMode.SetTemplate(pb.Default)
	bar.StandardMode.SetTemplateString(`{{counters . }} | {{bar . "" "█" "█" "" "" | rndcolor}} | {{percent . }} | {{speed . }} | {{string . "resolved"}}`)
	return
}

//Increment increment progress
func (bar *ProgressBar) Increment() error {
	if bar.CompatibleMode != nil {
		return bar.CompatibleMode.Add(1)
	}
	return bar.StandardMode.Increment().Err()
}

//SetResolved set resolved wallet number
func (bar *ProgressBar) SetResolved(resolved int) error {
	if bar.StandardMode != nil {
		return bar.StandardMode.Set("resolved", fmt.Sprintf("resolved: %v", resolved)).Err()
	}
	return nil
}

//Finish close progress bar
func (bar *ProgressBar) Finish() error {
	if bar.CompatibleMode != nil {
		return bar.CompatibleMode.Close()
	}
	return bar.StandardMode.Finish().Err()
}

func main() {

	interrupt := make(chan os.Signal, 1)

	signal.Notify(
		interrupt,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)

	fmt.Println("===============ETH Wallet Generator===============")
	fmt.Println(" ")

	number := flag.Int("n", 10, "set number of wallets to generate (set number to -1 for Infinite loop ∞)")
	dbPath := flag.String("db", "", "set sqlite output name eg. wallets.db (db file will create in /db)")
	concurrency := flag.Int("c", 1, "set number of concurrency")
	bits := flag.Int("bit", 256, "set number of entropy bits [128, 256]")
	contain := flag.String("contains", "", "used to check the given letters present in the given string or not")
	strict := flag.Bool("strict", false, "strict contains mode (required contains to use)")
	isDryrun := flag.Bool("dryrun", false, "generate wallet without result (used for benchamark speed)")
	isCompatible := flag.Bool("compatible", false, "logging compatible mode (turn on this to fix logging glitch)")
	flag.Parse()

	contains := strings.Split(*contain, ",")
	findContains := func(address string) bool {
		cb := func(contain string) bool {
			return strings.Contains(address, contain)
		}

		if *strict {
			if !have(contains, cb) {
				return false
			}
		} else {
			if !some(contains, cb) {
				return false
			}
		}
		return true
	}
	if *number < 0 {
		*number = -1
	}

	now := time.Now()
	resolvedCount := 0

	defer func() {
		fmt.Printf("\nResolved Speed: %.2f w/s\n", float64(resolvedCount)/time.Since(now).Seconds())
		fmt.Printf("Total Duration: %v\n", time.Since(now))
		fmt.Printf("Total Wallet Resolved: %d w\n", resolvedCount)

		fmt.Printf("\nCopyright (C) 2021 Planxnx <planxthanee@gmail.com>\n")
	}()

	bar := NewProgressBar(*number, *isCompatible)
	defer func() {
		bar.Finish()
		if *isDryrun {
			return
		}
		if result.Len() > 0 {
			fmt.Printf("\n%-42s %s\n", "Address", "Seed")
			fmt.Printf("%-42s %s\n", strings.Repeat("-", 42), strings.Repeat("-", 160))
			fmt.Println(result.String())
		}
	}()

	go func() {
		defer func() {
			interrupt <- syscall.SIGQUIT
		}()

		if *dbPath != "" {
			db, err := gorm.Open(sqlite.Open("./db/"+*dbPath), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
				DryRun: *isDryrun,
			})
			if err != nil {
				panic(err)
			}

			if !*isDryrun {
				db.AutoMigrate(&Wallet{})
			}

			for i := 0; i < *number || *number < 0; i += *concurrency {
				tx := db.Begin()
				txResolved := 0
				for j := 0; j < *concurrency && (i+j < *number || *number < 0); j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						wallet := generateNewWallet(*bits)
						bar.Increment()

						// if *contain != "" && !strings.Contains(wallet.Address, *contain) {
						// 	return
						// }

						if len(contains) != 0 && !findContains(wallet.Address) {
							return
						}

						txResolved++
						tx.Create(wallet)
					}(j)
				}
				wg.Wait()
				tx.Commit()
				resolvedCount += txResolved
				bar.SetResolved(resolvedCount)
			}
			return
		}

		for i := 0; i < *number || *number < 0; i += *concurrency {
			for j := 0; j < *concurrency && (i+j < *number || *number < 0); j++ {
				wg.Add(1)

				go func(j int) {
					defer wg.Done()

					wallet := generateNewWallet(*bits)
					bar.Increment()

					// if *contain != "" && !strings.Contains(wallet.Address, *contain) {
					// 	return
					// }

					if len(contains) != 0 && !findContains(wallet.Address) {
						return
					}

					fmt.Fprintf(&result, "%-18s %s\n", wallet.Address, wallet.Mnemonic)
					resolvedCount++
					bar.SetResolved(resolvedCount)
				}(j)
			}
			wg.Wait()
		}
		bar.Finish()
	}()
	<-interrupt
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
