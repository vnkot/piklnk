package repository_test

import (
	"testing"

	"github.com/vnkot/piklnk/internal/auth/domain"
	"github.com/vnkot/piklnk/internal/auth/repository"
	"github.com/vnkot/piklnk/pkg/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *db.Db {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect database:", err)
	}

	if err := gormDB.AutoMigrate(&repository.UserModel{}); err != nil {
		t.Fatal("migration failed:", err)
	}

	return &db.Db{DB: gormDB}
}

func cleanupTestDB(t *testing.T, db *db.Db) {
	if err := db.DB.Exec("DELETE FROM users").Error; err != nil {
		t.Fatal("cleanup failed:", err)
	}
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	testDB := setupTestDB(t)
	defer cleanupTestDB(t, testDB)

	repo := repository.NewUserRepository(testDB)

	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
	}
	if err := user.SetPassword("password123"); err != nil {
		t.Fatal("SetPassword failed:", err)
	}

	createdUser, err := repo.Create(user)
	if err != nil {
		t.Fatal("Create failed:", err)
	}
	if createdUser.ID == 0 {
		t.Error("User ID should be set")
	}

	foundUser, err := repo.FindByEmail("test@example.com")
	if err != nil {
		t.Fatal("FindByEmail failed:", err)
	}

	if createdUser.ID != foundUser.ID {
		t.Errorf("User IDs mismatch: created %d vs found %d", createdUser.ID, foundUser.ID)
	}
	if user.Email != foundUser.Email {
		t.Errorf("Emails mismatch: expected %s, got %s", user.Email, foundUser.Email)
	}
	if user.Name != foundUser.Name {
		t.Errorf("Names mismatch: expected %s, got %s", user.Name, foundUser.Name)
	}

	if !foundUser.IsValidPassword("password123") {
		t.Error("Password should be valid")
	}
	if foundUser.IsValidPassword("wrongpassword") {
		t.Error("Wrong password should be invalid")
	}
}

func TestUserRepository_FindByEmail_NotFound(t *testing.T) {
	testDB := setupTestDB(t)
	defer cleanupTestDB(t, testDB)

	repo := repository.NewUserRepository(testDB)

	_, err := repo.FindByEmail("notfound@example.com")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if err != gorm.ErrRecordNotFound {
		t.Errorf("Expected gorm.ErrRecordNotFound, got %v", err)
	}
}

func TestUserRepository_EmailUniqueness(t *testing.T) {
	testDB := setupTestDB(t)
	defer cleanupTestDB(t, testDB)

	repo := repository.NewUserRepository(testDB)
	email := "unique@example.com"

	user1 := &domain.User{Email: email}
	if _, err := repo.Create(user1); err != nil {
		t.Fatal("First create failed:", err)
	}

	user2 := &domain.User{Email: email}
	_, err := repo.Create(user2)
	if err == nil {
		t.Fatal("Expected error for duplicate email, got nil")
	}
}

func TestUserRepository_Conversion(t *testing.T) {
	testDB := setupTestDB(t)
	defer cleanupTestDB(t, testDB)

	repo := repository.NewUserRepository(testDB)

	original := &domain.User{
		Email: "conversion@test.com",
		Name:  "Conversion Test",
	}
	if err := original.SetPassword("testpassword"); err != nil {
		t.Fatal("SetPassword failed:", err)
	}

	created, err := repo.Create(original)
	if err != nil {
		t.Fatal("Create failed:", err)
	}

	model := repo.FromUserDomain(created)
	converted := repo.ToUserDomain(model)

	if created.ID != converted.ID {
		t.Errorf("ID mismatch: original %d vs converted %d", created.ID, converted.ID)
	}
	if created.Email != converted.Email {
		t.Errorf("Email mismatch: original %s vs converted %s", created.Email, converted.Email)
	}
	if created.Name != converted.Name {
		t.Errorf("Name mismatch: original %s vs converted %s", created.Name, converted.Name)
	}
	if created.Password != converted.Password {
		t.Error("Password mismatch")
	}
}
