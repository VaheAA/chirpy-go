package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "mysecret")

	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	userId, err := ValidateJWT(token, "mysecret")

	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	if userId.String() != userID.String() {
		t.Fatalf("ValidateJWT returned incorrect user ID")
	}

	// Case 2: wrong secret

	wrongSecretToken, _ := MakeJWT(userID, "mysecret")
	gotID, err := ValidateJWT(wrongSecretToken, "wrongsecret")

	if err == nil {
		t.Fatalf("ValidateJWT should have failed with wrong secret")
	}
	if gotID != uuid.Nil {
		t.Fatalf("ValidateJWT should return uuid.Nil on wrong secret")
	}

}
