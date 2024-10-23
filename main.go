package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

const (
	DefaultPasswordLength = 12
)

// compileChars creates a slice of runes based on the specified character types.
//
// Parameters:
//   - includes: A string specifying which character types to include.
//     'l' for lowercase letters, 'u' for uppercase letters,
//     'n' for numbers, and 's' for symbols.
//
// Returns:
//   A shuffled slice of runes containing the specified character types.
//
// If the 'includes' string is empty, all character types are included.
// The function uses predefined slices for each character type and
// combines them based on the 'includes' string. The resulting slice
// is shuffled before being returned.
func compileChars(includes string) []rune {
	var (
		LowercaseLetters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
		UppercaseLetters = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		Numbers          = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
		Symbols          = []rune{'~', '`', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '-', '+', '=', '{', '[', '}', ']', '|', '\\', ':', ';', '"', '\'', '<', ',', '>', '.', '?', '/'}
	)

	var chars []rune

	// add all characters if none are requested
	if len(includes) == 0 {
		chars = append(chars, LowercaseLetters...)
		chars = append(chars, UppercaseLetters...)
		chars = append(chars, Numbers...)
		chars = append(chars, Symbols...)
		return shuffle(chars)
	}

	for _, opt := range includes {
		switch opt {
		case 'l':
			chars = append(chars, LowercaseLetters...)
		case 'u':
			chars = append(chars, UppercaseLetters...)
		case 'n':
			chars = append(chars, Numbers...)
		case 's':
			chars = append(chars, Symbols...)
		}
	}
	return shuffle(chars)
}

// genPwd generates a random password based on specified criteria.
//
// Parameters:
//   - length: The desired length of the password.
//   - chars: A slice of runes representing the character set to use for generation.
//   - exc: A string containing characters to exclude from the password.
//
// Returns:
//   A string representing the generated password.
//
// The function generates a password by randomly selecting characters from the
// provided character set (chars), ensuring that the password meets the specified
// length and does not include any characters listed in the exclude string (exc).
func genPwd(length int, chars []rune, exc string) string {
	pwd := make([]rune, length)
	var i int
	for i < length {
		char := chars[rand.IntN(len(chars))]
		if strings.ContainsRune(exc, char) {
			continue
		}
		pwd[i] = char
		i++
	}
	return string(pwd)
}

// shuffle randomizes the order of elements in a rune slice.
//
// It takes a slice of runes as input and returns a new slice with the same
// elements in a randomized order. The original slice is not modified.
//
// Parameters:
//   - r: The input slice of runes to be shuffled.
//
// Returns:
//   A new slice of runes with the elements in a randomized order.
func shuffle(r []rune) []rune {
	out := slices.Clone(r)
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out
}

// genpass generates random passwords based on specified criteria.
//
// Usage:
//   pwg [flags]
//
// Flags:
//   -len int
//         password length (default 12)
//   -inc string
//         characters to include: l,n,u,s for lowercase, uppercase, numbers, symbols respectively (default "l,u,n,s")
//   -exc string
//         list characters to exclude
//
// Examples:
//   Generate a default password:
//     pwg
//
//   Generate a 16-character password:
//     pwg -len 16
//
//   Generate a password with only lowercase letters and numbers:
//     pwg -inc l,n
//
//   Generate a password excluding specific characters:
//     pwg -exc "0O1Il"
//
// Note: If no inclusion options are specified, all character types will be used.
func main() {
	length := flag.Int("len", DefaultPasswordLength, "password length")
	include := flag.String("inc", "l,u,n,s", "l,n,u,s for lowercase, uppercase, numbers, symbols respectively")
	exclude := flag.String("exc", "", "list characters to exclude")
	flag.Parse()
	fmt.Println(genPwd(*length, compileChars(*include), *exclude))
}
