package repository

import (
	"errors"
	"fmt"

	"github.com/ryo-y222/delivery-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("user repository create: %w", err)
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email=?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//service層にgormを伝えないようnilを返す。
			return nil, nil
		}
		// %s=email,%w=err
		return nil, fmt.Errorf("user repository get by email=%s: %w", email, err)
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// service層にgormを伝えないようnilを返していたがerrros.goに切り分け。
			return nil, fmt.Errorf("user repository get by id=%d: %w", id, ErrUserNotFound)
		}
		return nil, fmt.Errorf("user repository get by id=%d: %w", id, err)
	}
	return &user, nil
}

func (r *UserRepository) Update(id uint, user *model.User) error {
	result := r.db.Model(&model.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("user repository update id=%d: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user repository update id=%d: user not found", id)
	}
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&model.User{}, id)

	if result.Error != nil {
		return fmt.Errorf("user repository delete id=%d: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user repository delete id=%d: user not found", id)
	}

	return nil
}
