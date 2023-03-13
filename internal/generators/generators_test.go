package generators

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Planxnx/eth-wallet-gen/pkg/utils"
	"github.com/Planxnx/eth-wallet-gen/pkg/wallets"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TODO: refactor this test for better readability
func BenchmarkGenerte(b *testing.B) {
	type benchCase struct {
		name        string
		bitSize     int // 128 or 256
		concurrency int
		number      int
	}

	cases := []benchCase{
		{
			name:        "128bit-1concurrency-1000number",
			bitSize:     128,
			concurrency: 1,
			number:      1000,
		},
		{
			name:        "128bit-8concurrency-1000number",
			bitSize:     128,
			concurrency: 8,
			number:      1000,
		},
		{
			name:        "128bit-16concurrency-1000number",
			bitSize:     128,
			concurrency: 16,
			number:      1000,
		},
		{
			name:        "128bit-32concurrency-1000number",
			bitSize:     128,
			concurrency: 32,
			number:      1000,
		},
		{
			name:        "256bit-1concurrency-1000number",
			bitSize:     256,
			concurrency: 1,
			number:      1000,
		},
		{
			name:        "256bit-8concurrency-1000number",
			bitSize:     256,
			concurrency: 8,
			number:      1000,
		},
		{
			name:        "256bit-16concurrency-1000number",
			bitSize:     256,
			concurrency: 16,
			number:      1000,
		},
		{
			name:        "256bit-32concurrency-1000number",
			bitSize:     256,
			concurrency: 32,
			number:      1000,
		},
	}

	validateAddress := func(address string) bool {
		isValid := true
		cb := func(contain string) bool {
			return strings.Contains(address, contain)
		}
		if !utils.Have([]string{"0", "x"}, cb) {
			isValid = false
		}

		if !strings.HasPrefix(address, "0x") {
			isValid = false
		}

		return isValid
	}

	b.Run("Stdout", func(b *testing.B) {
		for _, c := range cases {
			b.Run(c.name, func(b *testing.B) {
				walletsRepo := wallets.NewRepository(wallets.RepositoryConfig{
					BitSize:            c.bitSize,
					AddresValidator:    validateAddress,
					DBTransactionsSize: uint64(c.concurrency),
				})

				generator := New(
					walletsRepo,
					Config{
						ProgressBar:      testProgressBar{},
						DryRun:           false,
						Concurrency:      c.concurrency,
						Number:           c.number,
						HideStdoutResult: true,
					},
				)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if err := generator.Start(context.Background()); err != nil {
						b.Errorf("Error: %+v", err)
					}
				}
			})
		}
	})
	b.Run("DB", func(b *testing.B) {
		for _, c := range cases {
			b.Run(c.name, func(b *testing.B) {
				dbPath := "./" + c.name
				db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
					Logger:                 logger.Default.LogMode(logger.Silent),
					DryRun:                 false,
					SkipDefaultTransaction: true,
				})
				if err != nil {
					b.Fatal(err)
				}

				defer func() {
					db, _ := db.DB()
					db.Close()

					if err := os.Remove(dbPath); err != nil {
						b.Fatal(err)
					}
				}()

				if err := db.AutoMigrate(&wallets.Wallet{}); err != nil {
					b.Fatal(err)
				}

				walletsRepo := wallets.NewRepository(wallets.RepositoryConfig{
					BitSize:            c.bitSize,
					AddresValidator:    validateAddress,
					DBTransactionsSize: uint64(c.concurrency),
					DB:                 db,
				})

				generator := New(
					walletsRepo,
					Config{
						ProgressBar:      testProgressBar{},
						DryRun:           false,
						Concurrency:      c.concurrency,
						Number:           c.number,
						HideStdoutResult: true,
					},
				)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if err := generator.Start(context.Background()); err != nil {
						b.Errorf("Error: %+v", err)
					}
				}
			})
		}
	})
}

type testProgressBar struct{}

func (testProgressBar) Increment() error {
	return nil
}

func (testProgressBar) SetResolved(resolved int) error {
	return nil
}

func (testProgressBar) Finish() error {
	return nil
}
