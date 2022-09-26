package utils

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
)

func JwtUsersClaims(ctx context.Context, cfg *configs.Configs, authRepo entities.AuthRepository, req *entities.UsersJwtClaimsReq, claimsType entities.ClaimsType) (string, error) {
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
		ss, err := token.SignedString([]byte(cfg.App.JwtSecretKey))
		if err != nil {
			log.Println(err.Error())
			return "", errors.New("error, can't claims an access token")
		}
		return ss, nil
	case entities.RefreshToken:
		if req.UsersRefreshToken.ExpiresAt == nil && req.UsersRefreshToken.IssuedAt == nil {
			expiresAt := time.Now().Add(7 * time.Hour)
			IssuedAt := time.Now()
			req.UsersRefreshToken.ExpiresAt = &expiresAt
			req.UsersRefreshToken.IssuedAt = &IssuedAt
		} else {
			expiresAt := req.UsersRefreshToken.ExpiresAt.Add(-time.Duration(time.Now().Unix()))
			IssuedAt := time.Now()
			req.UsersRefreshToken.ExpiresAt = &expiresAt
			req.UsersRefreshToken.IssuedAt = &IssuedAt
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
		ss, err := token.SignedString([]byte(cfg.App.JwtSecretKey))
		if err != nil {
			log.Println(err.Error())
			return "", errors.New("error, can't claims an access token")
		}

		if err := authRepo.UpdateUserRefreshToken(ctx, req.UsersRefreshToken.Id, ss); err != nil {
			return "", err
		}
		return ss, nil
	case entities.SessionToken:
		if req.UsersSessionToken.ExpiresAt == nil && req.UsersSessionToken.IssuedAt == nil {
			expiresAt := time.Now().Add(7 * time.Hour)
			IssuedAt := time.Now()
			req.UsersSessionToken.ExpiresAt = &expiresAt
			req.UsersSessionToken.IssuedAt = &IssuedAt
		} else {
			expiresAt := req.UsersSessionToken.ExpiresAt.Add(-time.Duration(time.Now().Unix()))
			IssuedAt := time.Now()
			req.UsersSessionToken.ExpiresAt = &expiresAt
			req.UsersSessionToken.IssuedAt = &IssuedAt
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
		ss, err := token.SignedString([]byte(cfg.App.JwtSecretKey))
		if err != nil {
			log.Println(err.Error())
			return "", errors.New("error, can't claims an access token")
		}
		return ss, nil
	default:
		return "", errors.New("error, claims type is invalid")
	}
}
