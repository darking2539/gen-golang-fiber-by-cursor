package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	domain "github.com/user/gen-golang-fiber-by-cursor/internal/domain/user"
	"github.com/user/gen-golang-fiber-by-cursor/internal/repository"
)

type Service interface {
	Login(ctx context.Context, username, password string) (string, error)
	ParseToken(tokenString string) (*jwt.RegisteredClaims, error)
}

type service struct {
	repo      repository.UserRepository
	jwtSecret []byte
}

func New(repo repository.UserRepository, jwtSecret string) Service {
	return &service{repo: repo, jwtSecret: []byte(jwtSecret)}
}

func (s *service) Login(ctx context.Context, username, password string) (string, error) {
	u, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if u.Password != password {
		return "", errors.New("invalid credentials")
	}
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *service) ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	tok, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tok.Claims.(*jwt.RegisteredClaims); ok && tok.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// Optional helper to convert claims to user
func ClaimsToUser(claims *jwt.RegisteredClaims, u *domain.User) {
	if claims == nil || u == nil {
		return
	}
	u.Username = claims.Subject
}
