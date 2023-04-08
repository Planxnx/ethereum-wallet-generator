package bip39

import (
	"fmt"
	"testing"

	"github.com/tyler-smith/go-bip39"
)

// nolint: wrapcheck,gocritic
func BenchmarkBIP39(b *testing.B) {
	bitSizes := []int{128, 256}

	generators := map[string]func(entropy []byte) (string, error){
		"old": bip39.NewMnemonic,
		"new": NewMnemonic,
	}

	for name, generator := range generators {
		for _, bitSize := range bitSizes {
			b.Run(fmt.Sprintf("%s:%d", name, bitSize), func(b *testing.B) {
				b.ResetTimer()
				entropy, err := bip39.NewEntropy(bitSize)
				if err != nil {
					b.Fatal(err)
				}
				for i := 0; i < b.N; i++ {
					_, _ = generator(entropy)
				}
			})
		}
	}
}
