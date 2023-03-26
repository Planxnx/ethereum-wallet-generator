package generators

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Planxnx/eth-wallet-gen/pkg/progressbar"
	"github.com/Planxnx/eth-wallet-gen/pkg/wallets"
)

type Config struct {
	ProgressBar      progressbar.ProgressBar
	DryRun           bool
	Concurrency      int
	Number           int
	HideStdoutResult bool
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
	var (
		bar           = g.config.ProgressBar
		resolvedCount atomic.Int64
		start         = time.Now()
	)
	defer func() {
	}()
	defer func() {
		_ = bar.Finish()

		if err := g.walletsRepo.Commit(); err != nil {
			// Ignore error
			log.Printf("Gerate Error: %+v", err)
		}

		if !g.config.HideStdoutResult {
			if result := g.walletsRepo.Results(); len(result) > 0 && !g.config.DryRun {
				fmt.Printf("\n%-42s %s\n", "Address", "Seed")
				fmt.Printf("%-42s %s\n", strings.Repeat("-", 42), strings.Repeat("-", 160))
				fmt.Println(result)
			}

			fmt.Printf("\nResolved Speed: %.2f w/s\n", float64(resolvedCount.Load())/time.Since(start).Seconds())
			fmt.Printf("Total Duration: %v\n", time.Since(start))
			fmt.Printf("Total Wallet Resolved: %d w\n", resolvedCount.Load())
			fmt.Printf("\nCopyright (C) 2023 Planxnx <planxthanee@gmail.com>\n")
		}
	}()

	var wg sync.WaitGroup
	commands := make(chan struct{})
	for i := 0; i < g.config.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range commands {
				ok, err := g.walletsRepo.Generate()
				if err != nil {
					// Ignore error
					log.Printf("Gerate Error: %+v", err)
					return
				}
				if ok {
					resolvedCount.Add(1)
				}

				_ = bar.Increment()
				_ = bar.SetResolved(int(resolvedCount.Load()))
			}
		}()
	}

mainloop:
	for i := 0; i < g.config.Number || g.config.Number < 0; i++ {
		select {
		case <-ctx.Done():
			break mainloop
		default:
			commands <- struct{}{}
		}
	}

	close(commands)
	wg.Wait()

	return nil
}
