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
	"encoding/binary"
	"math/big"
	"strings"

	"github.com/pkg/errors"
)

var (
	last11BitsMask  = big.NewInt(2047)
	shift11BitsMask = big.NewInt(2048)
	one             = big.NewInt(1)
	two             = big.NewInt(2)
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
	sentenceLength := (entropyBitLength + checksumBitLength) / 11

	// Validate that the requested size is supported.
	err := validateEntropyBitSize(entropyBitLength)
	if err != nil {
		return "", err
	}

	// Add checksum to entropy.
	entropy = addChecksum(entropy)

	// Break entropy up into sentenceLength chunks of 11 bits.
	// For each word AND mask the rightmost 11 bits and find the word at that index.
	// Then bitshift entropy 11 bits right and repeat.
	// Add to the last empty slot so we can work with LSBs instead of MSB.

	// Entropy as an int so we can bitmask without worrying about bytes slices.
	entropyInt := new(big.Int).SetBytes(entropy)

	// Throw away big.Int for AND masking.
	word := big.NewInt(0)

	var w strings.Builder
	w.Grow(sentenceLength * 11)
	for i := sentenceLength - 1; i >= 0; i-- {
		// Get 11 right most bits and bitshift 11 to the right for next time.
		word.And(entropyInt, last11BitsMask)
		entropyInt.Div(entropyInt, shift11BitsMask)

		// Get the bytes representing the 11 bits as a 2 byte slice.
		wordBytes := padByteSlice(word.Bytes(), 2)

		// Convert bytes to an index and add that word to the list.
		if i == sentenceLength-1 {
			w.WriteString(Words[binary.BigEndian.Uint16(wordBytes)])
		} else {
			w.WriteString(" ")
			w.WriteString(Words[binary.BigEndian.Uint16(wordBytes)])
		}
	}

	return w.String(), nil
}

// Appends to data the first (len(data) / 32)bits of the result of sha256(data)
// Currently only supports data up to 32 bytes.
func addChecksum(data []byte) []byte {
	// Get first byte of sha256
	hash := computeChecksum(data)
	firstChecksumByte := hash[0]

	// len() is in bytes so we divide by 4
	checksumBitLength := uint(len(data) / 4)

	// For each bit of check sum we want we shift the data one the left
	// and then set the (new) right most bit equal to checksum bit at that index
	// staring from the left
	dataBigInt := new(big.Int).SetBytes(data)

	for i := uint(0); i < checksumBitLength; i++ {
		// Bitshift 1 left
		dataBigInt.Mul(dataBigInt, two)

		// Set rightmost bit if leftmost checksum bit is set
		if firstChecksumByte&(1<<(7-i)) > 0 {
			dataBigInt.Or(dataBigInt, one)
		}
	}

	return dataBigInt.Bytes()
}

func computeChecksum(data []byte) []byte {
	hasher := sha256.New()
	_, _ = hasher.Write(data) // This error is guaranteed to be nil

	return hasher.Sum(nil)
}

// padByteSlice returns a byte slice of the given size with contents of the
// given slice left padded and any empty spaces filled with 0's.
func padByteSlice(slice []byte, length int) []byte {
	offset := length - len(slice)
	if offset <= 0 {
		return slice
	}

	newSlice := make([]byte, length)
	copy(newSlice[offset:], slice)

	return newSlice
}

// validateEntropyBitSize ensures that entropy is the correct size for being a mnemonic.
func validateEntropyBitSize(bitSize int) error {
	if (bitSize%32) != 0 || bitSize < 128 || bitSize > 256 {
		return errors.New("Entropy length must between range of [128,256] and a multiple of 32")
	}

	return nil
}
