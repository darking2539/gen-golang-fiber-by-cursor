package repository

import (
	"context"

	domain "github.com/user/gen-golang-fiber-by-cursor/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
}

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository { return &userRepository{db: db} }

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var u domain.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var u domain.User
	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}
