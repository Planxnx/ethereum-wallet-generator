package utils

// Has0xPrefix checks if the input string has 0x prefix or not.
//
// Returns `true“ if the input string has 0x prefix, otherwise `false`.
func Has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

// Add0xPrefix returns the input string with 0x prefix.
func Add0xPrefix(input string) string {
	if !Has0xPrefix(input) {
		return "0x" + input
	}
	return input
}
