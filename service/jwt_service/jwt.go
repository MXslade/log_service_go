package jwt_service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/MXslade/log_service_go/model"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const JwtExpiresAfterHours = 72

type jwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginData struct {
	Username string
	Password string
}

type JwtService struct {
	secret []byte
}

func New() (*JwtService, error) {
	secret, ok := os.LookupEnv("SECRET_PHRASE")
	if !ok {
		return nil, errors.New("No SECRET_PHRASE is specified. Cannot hash passwords!")
	}
	return &JwtService{secret: []byte(secret)}, nil
}

func (j *JwtService) CreateConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: j.secret,
	}
}

func (j *JwtService) CreateToken(ctx context.Context, admin *model.AdminModel) (string, error) {
	claims := &jwtCustomClaims{
		admin.Username,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * JwtExpiresAfterHours))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}

	return t, nil
}
