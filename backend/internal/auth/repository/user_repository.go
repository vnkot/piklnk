package repository

import (
	"github.com/vnkot/piklnk/internal/auth/domain"
	"github.com/vnkot/piklnk/pkg/db"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
	Name     string
}

func (UserModel) TableName() string {
	return "users"
}

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{database: database}
}

func (r *UserRepository) ToUserDomain(model *UserModel) *domain.User {
	return &domain.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}
}

func (r *UserRepository) FromUserDomain(domain *domain.User) *UserModel {
	return &UserModel{
		Model: gorm.Model{
			ID: domain.ID,
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	userModel := r.FromUserDomain(user)

	result := r.database.DB.Create(userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.ToUserDomain(userModel), nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user UserModel

	result := r.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.ToUserDomain(&user), nil
}
