package domain_test

import (
	"testing"

	"github.com/vnkot/piklnk/internal/auth/domain"
)

func TestUser_SetPassword(t *testing.T) {
	t.Run("successful password hashing", func(t *testing.T) {
		user := &domain.User{}
		password := "securePassword123!"

		err := user.SetPassword(password)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if user.Password == "" {
			t.Error("Hashed password should not be empty")
		}
	})
}

func TestUser_IsValidPassword(t *testing.T) {
	setupUser := func() *domain.User {
		user := &domain.User{}
		if err := user.SetPassword("correctPassword"); err != nil {
			t.Fatalf("Setup failed: %v", err)
		}
		return user
	}

	t.Run("correct password", func(t *testing.T) {
		user := setupUser()
		if !user.IsValidPassword("correctPassword") {
			t.Error("Expected valid password check")
		}
	})

	t.Run("incorrect password", func(t *testing.T) {
		user := setupUser()
		if user.IsValidPassword("wrongPassword") {
			t.Error("Expected invalid password check")
		}
	})

	t.Run("different case password", func(t *testing.T) {
		user := setupUser()
		if user.IsValidPassword("CORRECTpassword") {
			t.Error("Expected case sensitivity")
		}
	})

	t.Run("empty password against non-empty hash", func(t *testing.T) {
		user := setupUser()
		if user.IsValidPassword("") {
			t.Error("Empty password should not validate")
		}
	})
}

func TestPasswordHashingSecurity(t *testing.T) {
	t.Run("unique salts per hash", func(t *testing.T) {
		user1 := &domain.User{}
		user2 := &domain.User{}
		password := "samePassword"

		if err := user1.SetPassword(password); err != nil {
			t.Fatalf("Setup user1 failed: %v", err)
		}
		if err := user2.SetPassword(password); err != nil {
			t.Fatalf("Setup user2 failed: %v", err)
		}

		if user1.Password == user2.Password {
			t.Error("Identical passwords produced same hash")
		}
	})

	t.Run("empty password hash validation", func(t *testing.T) {
		user := &domain.User{Password: ""}
		if user.IsValidPassword("anyPassword") {
			t.Error("Validation should fail for empty stored password")
		}
	})
}
