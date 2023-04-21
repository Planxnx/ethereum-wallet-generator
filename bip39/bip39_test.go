package bip39

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGenerateFromEntropy(t *testing.T) {
	type TestSpec struct {
		entropy          []byte
		expectedMnemonic string
	}

	testCases := map[string][]TestSpec{
		"128bits": {
			{
				entropy:          []byte{157, 167, 82, 14, 138, 57, 220, 233, 60, 83, 236, 19, 121, 207, 199, 151},
				expectedMnemonic: "oval dentist lonely behave oven innocent vanish laugh beach solar vehicle connect",
			},
			{
				entropy:          []byte{103, 15, 41, 115, 192, 199, 222, 77, 33, 34, 50, 57, 174, 229, 13, 231},
				expectedMnemonic: "grow junk friend light law charge loyal edge defy jaguar drop sort",
			},
			{
				entropy:          []byte{172, 237, 228, 119, 25, 39, 253, 133, 139, 151, 201, 62, 48, 247, 58, 15},
				expectedMnemonic: "provide hundred build crane lemon security comic weird dilemma marble soldier butter",
			},
			{
				entropy:          []byte{41, 126, 238, 172, 119, 220, 20, 164, 174, 84, 12, 165, 14, 128, 206, 240},
				expectedMnemonic: "city wash prison use scout false rich light pink inject crisp ticket",
			},
			{
				entropy:          []byte{21, 42, 240, 221, 4, 84, 149, 97, 139, 174, 9, 186, 150, 186, 144, 14},
				expectedMnemonic: "benefit fiscal dance anger enable radio concert scorpion ritual remind piano bronze",
			},
		},
		"256bits": {
			{
				entropy:          []byte{6, 137, 161, 26, 237, 39, 33, 110, 228, 193, 105, 245, 181, 108, 133, 38, 168, 102, 174, 148, 240, 219, 81, 74, 194, 162, 228, 218, 16, 148, 57, 156},
				expectedMnemonic: "alley escape effort surge improve resist narrow coffee volcano problem candy essence major firm fatigue bread eyebrow figure post situate patient energy toy menu",
			},
			{
				entropy:          []byte{24, 191, 8, 191, 188, 214, 228, 149, 31, 99, 124, 92, 243, 25, 195, 46, 56, 206, 78, 177, 177, 183, 231, 19, 100, 156, 57, 0, 124, 102, 156, 96},
				expectedMnemonic: "board weapon copper keen hour enhance laugh hurdle friend occur ignore fragile mind chef short dad train open check improve amazing crew immense bonus",
			},
		},
	}

	for name, testSpecs := range testCases {
		t.Run(name, func(t *testing.T) {
			for _, testSpec := range testSpecs {
				actualMnemonic, err := NewMnemonic(testSpec.entropy)
				if err != nil {
					t.Errorf("failed to generate mnemonic %+v", err)
				}

				assert.Equal(t, testSpec.expectedMnemonic, actualMnemonic)
			}
		})
	}
}
