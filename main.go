// go-random-password-generator: A simple CLI tool to generate secure random passwords with customizable options.
package main

import (
	"crypto/rand" // For cryptographically secure random number generation
	"flag"        // For parsing command-line flags
	"fmt"         // For formatted I/O
	"math/big"    // For handling big integers (needed by crypto/rand)
	"os"          // For exiting on error
)

// Command-line flag variables
var (
	length       int  // Desired password length
	useLowercase bool // Include lowercase letters
	useUppercase bool // Include uppercase letters
	useNumbers   bool // Include numbers
	useSymbols   bool // Include symbols
)

// Character sets for password generation
const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers          = "0123456789"
	symbols          = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

// Initialize command-line flags with default values and descriptions
func init() {
	flag.IntVar(&length, "length", 12, "length of the password")
	flag.BoolVar(&useLowercase, "lower", false, "include lowercase letters")
	flag.BoolVar(&useUppercase, "upper", false, "include uppercase letters")
	flag.BoolVar(&useNumbers, "number", false, "include numbers")
	flag.BoolVar(&useSymbols, "symbol", false, "include symbols")
}

// Entry point of the program
func main() {
	flag.Parse() // Parse command-line flags

	// If no character type flags are set, include all by default for convenience
	if !(useLowercase || useUppercase || useNumbers || useSymbols) {
		useLowercase = true
		useUppercase = true
		useNumbers = true
		useSymbols = true
	}

	// Build the character set based on user options
	charset := ""
	if useLowercase {
		charset += lowercaseLetters
	}
	if useUppercase {
		charset += uppercaseLetters
	}
	if useNumbers {
		charset += numbers
	}
	if useSymbols {
		charset += symbols
	}

	// Validate password length
	if length <= 0 {
		fmt.Println("Password length must be greater than zero")
		os.Exit(1)
	}

	// Generate the password
	password, err := generatePassword(length, charset)

	if err != nil {
		fmt.Println("Error generating password:", err)
		os.Exit(1)
	}

	fmt.Println("Generated password:", password)
}

// generatePassword creates a random password of the given length using the provided charset.
// It uses crypto/rand for cryptographically secure random number generation.

func generatePassword(length int, charset string) (string, error) {

	password := make([]byte, length) // Create a byte slice to hold the password characters

	for i := 0; i < length; i++ {
		// Generate a secure random index for the charset
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		if err != nil {
			return "", err // Return error if random generation fails
		}

		password[i] = charset[num.Int64()]
	}

	return string(password), nil // Return the generated password as a string
}
