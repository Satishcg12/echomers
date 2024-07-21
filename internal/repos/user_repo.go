package repos

import (
	"github.com/satishcg12/echomers/internal/types"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindAll() ([]types.Users, error)
		FindByID(id uint) (types.Users, error)
		FindByEmail(email string) (types.Users, error)
		FindByUsername(username string) (types.Users, error)
		Create(user types.Users) (types.Users, error)
		Update(user types.Users) (types.Users, error)
		Delete(id uint) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]types.Users, error) {
	var users []types.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (types.Users, error) {
	var user types.Users
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (types.Users, error) {
	var user types.Users
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindByUsername(username string) (types.Users, error) {
	var user types.Users
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) Create(user types.Users) (types.Users, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Update(user types.Users) (types.Users, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&types.Users{}, id).Error
}
