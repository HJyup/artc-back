package auth

import (
	"github.com/google/uuid"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, uuid.New())
	if err != nil {
		t.Errorf("error creating jwt token: %v", err)
	}

	if token == "" {
		t.Errorf("expected token, got empty string")
	}
}
