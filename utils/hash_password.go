package utils

import (
	"crypto/rand"
	"errors"

	"github.com/BearTS/go-gin-monolith/config"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password with a random salt.
func HashPassword(password string) ([]byte, error) {
	saltSize := config.Password.SaltLength
	// Generate a random salt
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	// Hash the password with the salt using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password+string(salt)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Combine the salt and hash
	return append(salt, hash...), nil
}

// VerifyPassword compares a bcrypt hash with a plaintext password and returns true if they match.
func VerifyPassword(hashedPassword []byte, password string) (bool, error) {
	saltSize := config.Password.SaltLength

	hash := hashedPassword

	// Extract the salt from the hash
	if len(hash) < saltSize+bcrypt.MinCost {
		return false, errors.New("invalid hash")
	}
	salt := hash[:saltSize]
	hashWithoutSalt := hash[saltSize:]

	// Verify the password with the salt using bcrypt
	if err := bcrypt.CompareHashAndPassword(hashWithoutSalt, []byte(password+string(salt))); err != nil {
		return false, err
	}

	return true, nil
}
