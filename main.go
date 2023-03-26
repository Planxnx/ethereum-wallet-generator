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

	"github.com/Planxnx/eth-wallet-gen/internal/generators"
	"github.com/Planxnx/eth-wallet-gen/pkg/progressbar"
	"github.com/Planxnx/eth-wallet-gen/pkg/utils"
	"github.com/Planxnx/eth-wallet-gen/pkg/wallets"
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

	number := flag.Int("n", 10, "set number of wallets to generate (set number to -1 for Infinite loop ∞)")
	dbPath := flag.String("db", "", "set sqlite output name eg. wallets.db (db file will create in /db)")
	concurrency := flag.Int("c", 1, "set concurrency value")
	bits := flag.Int("bit", 128, "set number of entropy bits [128, 256]")
	strict := flag.Bool("strict", false, "strict contains mode (required contains to use)")
	contain := flag.String("contains", "", "used to check if the given letters are present in the given string")
	prefix := flag.String("prefix", "", "used to check if the given letters are present in the prefix string")
	suffix := flag.String("suffix", "", "used to check if the given letters are present in the suffix string")
	regEx := flag.String("regex", "", "used to check if the given letters are present in the regex format")
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

	// now := time.Now()
	// resolvedCount := 0

	// defer func() {
	// 	fmt.Printf("\nResolved Speed: %.2f w/s\n", float64(resolvedCount)/time.Since(now).Seconds())
	// 	fmt.Printf("Total Duration: %v\n", time.Since(now))
	// 	fmt.Printf("Total Wallet Resolved: %d w\n", resolvedCount)

	// 	fmt.Printf("\nCopyright (C) 2023 Planxnx <planxthanee@gmail.com>\n")
	// }()

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
		log.Printf("gracefully shutting down...\n")

		if err := walletsRepo.Close(); err != nil {
			log.Printf("WalletsRepo Close Error: %+v", err)
		}

		if db != nil {
			if err := utils.Must(db.DB()).Close(); err != nil {
				log.Printf("DB Close Error: %+v", err)
			}
		}
	}()

	if err := generator.Start(ctx); err != nil {
		log.Printf("Generator Error: %+v", err)
	}
}
