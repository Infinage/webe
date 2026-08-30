package main

import (
	"strings"
	"testing"
	"unicode"
)

func Test_GenerateId(t *testing.T) {
	id := generateId()
	if len(id) != 32 {
		t.Errorf("Expected length: 32, got: %d", len(id))
	} else if strings.ContainsFunc(id, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}) {
		t.Errorf("Only expected Alphanumeric characters: %s", id)
	}
}
