package profile

import (
	"context"

	domain "github.com/user/gen-golang-fiber-by-cursor/internal/domain/user"
	"github.com/user/gen-golang-fiber-by-cursor/internal/repository"
)

type Service interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
}

type service struct{ repo repository.UserRepository }

func New(repo repository.UserRepository) Service { return &service{repo: repo} }

func (s *service) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.repo.FindByUsername(ctx, username)
}
