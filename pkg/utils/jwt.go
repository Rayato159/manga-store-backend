package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
)

func JwtUsersClaims(cfg *configs.Configs, req *entities.UsersJwtClaimsReq, claimsType entities.ClaimsType) (string, error) {
	switch claimsType {
	case entities.AccessToken:
		claims := entities.UsersJwtTokenMapClaims{
			Id:   req.UsersAccessToken.Id,
			Role: req.UsersAccessToken.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "access_token",
				Subject:   "users_access_token",
				ID:        uuid.New().String(),
				Audience:  []string{"users"},
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(cfg.App.JwtSecretKey)
		if err != nil {
			return "", errors.New("error, can't claims an access token")
		}
		return ss, nil
	case entities.RefreshToken:
		if req.UsersRefreshToken.ExpiresAt == nil && req.UsersRefreshToken.IssuedAt == nil {
			*req.UsersRefreshToken.ExpiresAt = time.Now().Add(7 * time.Hour)
			*req.UsersRefreshToken.IssuedAt = time.Now()
		} else {
			*req.UsersRefreshToken.ExpiresAt = req.UsersRefreshToken.ExpiresAt.Add(-time.Duration(time.Now().Unix()))
			*req.UsersRefreshToken.IssuedAt = time.Now()
		}
		claims := entities.UsersJwtTokenMapClaims{
			Id:   req.UsersRefreshToken.Id,
			Role: req.UsersRefreshToken.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(*req.UsersRefreshToken.ExpiresAt),
				IssuedAt:  jwt.NewNumericDate(*req.UsersRefreshToken.IssuedAt),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "users_refresh_token",
				Subject:   "users",
				ID:        uuid.New().String(),
				Audience:  []string{"users"},
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(cfg.App.JwtSecretKey)
		if err != nil {
			return "", errors.New("error, can't claims an access token")
		}
		return ss, nil
	case entities.SessionToken:
		if req.UsersSessionToken.ExpiresAt == nil && req.UsersSessionToken.IssuedAt == nil {
			*req.UsersSessionToken.ExpiresAt = time.Now().Add(7 * time.Hour)
			*req.UsersSessionToken.IssuedAt = time.Now()
		} else {
			*req.UsersSessionToken.ExpiresAt = req.UsersSessionToken.ExpiresAt.Add(-time.Duration(time.Now().Unix()))
			*req.UsersSessionToken.IssuedAt = time.Now()
		}
		claims := entities.UsersJwtSessionMapClaims{
			Id:       req.UsersSessionToken.Id,
			Username: req.UsersSessionToken.Username,
			Role:     req.UsersSessionToken.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(*req.UsersSessionToken.ExpiresAt),
				IssuedAt:  jwt.NewNumericDate(*req.UsersSessionToken.IssuedAt),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "users_session_token",
				Subject:   "users",
				ID:        uuid.New().String(),
				Audience:  []string{"users"},
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(cfg.App.JwtSecretKey)
		if err != nil {
			return "", errors.New("error, can't claims an access token")
		}
		return ss, nil
	default:
		return "", errors.New("error, claims type is invalid")
	}
}
