// nolint
package bip39

import (
	_ "embed"
	"hash/crc32"
	"strings"
)

//go:embed wordlist.txt
var words string

// Words a slice of mnemonic words taken from the bip39 specification
// https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt
var (
	Words = strings.Split(strings.TrimSpace(words), "\n")
)

func init() {
	// Ensure word list is correct
	// $ wget https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt
	// $ crc32 english.txt
	// OUTPUT: c1dbd296
	checksum := crc32.ChecksumIEEE([]byte(words))
	if checksum != 0xc1dbd296 {
		panic("english checksum invalid")
	}
}
