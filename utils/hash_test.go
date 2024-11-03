package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mySecretPassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if hash == "" {
		t.Fatal("HashPassword returned an empty string")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mySecretPassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatal("CheckPasswordHash returned false for a valid password")
	}

	// Test with an incorrect password
	incorrectPassword := "wrongPassword"
	if CheckPasswordHash(incorrectPassword, hash) {
		t.Fatal("CheckPasswordHash returned true for an invalid password")
	}
}
