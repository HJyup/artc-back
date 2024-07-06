package auth

import (
	"github.com/nrednav/cuid2"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, cuid2.Generate())
	if err != nil {
		t.Errorf("error creating jwt token: %v", err)
	}

	if token == "" {
		t.Errorf("expected token, got empty string")
	}
}
