package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"io"
)

// sha256Base64 hashes a string to SHA256 and encodes it to base64.
// It panics on error.
func sha256Base64(str string) string {
	hash := sha256.New()

	if _, err := io.WriteString(hash, str); err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// hashPassword returns a bcrypt encoded password hash.
// The password is hashed using SHA256 first to limit its length.
func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(sha256Base64(password)), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// comparePassword compares the plain text password and the hash and returns true if they're equal or false otherwise.
func comparePassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(sha256Base64(password))) == nil
}
