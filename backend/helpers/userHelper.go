package helpers

import (
	"math/rand"
	"time"
)

// Function to generate a random string of a specified length
func RandomString(length int) string {
	// Define the character set from which to generate the string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create the random string
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}