package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "mysecret", time.Hour)

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

	// Case 2: expired token
	expiredToken, _ := MakeJWT(userID, "mysecret", -time.Second)
	userId, err = ValidateJWT(expiredToken, "mysecret")

	if err == nil {
		t.Fatalf("ValidateJWT did not return error for expired token")
	}

	if userId != uuid.Nil {
		t.Fatalf("ValidateJWT returned non-nil user ID for expired token")
	}

	// Case 3: wrong secret

	wrongSecretToken, _ := MakeJWT(userID, "mysecret", time.Hour)
	gotID, err := ValidateJWT(wrongSecretToken, "wrongsecret")

	if err == nil {
		t.Fatalf("ValidateJWT should have failed with wrong secret")
	}
	if gotID != uuid.Nil {
		t.Fatalf("ValidateJWT should return uuid.Nil on wrong secret")
	}

}
