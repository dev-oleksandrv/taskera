package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hash == "" {
		t.Fatalf("expected non-empty hash")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	match := CheckPasswordHash(password, hash)
	if !match {
		t.Errorf("expected password to match hash")
	}

	wrongPassword := "wrongpassword"
	match = CheckPasswordHash(wrongPassword, hash)
	if match {
		t.Errorf("expected password to not match hash")
	}
}
