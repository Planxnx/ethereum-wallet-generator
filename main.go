package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Planxnx/ethereum-wallet-generator/internal/generators"
	"github.com/Planxnx/ethereum-wallet-generator/pkg/progressbar"
	"github.com/Planxnx/ethereum-wallet-generator/pkg/utils"
	"github.com/Planxnx/ethereum-wallet-generator/pkg/wallets"
)

func init() {
	if _, err := os.Stat("db"); os.IsNotExist(err) {
		if err := os.Mkdir("db", 0o750); err != nil {
			panic(err)
		}
	}
}

func main() {
	// Context with gracefully shutdown signal
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)
	defer stop()

	fmt.Println("===============ETH Wallet Generator===============")
	fmt.Println(" ")

	number := flag.Int("n", 10, "set number of generate times (not number of result wallets) (set number to -1 for Infinite loop ∞)")
	dbPath := flag.String("db", "", "set sqlite output name eg. wallets.db (db file will create in /db)")
	concurrency := flag.Int("c", 1, "set concurrency value")
	bits := flag.Int("bit", 128, "set number of entropy bits [128, 256]")
	strict := flag.Bool("strict", false, "strict contains mode (required contains to use)")
	contain := flag.String("contains", "", "show only result that contained with the given letters (support for multiple characters)")
	prefix := flag.String("prefix", "", "show only result that prefix was matched with the given letters  (support for single character)")
	suffix := flag.String("suffix", "", "show only result that suffix was matched with the given letters (support for single character)")
	regEx := flag.String("regex", "", "show only result that was matched with given regex (eg. ^0x99 or ^0x00)")
	isDryrun := flag.Bool("dryrun", false, "generate wallet without a result (used for benchmark speed)")
	isCompatible := flag.Bool("compatible", false, "logging compatible mode (turn this on to fix logging glitch)")
	flag.Parse()

	r, err := regexp.Compile(*regEx)
	if err != nil {
		panic(err)
	}
	contains := strings.Split(*contain, ",")
	validateAddress := func(address string) bool {
		isValid := true
		if len(contains) > 0 {
			cb := func(contain string) bool {
				return strings.Contains(address, contain)
			}
			if *strict {
				if !utils.Have(contains, cb) {
					isValid = false
				}
			} else {
				if !utils.Some(contains, cb) {
					isValid = false
				}
			}
		}

		if *prefix != "" {
			if !strings.HasPrefix(address, *prefix) {
				isValid = false
			}
		}

		if *suffix != "" {
			if !strings.HasSuffix(address, *suffix) {
				isValid = false
			}
		}

		if *regEx != "" && !r.MatchString(address) {
			isValid = false
		}

		return isValid
	}
	if *number < 0 {
		*number = -1
	}

	var bar progressbar.ProgressBar
	if *isCompatible {
		bar = progressbar.NewCompatibleProgressBar(*number)
	} else {
		bar = progressbar.NewStandardProgressBar(*number)
	}

	var db *gorm.DB
	if *dbPath != "" {
		db, err = gorm.Open(sqlite.Open("./db/"+*dbPath), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			DryRun:                 *isDryrun,
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}

		defer func() {
			db, _ := db.DB()
			db.Close()
		}()

		if !*isDryrun {
			if err := db.AutoMigrate(&wallets.Wallet{}); err != nil {
				panic(err)
			}
		}
	}

	walletsRepo := wallets.NewRepository(wallets.RepositoryConfig{
		BitSize:            *bits,
		AddresValidator:    validateAddress,
		DB:                 db,
		DBTransactionsSize: uint64(*concurrency),
	})

	generator := generators.New(
		walletsRepo,
		generators.Config{
			ProgressBar: bar,
			DryRun:      *isDryrun,
			Concurrency: *concurrency,
			Number:      *number,
		},
	)

	go func() {
		<-ctx.Done()

		if err := generator.Shutdown(); err != nil {
			log.Printf("Generator Shutdown Error: %+v", err)
		}

		if err := walletsRepo.Close(); err != nil {
			log.Printf("WalletsRepo Close Error: %+v", err)
		}

		if db != nil {
			if err := utils.Must(db.DB()).Close(); err != nil {
				log.Printf("DB Close Error: %+v", err)
			}
		}
	}()

	if err := generator.Start(); err != nil {
		log.Printf("Generator Error: %+v", err)
	}
}
