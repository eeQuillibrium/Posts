package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "sl;fg;dl2#$#t';,gdf#G}%${}#QR{A@#$}!F:sd,lgsmkgdlkfSDF.'sf,S"

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"userid,omitempty"`
}

func GenerateToken(
	ctx context.Context,
	userID int,
	tokenTTL time.Duration,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		},
		userID,
	})

	return token.SignedString([]byte(secretKey))
}

func ParseToken(
	ctx context.Context,
	tokenStr string,
) (int, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims doesn't have *TokenClaims type")
	}

	return claims.UserID, nil
}
