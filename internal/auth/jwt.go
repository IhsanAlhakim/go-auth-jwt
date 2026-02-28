package auth

import (
	"errors"
	"time"

	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
}

var (
	ErrInvalidToken         = errors.New("Invalid Token")
	ErrInvalidSigningMethod = errors.New("Invalid Signing Method")
)

var (
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
)

func GenerateToken(userId string, username string, cfg config.Config) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			Issuer:    cfg.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		Name: username,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	jwtSigKey := []byte(cfg.JWTSigKey)

	signedToken, err := token.SignedString(jwtSigKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(cfg config.Config, tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		} else if method != JWT_SIGNING_METHOD {
			return nil, ErrInvalidSigningMethod
		}

		jwtSigKey := []byte(cfg.JWTSigKey)
		return jwtSigKey, nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
