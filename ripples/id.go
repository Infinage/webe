package main

import "math/rand/v2"

const charset = "abcdefghijklmnopqrstuvwxyz0ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

// generateId creates a 32 byte long random string
func generateId() string {
	const size = 32
	b := make([]byte, size)
	for i := range size {
		b[i] = charset[rand.IntN(len(charset))]
	}
	return string(b)
}
