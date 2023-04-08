package generators

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/planxnx/ethereum-wallet-generator/internal/progressbar"
	"github.com/planxnx/ethereum-wallet-generator/internal/repository"
	"github.com/planxnx/ethereum-wallet-generator/wallets"
)

type Config struct {
	AddresValidator func(address string) bool
	ProgressBar     progressbar.ProgressBar
	DryRun          bool
	Concurrency     int
	Number          int
	Limit           int
}

type Generator struct {
	walletGen wallets.Generator
	repo      repository.Repository
	config    Config

	isShutdown     atomic.Bool
	shutdownSignal chan struct{}
	shutDownWg     sync.WaitGroup
}

func New(walletGen wallets.Generator, repo repository.Repository, cfg Config) *Generator {
	return &Generator{
		walletGen:      walletGen,
		repo:           repo,
		config:         cfg,
		shutdownSignal: make(chan struct{}),
	}
}

func (g *Generator) Start() (err error) {
	g.isShutdown.Store(false)
	g.shutDownWg.Add(1)
	defer g.shutDownWg.Done()

	var (
		bar           = g.config.ProgressBar
		resolvedCount atomic.Int64
		start         = time.Now()
	)
	defer func() {
		_ = bar.Finish()

		if err := g.repo.Close(); err != nil {
			// Ignore error
			log.Printf("[ERROR] failed to close repository: %+v\n", err)
		}

		if w := g.repo.Result(); len(w) > 0 && !g.config.DryRun {
			var result strings.Builder
			for _, wallet := range w {
				if _, err := fmt.Fprintf(&result, "%-42s %s\n", wallet.Address, wallet.Mnemonic); err != nil {
					continue
				}
			}

			fmt.Printf("\n%-42s %s\n", "Address", "Seed")
			fmt.Printf("%-42s %s\n", strings.Repeat("-", 42), strings.Repeat("-", 90))
			fmt.Println(result.String())
		}

		fmt.Printf("\nResolved Speed: %.2f w/s\n", float64(resolvedCount.Load())/time.Since(start).Seconds())
		fmt.Printf("Total Duration: %v\n", time.Since(start))
		fmt.Printf("Total Wallet Resolved: %d w\n", resolvedCount.Load())
		fmt.Printf("\nCopyright (C) 2023 Planxnx <planxthanee@gmail.com>\n")

		g.isShutdown.Store(true)
	}()

	var wg sync.WaitGroup
	commands := make(chan struct{})
	for i := 0; i < g.config.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range commands {
				if !(resolvedCount.Load() < int64(g.config.Limit) || g.config.Limit < 0) {
					return
				}

				wallet, err := g.walletGen()
				if err != nil {
					// Ignore error
					log.Printf("[ERROR] failed to generate wallet: %+v\n", err)
					continue
				}

				isOk := true
				if g.config.AddresValidator != nil {
					isOk = g.config.AddresValidator(wallet.Address)
				}

				if isOk {
					if err := g.repo.Insert(wallet); err != nil {
						// Ignore error
						log.Printf("[ERROR] failed to insert wallet to db: %+v\n", err)
						continue
					}
					resolvedCount.Add(1)
				}

				_ = bar.SetResolved(int(resolvedCount.Load()))
				_ = bar.Increment()
			}
		}()
	}

mainloop:
	for i := 0; (i < g.config.Number || g.config.Number < 0) && (resolvedCount.Load() < int64(g.config.Limit) || g.config.Limit < 0); i++ {
		select {
		case <-g.shutdownSignal:
			break mainloop
		default:
			commands <- struct{}{}
		}
	}

	close(commands)
	wg.Wait()
	return nil
}

func (g *Generator) Shutdown() (err error) {
	if g.isShutdown.Load() {
		return nil
	}
	go func() {
		g.shutdownSignal <- struct{}{}
	}()
	g.shutDownWg.Wait()
	return nil
}
