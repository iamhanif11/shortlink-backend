package pkg

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// generate random slug 6 karakter
func GenRandomSlug() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var slugRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	s := make([]byte, 6)
	for i := range s {
		s[i] = charset[slugRand.Intn(len(charset))]
	}
	return string(s)
}

// validasi slug
func IsValidSlug(slug string) (bool, string) {
	if len(slug) < 3 || len(slug) > 50 {
		return false, "Slug length must be 3-50 characters"
	}
	reservedWord := map[string]bool{
		"api":       true,
		"login":     true,
		"register":  true,
		"dashboard": true,
	}
	if reservedWord[strings.ToLower(slug)] {
		return false, "Slug Cannot Use Reserved Words"
	}

	var validWithRegex = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !validWithRegex.MatchString(slug) {
		return false, "Slug can Only Contain Alphabet & Numeric and Hyphens"
	}

	return true, ""
}
