package auth_service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"os"
)

type AuthService interface {
	HashPassword(ctx context.Context, password string) string
	VerifyHash(ctx context.Context, toHash string, actualHash string) bool
}

type authService struct {
	secret string
	h      hash.Hash
}

func New() (*authService, error) {
	secret, ok := os.LookupEnv("SECRET_PHRASE")
	if !ok {
		return nil, errors.New("No SECRET_PHRASE is specified. Cannot hash passwords!")
	}
	return &authService{secret: secret, h: sha256.New()}, nil
}

func (a *authService) HashPassword(ctx context.Context, password string) string {
	a.h.Write([]byte(password + a.secret))
	result := a.h.Sum(nil)
	a.h.Reset()
	return hex.EncodeToString(result)
}

func (a *authService) VerifyHash(ctx context.Context, toHash string, actualHash string) bool {
	a.h.Write([]byte(toHash + a.secret))
	result := a.h.Sum(nil)
	a.h.Reset()
	return hex.EncodeToString(result) == actualHash
}
