package service_test

import (
	"errors"
	"testing"

	"github.com/vnkot/piklnk/internal/auth/domain"
	"github.com/vnkot/piklnk/internal/auth/service"
)

type MockUserRepository struct {
	users    map[string]*domain.User
	createFn func(user *domain.User) (*domain.User, error)
	findFn   func(email string) (*domain.User, error)
}

func (m *MockUserRepository) Create(user *domain.User) (*domain.User, error) {
	if m.createFn != nil {
		return m.createFn(user)
	}

	if m.users == nil {
		m.users = make(map[string]*domain.User)
	}

	m.users[user.Email] = user
	user.ID = uint(len(m.users))

	return user, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	if m.findFn != nil {
		return m.findFn(email)
	}

	if user, ok := m.users[email]; ok {
		return user, nil
	}

	return nil, nil
}

func TestAuthService_Login_Success(t *testing.T) {
	userPassword := "password123"
	userEmail := "test@example.com"

	user := &domain.User{
		Email: userEmail,
	}
	user.SetPassword(userPassword)

	repo := &MockUserRepository{}
	repo.Create(user)

	userID, err := service.NewAuthService(repo).Login(userEmail, userPassword)

	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	if userID != 1 {
		t.Errorf("Expected user ID 1, got %d", userID)
	}
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	testUserPassword := "password123"
	wrongPassword := "password123Wrong"
	testUserEmail := "test@example.com"

	user := &domain.User{
		Email: testUserEmail,
	}
	user.SetPassword(testUserPassword)

	repo := &MockUserRepository{}
	repo.Create(user)

	_, err := service.NewAuthService(repo).Login(testUserEmail, wrongPassword)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	testUserPassword := "password123"
	testUserEmail := "test@example.com"

	_, err := service.NewAuthService(&MockUserRepository{}).Login(testUserEmail, testUserPassword)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestAuthService_Register_Success(t *testing.T) {
	testUserName := "user"
	testUserPassword := "password123"
	testUserEmail := "test@example.com"

	repo := &MockUserRepository{}
	authService := service.NewAuthService(repo)

	userID, err := authService.Register(testUserEmail, testUserPassword, testUserName)
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	if userID == 0 {
		t.Error("Expected user ID, got 0")
	}

	user, err := repo.FindByEmail(testUserEmail)
	if err != nil {
		t.Fatal("FindByEmail failed:", err)
	}
	if user == nil {
		t.Fatal("User not created")
	}
	if user.Name != testUserName {
		t.Errorf("Expected name 'New User', got '%s'", user.Name)
	}
}

func TestAuthService_Register_UserExists(t *testing.T) {
	repo := &MockUserRepository{}
	repo.Create(&domain.User{Email: "exists@example.com"})

	authService := service.NewAuthService(repo)

	_, err := authService.Register("exists@example.com", "password123", "Existing User")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestAuthService_Register_CreateError(t *testing.T) {
	testUserName := "user"
	testUserPassword := "password123"
	testUserEmail := "test@example.com"

	authService := service.NewAuthService(&MockUserRepository{
		createFn: func(user *domain.User) (*domain.User, error) {
			return nil, errors.New("database error")
		},
	})

	_, err := authService.Register(testUserEmail, testUserPassword, testUserName)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "database error" {
		t.Errorf("Expected 'database error', got: %v", err)
	}
}
