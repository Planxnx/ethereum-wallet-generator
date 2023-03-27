package generators

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Planxnx/ethereum-wallet-generator/pkg/progressbar"
	"github.com/Planxnx/ethereum-wallet-generator/pkg/wallets"
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

	isShutdown     atomic.Bool
	shutdownSignal chan struct{}
	shutDownWg     sync.WaitGroup
}

func New(walletsRepo *wallets.Repository, cfg Config) *Generator {
	return &Generator{
		walletsRepo:    walletsRepo,
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

		g.isShutdown.Store(true)
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
