package jwt_service

import (
	"context"
	"time"

	"github.com/MXslade/log_service_go/db/repo/admin_repo"
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

func CreateConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
}

func CreateToken(ctx context.Context, admin *admin_repo.AdminModel) (string, error) {
	claims := &jwtCustomClaims{
		admin.Username,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * JwtExpiresAfterHours))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
