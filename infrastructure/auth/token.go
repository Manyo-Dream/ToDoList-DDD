package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecert = []byte("very-secret-jwt")

type Claims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type TokenService interface {
	GenerateToken(ctx context.Context, id int64, username string) (string, error)
	ParseToken(ctx context.Context, token string) (*Claims, error)
}

type JWTTokenServiceImpl struct {
}

func NewJWTTokenService() TokenService {
	return &JWTTokenServiceImpl{}
}

func (t *JWTTokenServiceImpl) GenerateToken(ctx context.Context, id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:       uint64(id),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "todo-list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecert)
	return token, err
}

func (t *JWTTokenServiceImpl) ParseToken(ctx context.Context, token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecert, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
