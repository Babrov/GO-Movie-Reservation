package repositories

import (
	"movies-app/constants"
	"movies-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Create(email, password, name string) error {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     constants.RoleUser,
	}

	return repo.DB.Create(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("Email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) SetRole(email, role string) error {
	user, err := r.GetByEmail(email)

	if err != nil {
		return err
	}

	user.Role = role
	return r.DB.Save(&user).Error
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}
