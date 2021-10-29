package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Planxnx/eth-wallet-gen/generator"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"eth-wallet-gen/common"
)

var (
	result strings.Builder
)

func init() {
	if _, err := os.Stat("db"); os.IsNotExist(err) {
		if err := os.Mkdir("db", 0750); err != nil {
			panic(err)
		}
	}
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
	concurrency := flag.Int("c", 1, "set concurrency value")
	bits := flag.Int("bit", 256, "set number of entropy bits [128, 256]")
	strict := flag.Bool("strict", false, "strict contains mode (required contains to use)")
	contain := flag.String("contains", "", "used to check if the given letters are present in the given string")
	prefix := flag.String("prefix", "", "used to check if the given letters are present in the prefix string")
	suffix := flag.String("suffix", "", "used to check if the given letters are present in the suffix string")
	isDryrun := flag.Bool("dryrun", false, "generate wallet without a result (used for benchmark speed)")
	isCompatible := flag.Bool("compatible", false, "logging compatible mode (turn this on to fix logging glitch)")
	flag.Parse()

	contains := strings.Split(*contain, ",")
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

	var bar *common.ProgressBar
	if *isCompatible {
		bar = common.NewCompatibleProgressBar(*number)
	} else {
		bar = common.NewStandardProgressBar(*number)
	}

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

	var db *gorm.DB
	if *dbPath != "" {
		db, err := gorm.Open(sqlite.Open("./db/"+*dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			DryRun: *isDryrun,
		})
		if err != nil {
			panic(err)
		}
		if !*isDryrun {
			db.AutoMigrate(&generator.Wallet{})
		}
	}

	walletgen, err := generator.NewWalletGenerator(*bits, "")
	if err != nil {
		panic(err)
	}

	index, wallets := walletgen.GenerateMultipleWallets(generator.Config{
		Number:      *number,
		Concurrency: *concurrency,
		Strict:      *strict,
		Contains:    contains,
		Prefix:      *prefix,
		Suffix:      *suffix,
	})

	go func() {
		defer func() {
			bar.Finish()
			interrupt <- syscall.SIGQUIT
		}()

		for range index {
			if *number < 0 {
				bar.Increment()
			}
		}
	}()
	go func() {

		for wallet := range wallets {
			resolvedCount++
			bar.SetResolved(resolvedCount)

			if *dbPath != "" {
				//NOTE: Performance Issues (too much db connections!)
				db.Create(wallet)
			} else {
				fmt.Fprintf(&result, "%-18s %s\n", wallet.Address, wallet.Mnemonic)
			}

			if *number > 0 {
				bar.Increment()
			}
		}
	}()

	<-interrupt
}
