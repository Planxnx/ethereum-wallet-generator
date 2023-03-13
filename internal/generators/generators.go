package generators

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Planxnx/eth-wallet-gen/pkg/progressbar"
	"github.com/Planxnx/eth-wallet-gen/pkg/wallets"
)

type Config struct {
	ProgressBar *progressbar.ProgressBar
	DryRun      bool
	Concurrency int
	Number      int
}

type Generator struct {
	walletsRepo *wallets.Repository
	config      Config
}

func New(walletsRepo *wallets.Repository, cfg Config) *Generator {
	return &Generator{
		walletsRepo: walletsRepo,
		config:      cfg,
	}
}

func (g *Generator) Start(ctx context.Context) (err error) {
	bar := g.config.ProgressBar
	defer func() {
		_ = bar.Finish()
		if g.config.DryRun {
			return
		}

		if result := g.walletsRepo.Results(); len(result) > 0 {
			fmt.Printf("\n%-42s %s\n", "Address", "Seed")
			fmt.Printf("%-42s %s\n", strings.Repeat("-", 42), strings.Repeat("-", 160))
			fmt.Println(result)
		}
	}()

	resolvedCount := 0
	var wg sync.WaitGroup
	// generate wallets without db
	semph := make(chan int, g.config.Concurrency)
	for i := 0; i < g.config.Number || g.config.Number < 0; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			semph <- 1
			wg.Add(1)

			go func() {
				defer func() {
					<-semph
					wg.Done()
				}()

				ok, err := g.walletsRepo.Generate()
				if err != nil {
					log.Printf("Error: %+v", err)
					return
				}
				if !ok {
					return
				}

				resolvedCount++
				_ = g.config.ProgressBar.Increment()
				_ = g.config.ProgressBar.SetResolved(resolvedCount)
			}()
		}
	}
	wg.Wait()

	return nil
}
