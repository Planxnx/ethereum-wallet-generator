// Package bip39 is the Golang implementation of the BIP39 spec.
//
// This package is a optimized version of github.com/tyler-smith/go-bip39.
// Thanks to Tyler Smith for the original implementation.
//
// The official BIP39 spec can be found at
// https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"strings"

	"github.com/holiman/uint256"
	"github.com/pkg/errors"
)

var (
	one = uint256.NewInt(1)
	two = uint256.NewInt(2)

	bitsChunkSize   = 11
	shift11BitsMask = new(uint256.Int).Lsh(one, uint(bitsChunkSize)) // 2^11 = 2048
	last11BitsMask  = new(uint256.Int).Sub(shift11BitsMask, one)     // 2^11 - 1 = 2047
)

// NewEntropy will create random entropy bytes
// so long as the requested size bitSize is an appropriate size.
//
// bitSize has to be a multiple 32 and be within the inclusive range of {128, 256}.
func NewEntropy(bitSize int) ([]byte, error) {
	if err := validateEntropyBitSize(bitSize); err != nil {
		return nil, errors.WithStack(err)
	}

	entropy := make([]byte, bitSize/8)
	if _, err := rand.Read(entropy); err != nil {
		return nil, errors.WithStack(err)
	}

	return entropy, nil
}

// NewMnemonic will return a string consisting of the mnemonic words for
// the given entropy.
// If the provide entropy is invalid, an error will be returned.
func NewMnemonic(entropy []byte) (string, error) {
	// Compute some lengths for convenience.
	entropyBitLength := len(entropy) * 8
	checksumBitLength := entropyBitLength / 32
	sentenceLength := (entropyBitLength + checksumBitLength) / bitsChunkSize

	// Validate that the requested size is supported.
	err := validateEntropyBitSize(entropyBitLength)
	if err != nil {
		return "", err
	}

	// Add checksum to entropy.
	// Entropy as an int so we can bitmask without worrying about bytes slices.
	entropyInt := addChecksum(entropy)

	// Throw away uint256.Int for AND masking.
	word := uint256.NewInt(0)

	// Break entropy up into sentenceLength chunks of 11 bits.
	// For each word AND mask the rightmost 11 bits and find the word at that index.
	// Then bitshift entropy 11 bits right and repeat.
	// Add to the last empty slot so we can work with LSBs instead of MSB.

	var w strings.Builder
	w.Grow(sentenceLength * bitsChunkSize)
	for i := sentenceLength - 1; i >= 0; i-- {
		// Get 11 right most bits and bitshift 11 to the right for next time.
		word.And(entropyInt, last11BitsMask)
		entropyInt.Div(entropyInt, shift11BitsMask)

		// Convert bytes to an index and add that word to the list.
		if i == sentenceLength-1 {
			w.WriteString(Words[word.Uint64()])
		} else {
			w.WriteString(" ")
			w.WriteString(Words[word.Uint64()])
		}
	}

	return w.String(), nil
}

// Appends to data the first (len(data) / 32)bits of the result of sha256(data)
// abd returns the result as a uint256.Int.
//
// Currently only supports data up to 32 bytes.
func addChecksum(data []byte) *uint256.Int {
	// Get first byte of sha256
	hash := computeChecksum(data)
	firstChecksumByte := hash[0]

	// len() is in bytes so we divide by 4
	checksumBitLength := uint(len(data) / 4)

	// For each bit of check sum we want we shift the data one the left
	// and then set the (new) right most bit equal to checksum bit at that index
	// staring from the left
	dataInt := new(uint256.Int).SetBytes(data)

	for i := uint(0); i < checksumBitLength; i++ {
		// Bitshift 1 left
		dataInt.Mul(dataInt, two)

		// Set rightmost bit if leftmost checksum bit is set
		if firstChecksumByte&(1<<(7-i)) > 0 {
			dataInt.Or(dataInt, one)
		}
	}

	return dataInt
}

func computeChecksum(data []byte) []byte {
	hasher := sha256.New()
	_, _ = hasher.Write(data) // This error is guaranteed to be nil

	return hasher.Sum(nil)
}

// validateEntropyBitSize ensures that entropy is the correct size for being a mnemonic.
func validateEntropyBitSize(bitSize int) error {
	if (bitSize%32) != 0 || bitSize < 128 || bitSize > 256 {
		return errors.New("Entropy length must between range of [128,256] and a multiple of 32")
	}

	return nil
}
